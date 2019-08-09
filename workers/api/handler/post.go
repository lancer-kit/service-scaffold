package handler

import (
	"net/http"

	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/service-scaffold/models"
)

func AddBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLogEntry(r)

	data := new(models.BuzzFeed)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into database")
	dataQ := models.NewQ(nil).BuzzFeed()
	err = dataQ.Insert(data)
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}

func AddDocument(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLogEntry(r)

	data := new(models.CustomDocument)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into couchdb")
	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		logger.WithError(err).Error("Can not establish connection with database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = docQ.AddDocument(data)
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}
