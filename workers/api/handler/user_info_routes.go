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

func (h *Handler) AddUserInfo(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	data := new(models.UserInfo)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into couchdb")
	userInfo, err := h.couch.UserInfo()
	if err != nil {
		logger.WithError(err).Error("Can not establish connection with database")
		render.ServerError(w)
		return
	}

	err = userInfo.AddUserInfo(data)
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	h.bus <- models.Event{
		Kind: models.EventUserInfoAdd,
		Info: map[string]interface{}{"id": data.UserID},
	}

	render.WriteJSON(w, 201, data)
}

func (h *Handler) GetAllUserInfo(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	userInfo, err := h.couch.UserInfo()
	if err != nil {
		logger.WithError(err).Error("unable to create custom doc")
		render.ServerError(w)
		return
	}

	res, err := userInfo.AllUserInfo(pageQuery)
	if err != nil {
		logger.WithError(err).Error("can not to get documents")
		render.ServerError(w)
		return
	}

	render.RenderListWithPages(w, pageQuery, int64(len(res)), res)
}

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("user_info_id", rawID)
	userInfoID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	userInfo, err := h.couch.UserInfo()
	if err != nil {
		logger.WithError(err).Error("unable to create custom doc")
		render.ServerError(w)
		return
	}

	res, err := userInfo.GetUserInfo(userInfoID)
	if err != nil {
		logger.WithError(err).Error("can not to get document")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}

func (h *Handler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("user_info_id", rawID)
	userInfoID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	data := new(models.UserInfo)
	err = httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	userInfo, err := h.couch.UserInfo()
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.ServerError(w)
		return
	}

	err = userInfo.UpdateUserInfo(userInfoID, data)
	if err != nil {
		logger.WithError(err).Error("сan not update document")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	h.bus <- models.Event{
		Kind: models.EventUserInfoUpdate,
		Info: map[string]interface{}{"id": userInfoID},
	}
	render.Success(w, "UserInfo was updated successful")
}

func (h *Handler) DeleteUserInfo(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	logger := log.IncludeRequest(h.log, r).WithField("user_info_id", rawID)
	userInfoID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("can not parse uid")
		render.BadRequest(w, "invalid buzz_id, should be a number")
		return
	}

	userInfo, err := h.couch.UserInfo()
	if err != nil {
		logger.WithError(err).Error("can not create custom document")
		render.ServerError(w)
		return
	}

	err = userInfo.DeleteUserInfo(userInfoID)
	if err != nil {
		logger.WithError(err).Error("can not delete document")
		render.ServerError(w)
		return
	}
	h.bus <- models.Event{
		Kind: models.EventUserInfoDelete,
		Info: map[string]interface{}{"id": userInfoID},
	}

	render.Success(w, "UserInfo was successfully deleted")
}
