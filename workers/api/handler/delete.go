package handler

import (
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/service-scaffold/models"
)

func DeleteBuzz(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	dataQ := models.NewQ(nil).BuzzFeed()
	err = dataQ.DeleteByID(int64(idINT))
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Data has been deleted successfully")
	render.Success(w, "success")
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = docQ.DeleteDocument(idINT)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	render.Success(w, "Document was successfully deleted")
}
