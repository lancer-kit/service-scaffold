package log

// Config is a options for the initialization
// of the default logrus.Entry.
type Config struct {
	// AppName identifier of the app.
	AppName string `json:"app_name" yaml:"app_name"`
	// Level is a string representation of the `lorgus.Level`.
	Level string `json:"level" yaml:"level"`
	// Sentry is a DSN string for sentry hook.
	Sentry string `json:"sentry" yaml:"sentry"`
	// AddTrace enable adding of the filename field into log.
	AddTrace bool `json:"add_trace" yaml:"add_trace"`
	// JSON enable json formatted output.
	JSON bool `json:"json" yaml:"json"`
}
