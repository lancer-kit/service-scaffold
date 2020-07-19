package handler

import (
	"net/http"
	"strconv"

	"lancer-kit/service-scaffold/models"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
)

func (h *Handler) DeleteBuzz(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_uid", uid)

	idINT, err := strconv.Atoi(uid)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid uid, should be a number")
		return
	}

	dataQ := models.NewQ(nil).BuzzFeed()
	err = dataQ.DeleteByID(int64(idINT))
	if err != nil {
		logger.WithError(err).Error("can not delete BuzzFeed")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been deleted successfully")
	render.Success(w, "success")
}

func (h *Handler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_uid", uid)

	idINT, err := strconv.Atoi(uid)
	if err != nil {
		logger.WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	docQ, err := models.CreateCustomDocumentQ(h.Cfg)
	if err != nil {
		logger.WithError(err).Error("can not create custom document")
		render.ServerError(w)
		return
	}

	err = docQ.DeleteDocument(idINT)
	if err != nil {
		logger.WithError(err).Error("can not delete document")
		render.ServerError(w)
		return
	}

	render.Success(w, "Document was successfully deleted")
}
