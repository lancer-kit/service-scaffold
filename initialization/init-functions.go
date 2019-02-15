package initialization

import (
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/natswrap"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
)

type initModule string

var (
	DB   initModule = "database connection"
	NATS initModule = "NATS"
)

func initDatabase(cfg *config.Cfg, entry *logrus.Entry) error {
	return db.Init(cfg.DB, entry)
}

func initNATS(cfg *config.Cfg, entry *logrus.Entry) error {
	natswrap.SetConfig(&cfg.NATS)
	_, err := natswrap.GetConn()
	return err
}
