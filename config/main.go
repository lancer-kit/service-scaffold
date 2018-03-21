package config

import (
	"time"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/sirupsen/logrus"

	"gitlab.inn4science.com/vcg/go-common"
	"gitlab.inn4science.com/vcg/go-common/db"
	"gitlab.inn4science.com/vcg/go-common/log"
)

const ServiceName = "courier"

// config is a `Cfg` singleton var,
// for access use the `Config` method.
var config *Cfg

func Init(path string) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	config = new(Cfg)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("raw_config", rawConfig).
			Fatal("unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		logrus.New().
			WithError(err).
			Fatal("Invalid configuration")
	}

	if config.Workers == nil {
		config.FillDefaultServices()
	}

}

// Config returns the config obj.
func Config() *Cfg {
	return config
}

// InitLog initializes logger.
func InitLog() {
	_, err := log.Init(config.Log)
	if err != nil {
		log.Default.
			WithError(err).
			Fatal("Unable to init log")
	}
}

// InitDB initializes database connector.
func InitDB() {
	if !config.WaitForDB {
		err := db.Init(config.DB, log.Default)
		if err != nil {
			log.Default.WithError(err).Fatal("Can't to init database connection")
		}
		return
	}

	vcgtools.RetryIncrementally(
		5*time.Second,
		func() bool {
			err := db.Init(config.DB, log.Default)
			if err != nil {
				log.Default.WithError(err).Warning("Can't to init database connection")
			}
			return err == nil
		})
}
