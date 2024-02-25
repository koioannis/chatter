package logging

import (
	"strings"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type LoggerAdapter struct {
	*logrus.Logger
}

func NewLoggerAdapter(logger *logrus.Logger) *LoggerAdapter {
	return &LoggerAdapter{logger}
}

// LogEvent implements fxevent.Logger.
func (l *LoggerAdapter) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.WithFields(
			logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			},
		).Info("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.WithFields(
				logrus.Fields{
					"callee": e.FunctionName,
					"caller": e.CallerName,
				},
			).Error("OnStart hook failed", e.Err)
		} else {
			l.WithFields(
				logrus.Fields{
					"callee":  e.FunctionName,
					"caller":  e.CallerName,
					"runtime": e.Runtime.String(),
				},
			).Info("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.WithFields(logrus.Fields{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		}).Info("OnStop hook executing")

	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
				"error":  e.Err, // Note: Logrus does not have a direct method like zap.Error, so we add the error to the fields manually.
			}).Error("OnStop hook failed")
		} else {
			l.WithFields(logrus.Fields{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			}).Info("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"type":        e.TypeName,
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
				"error":       e.Err,
			}).Error("error encountered while applying options")
		} else {
			l.WithFields(logrus.Fields{
				"type":        e.TypeName,
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
			}).Info("supplied")
		}

	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.WithFields(logrus.Fields{
				"type":        rtype,
				"constructor": e.ConstructorName,
				"module":      e.ModuleName,
			}).Info("provided")
		}
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"module":      e.ModuleName,
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"error":       e.Err,
			}).Error("error encountered while applying options")
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.WithFields(logrus.Fields{
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
				"type":        rtype,
			}).Info("replaced")
		}
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
			}).Error("error encountered while replacing", e.Err)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.WithFields(logrus.Fields{
				"decorator":   e.DecoratorName,
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
				"type":        rtype,
			}).Info("decorated")
		}
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"stacktrace":  e.StackTrace,
				"moduletrace": e.ModuleTrace,
				"module":      e.ModuleName,
			}).Error("error encountered while applying options", e.Err)
		}
	case *fxevent.Run:
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"name":   e.Name,
				"kind":   e.Kind,
				"module": e.ModuleName, // Assuming moduleField(e.ModuleName) maps to a simple field, adjust accordingly
			}).WithError(e.Err).Error("error returned")
		} else {
			l.WithFields(logrus.Fields{
				"name":   e.Name,
				"kind":   e.Kind,
				"module": e.ModuleName, // Adjust if moduleField(e.ModuleName) needs different handling
			}).Info("run")
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.WithFields(logrus.Fields{
			"function": e.FunctionName,
			"module":   e.ModuleName, // Adjust if necessary
		}).Info("invoking")
	case *fxevent.Invoked:
		if e.Err != nil {
			l.WithFields(logrus.Fields{
				"error":    e.Err,
				"stack":    e.Trace, // Consider whether to include based on log readability
				"function": e.FunctionName,
				"module":   e.ModuleName, // Adjust if necessary
			}).Error("invoke failed")
		}
	case *fxevent.Stopping:
		l.WithFields(logrus.Fields{
			"signal": strings.ToUpper(e.Signal.String()),
		}).Info("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.WithError(e.Err).Error("stop failed")
		}
	case *fxevent.RollingBack:
		l.WithError(e.StartErr).Error("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.WithError(e.Err).Error("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.WithError(e.Err).Error("start failed")
		} else {
			l.Info("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.WithError(e.Err).Error("custom logger initialization failed")
		} else {
			l.WithFields(logrus.Fields{
				"function": e.ConstructorName,
			}).Info("initialized custom fxevent.Logger")
		}

	}
}
