package handler

import (
	"net/http"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func AllBuzz(w http.ResponseWriter, r *http.Request) {
	dbQuery := models.NewQ(nil).BuzzFeed()
	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		render.BadRequest(w, err)
		return
	}

	dbQuery.SetPage(&pageQuery)

	ols, err := dbQuery.Select()
	if err != nil {
		log.Default.Info(err)
		render.ServerError(w)
		return
	}
	if ols == nil {
		render.ResultNotFound.Render(w)
		return
	}

	render.WriteJSON(w, http.StatusOK, ols)
}

func GetValueFromMiddleware(w http.ResponseWriter, r *http.Request) {
	testParam := r.Context().Value("some_param")
	render.Success(w, testParam)
}
