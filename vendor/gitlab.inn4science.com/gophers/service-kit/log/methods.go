package log

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/olivere/elastic"
	"github.com/onrik/logrus/filename"
	"github.com/onrik/logrus/sentry"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
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

// AddElasticHook adds hook that sends all error,
// logs to the ElasticSearch.
func AddElasticHook(config ElasticConfig) {
	client, err := elastic.NewClient(elastic.SetURL(config.URL),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(config.Username, config.Password),
	)
	if err != nil {
		Default.WithError(err).Error("unable to init elastic")
		return
	}
	host, _ := os.Hostname()
	level, _ := logrus.ParseLevel(config.Level)
	hook, err := elogrus.NewElasticHook(client, host, level, config.Index)
	if err != nil {
		Default.WithError(err).Error("unable to init elastic hook")
		return
	}
	Default.Logger.Hooks.Add(hook)
}

// DefaultForRequest returns default logger with included http.Request details.
func DefaultForRequest(r *http.Request) *logrus.Entry {
	return IncludeRequest(Default, r)
}

// IncludeRequest includes http.Request details into the log.Entry.
func IncludeRequest(log *logrus.Entry, r *http.Request) *logrus.Entry {
	reqID := middleware.GetReqID(r.Context())

	return log.
		WithFields(logrus.Fields{
			"req_id": reqID,
			"path":   r.URL.Path,
			"method": r.Method,
			"sender": r.Header.Get("X-Forwarded-For"),
		})
}
