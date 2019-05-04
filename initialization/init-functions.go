package initialization

import (
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/sirupsen/logrus"
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
	natsx.SetConfig(&cfg.NATS)
	_, err := natsx.GetConn()
	return err
}
