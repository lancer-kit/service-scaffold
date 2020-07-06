package initialization

import (
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/natsx"
	"github.com/sirupsen/logrus"

	"lancer-kit/service-scaffold/config"
)

type initModule string

const (
	DB   initModule = "database connection"
	NATS initModule = "NATS"
)

func initDatabase(cfg *config.Cfg, entry *logrus.Entry) error {
	return db.Init(cfg.DB.ConnURL, entry)
}

// nolint: unparam
func initNATS(cfg *config.Cfg, entry *logrus.Entry) error {
	natsx.SetConfig(&cfg.NATS)
	_, err := natsx.GetConn()
	return err
}
