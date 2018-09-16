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
	//Elastic is a config for the ElasticSearch hook.
	Elastic *ElasticConfig `json:"elastic" yaml:"elastic"`
}

// ElasticConfig is a set of params for the ElasticSearch node.
type ElasticConfig struct {
	// URL is endpoint of the ElasticSearch node.
	URL string `json:"url" yaml:"url"`
	// Username for HTTP Basic Auth for the ElasticSearch node.
	Username string `json:"username" yaml:"username"`
	// Password for HTTP Basic Auth for the ElasticSearch node.
	Password string `json:"password" yaml:"password"`
	// Index is name of the index in ElasticSearch.
	Index string `json:"index" yaml:"index"`
	// Level is a log level for ElasticSearch events.
	Level string `json:"level" yaml:"level"`
}
