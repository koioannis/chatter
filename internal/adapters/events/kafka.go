package events

import (
	"context"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/koioannis/chatter/internal/core/domain"
	"github.com/koioannis/chatter/internal/serialization"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func newWriter() *kafka.Writer {
	kafkaAddr := os.Getenv("KAFKA_ADDR")

	w := kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: []string{kafkaAddr},
			Topic:   "messages",
		},
	)

	return w
}

type KafkaPublisher struct {
	w   *kafka.Writer
	ser serialization.MessageSerializer
}

func NewKafkaMessagePublisher(serializer serialization.MessageSerializer) *KafkaPublisher {
	w := newWriter()

	return &KafkaPublisher{
		w:   w,
		ser: serializer,
	}
}

func (p *KafkaPublisher) Stop() error {
	return p.w.Close()
}

func (p *KafkaPublisher) Publish(message *domain.Message) error {
	bb, err := p.ser.Serialize(message)
	if err != nil {
		return err
	}

	return p.w.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: bb,
		},
	)
}

func newReader() *kafka.Reader {
	kafkaAddr := os.Getenv("KAFKA_ADDR")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{kafkaAddr},
		GroupID:     "message-notifications",
		Topic:       "messages",
		StartOffset: kafka.LastOffset,
		MaxBytes:    10e6, // 10 mb
	})
	return r
}

type KafkaSubscriber struct {
	mutex        sync.RWMutex
	reader       *kafka.Reader
	deserializer serialization.MessageDeserializer
	log          *logrus.Logger

	subs map[uuid.UUID]map[string]chan *domain.Message
}

func NewKafkaMessageSubscriber(deserializer serialization.MessageDeserializer, logger *logrus.Logger) *KafkaSubscriber {
	reader := newReader()

	subscriber := &KafkaSubscriber{
		reader:       reader,
		deserializer: deserializer,
		subs:         make(map[uuid.UUID]map[string]chan *domain.Message),
	}

	return subscriber
}

func (s *KafkaSubscriber) Start(ctx context.Context) error {
	for {
		m, err := s.reader.ReadMessage(ctx)
		if err != nil {
			s.reader.Close()
			return ctx.Err()
		}
		message, err := s.deserializer.Deserialize(m.Value)
		if err != nil {
			s.log.Error(err)
			continue
		}

		s.mutex.RLock()
		subs, ok := s.subs[message.RoomId]
		if !ok {
			continue
		}
		for _, sub := range subs {
			go func(sub chan<- *domain.Message, msg *domain.Message) {
				select {
				case sub <- msg:
				case <-ctx.Done():
				}
			}(sub, message)
		}
		s.mutex.RUnlock()
	}
}

func (s *KafkaSubscriber) Subscribe(roomId uuid.UUID, clientID string) <-chan *domain.Message {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.subs[roomId]
	if !ok {
		s.subs[roomId] = make(map[string]chan *domain.Message)
	}

	ch := make(chan *domain.Message, 100)
	s.subs[roomId][clientID] = ch

	return ch
}

func (s *KafkaSubscriber) Unsubscribe(roomId uuid.UUID, clientID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if subs, ok := s.subs[roomId]; ok {
		delete(subs, clientID)
	}
}
