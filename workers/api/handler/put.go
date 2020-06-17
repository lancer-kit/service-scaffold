package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"

	"lancer-kit/service-scaffold/models"
)

func ChangeBuzz(w http.ResponseWriter, r *http.Request) {
	type inputData struct {
		Description string `json:"description"`
	}

	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_id", uid)

	idINT, err := strconv.Atoi(uid)
	if err != nil {
		render.BadRequest(w, "invalid id, must be a number")
		return
	}

	data := new(inputData)
	err = httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	dataQ := models.NewQ(nil).BuzzFeed()
	err = dataQ.UpdateBuzzDescription(int64(idINT), data.Description)
	if err != nil {
		logger.WithError(err).Error("Can not update buzzfeed description")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.Success(w, data)
}

func ChangeDocument(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_id", uid)
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		render.BadRequest(w, "invalid id, must be a number")
		return
	}

	data := new(models.CustomDocument)
	err = httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.ServerError(w)
		return
	}

	err = docQ.UpdateDocument(idINT, data)
	if err != nil {
		logger.WithError(err).Error("сan not update document")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.Success(w, "Document was updated successful")
}
