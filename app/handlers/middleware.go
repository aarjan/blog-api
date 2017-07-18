package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aarjan/blog/app/api"
)

type AppHandler struct {
	api.AppServicer
	HanlderFunc func(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error)
}

func (app AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Println(r.URL.Path)
	val, err := app.HanlderFunc(w, r, app.AppServicer)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			msg := wrapper{nil, Meta{false, 404, err.Error()}}
			json.NewEncoder(w).Encode(msg)
			return
		}
		msg := wrapper{nil, Meta{false, 505, err.Error()}}
		json.NewEncoder(w).Encode(msg)

		return
	}

	msg := wrapper{val, Meta{true, 200, ""}}
	json.NewEncoder(w).Encode(msg)

}

func AuthMiddleware(h http.Handler) http.Handler {
	return nil
}
