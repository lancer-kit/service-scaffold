package config

import (
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/natswrap"
	"gitlab.inn4science.com/gophers/service-kit/tools"
	"gopkg.in/yaml.v2"
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
		config.FillDefaultWorkers()
	}

	initLog()
	initDB()
	initNats()
}

// Config returns the config obj.
func Config() *Cfg {
	return config
}

func initLog() {
	_, err := log.Init(config.Log)
	if err != nil {
		log.Default.
			WithError(err).
			Fatal("Unable to init log")
	}
}

func initDB() {
	if !config.WaitForDB {
		err := db.Init(config.DB, log.Default)
		if err != nil {
			log.Default.WithError(err).Fatal("Can't to init database connection")
		}
		return
	}

	tools.RetryIncrementally(
		5*time.Second,
		func() bool {
			err := db.Init(config.DB, log.Default)
			if err != nil {
				log.Default.WithError(err).Warning("Can't to init database connection")
			}
			return err == nil
		})
}

func initNats() {
	natswrap.SetConfig(&config.NATS)
}
