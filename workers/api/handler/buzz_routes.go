package handler

import (
	"net/http"
	"strconv"

	"lancer-kit/service-scaffold/models"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
)

func (h *Handler) AddBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	data := new(models.BuzzFeed)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into database")
	dataQ := h.pg.Clone().BuzzFeed()
	err = dataQ.Insert(data)
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	h.bus <- models.Event{
		Kind: models.EventBuzzAdd,
		Info: map[string]interface{}{
			"name": data.Name,
		},
	}

	render.WriteJSON(w, 201, data)
}

func (h *Handler) AllBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	dbQuery := h.pg.Clone().BuzzFeed()
	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	ols, total, err := dbQuery.SelectPage(&pageQuery)
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
	render.RenderListWithPages(w, pageQuery, total, ols)
}

func (h *Handler) GetBuzz(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("buzz_id", rawID)
	buzzID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	dataQ := h.pg.Clone().BuzzFeed()
	res, err := dataQ.GetByID(buzzID)
	if err != nil {
		logger.WithError(err).Error("can not get by id")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}

func (h *Handler) ChangeBuzz(w http.ResponseWriter, r *http.Request) {
	type inputData struct {
		Description string `json:"description"`
	}

	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("buzz_id", rawID)
	buzzID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	data := new(inputData)
	err = httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	dataQ := h.pg.Clone().BuzzFeed()
	err = dataQ.UpdateBuzzDescription(buzzID, data.Description)
	if err != nil {
		logger.WithError(err).Error("Can not update buzzfeed description")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	h.bus <- models.Event{
		Kind: models.EventBuzzUpdate,
		Info: map[string]interface{}{"id": buzzID},
	}

	render.Success(w, data)
}

func (h *Handler) DeleteBuzz(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("buzz_id", rawID)
	buzzID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	dataQ := h.pg.Clone().BuzzFeed()
	err = dataQ.DeleteByID(buzzID)
	if err != nil {
		logger.WithError(err).Error("can not delete BuzzFeed")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been deleted successfully")
	h.bus <- models.Event{
		Kind: models.EventBuzzDelete,
		Info: map[string]interface{}{"id": buzzID},
	}

	render.Success(w, "success")
}
