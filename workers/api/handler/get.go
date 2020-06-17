package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"

	"github.com/lancer-kit/service-scaffold/models"
)

func AllBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLogEntry(r)

	dbQuery := models.NewQ(nil).BuzzFeed()
	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	ols, err := dbQuery.SetPage(&pageQuery).Select()
	if err != nil {
		logger.WithError(err).Error("unable to select")
		render.ServerError(w)
		return
	}
	if len(ols) == 0 {
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Buzz instances was successfully obtained")
	render.RenderListWithPages(w, pageQuery, int64(len(ols)), ols)
}

func GetBuzz(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_uid", uid)

	idINT, err := strconv.Atoi(uid)
	if err != nil {
		render.BadRequest(w, "invalid id, must be a number")
		return
	}

	dataQ := models.NewQ(nil).BuzzFeed()
	res, err := dataQ.GetByID(int64(idINT))
	if err != nil {
		logger.WithError(err).Error("can not get by id")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}

func GetAllDocument(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLogEntry(r)

	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		logger.WithError(err).Error("unable to create custom doc")
		render.ServerError(w)
		return
	}

	res, err := docQ.GetAllDocument(pageQuery)
	if err != nil {
		logger.WithError(err).Error("can not to get documents")
		render.ServerError(w)
		return
	}

	render.RenderListWithPages(w, pageQuery, int64(len(res)), res)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	logger := log.GetLogEntry(r).WithField("query_uid", uid)

	idINT, err := strconv.Atoi(uid)
	if err != nil {
		render.BadRequest(w, "invalid id, must be a number")
		return
	}

	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		logger.WithError(err).Error("unable to create custom doc")
		render.ServerError(w)
		return
	}

	res, err := docQ.GetDocument(idINT)
	if err != nil {
		logger.WithError(err).Error("can not to get document")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}
