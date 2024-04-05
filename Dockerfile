FROM golang:1.22-alpine as build

# Required for tailwind executable
RUN apk add --update gcompat curl build-base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . . 

# Generate templ go files
RUN go install github.com/a-h/templ/cmd/templ@v0.2.648
RUN templ generate

# Generate tailwind styles.css
RUN wget -q -O tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.3/tailwindcss-linux-x64
RUN chmod +x ./tailwind
RUN ./tailwind  -i ./static/tailwind.css -o ./static/dist/styles.css

RUN CGO_ENABLED=0 go build -o /app/chatter cmd/main.go

FROM alpine:3.14 as release
WORKDIR /app

COPY --from=build /app/chatter chatter
COPY --from=build /app/static/dist static/dist

CMD [ "/app/chatter" ]

