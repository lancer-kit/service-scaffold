package db

type Config struct {
	MaxIdleConns int   `json:"max_idle" yaml:"max_idle"`
	MaxOpenConns int   `json:"max_open" yaml:"max_open"`
	MaxLifetime  int64 `json:"max_lifetime" yaml:"max_lifetime"`
}
