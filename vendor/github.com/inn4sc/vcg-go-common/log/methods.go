package log

import (
	"net/http"

	"github.com/onrik/logrus/filename"
	"github.com/onrik/logrus/sentry"
	"github.com/sirupsen/logrus"
)

// AddSentryHook adds hook that sends all error,
// fatal and panic log lines to the sentry service.
func AddSentryHook(dsn string) {
	sentryHook := sentry.NewHook(dsn,
		logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel)
	Default.Logger.AddHook(sentryHook)
}

// AddFilenameHook adds hook that includes
// filename and line number into the log.
func AddFilenameHook() {
	filenameHook := filename.NewHook()
	filenameHook.Field = "file"
	Default.Logger.AddHook(filenameHook)
}

// DefaultForRequest returns default logger with included http.Request details.
func DefaultForRequest(r *http.Request) *logrus.Entry {
	return IncludeRequest(Default, r)
}

// IncludeRequest includes http.Request details into the log.Entry.
func IncludeRequest(log *logrus.Entry, r *http.Request) *logrus.Entry {
	return log.
		WithField("path", r.URL.Path).
		WithField("method", r.Method).
		WithField("sender", r.Header.Get("X-Forwarded-For"))
}
