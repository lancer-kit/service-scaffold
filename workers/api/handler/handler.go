package handler

import (
	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/models"
	"lancer-kit/service-scaffold/repo"

	"github.com/lancer-kit/armory/db"
	"github.com/sirupsen/logrus"
)

// Handler contains realization of http handlers
type Handler struct {
	pg    repo.PGRepoI
	couch *repo.CouchRepo
	log   *logrus.Entry
	bus   chan<- models.Event
}

func NewHandler(cfg *config.Cfg, entry *logrus.Entry, bus chan<- models.Event) *Handler {
	conn, err := db.NewConnector(cfg.DB.Config(), entry)
	if err != nil {
		entry.WithError(err).Fatal("unable to init db connector")
	}

	return &Handler{
		pg:    repo.NewPGRepo(conn),
		couch: repo.NewCouchRepo(cfg.CouchDB),
		log:   entry.WithField("app_layer", "api.Handler"),
		bus:   bus,
	}

}
