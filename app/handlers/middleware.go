package handlers

import (
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	log.Println(r.URL.Path, r.Method)
	val, err := app.HanlderFunc(w, r, app.AppServicer)

	if err != nil {
		log.Println(err)

		msg := wrapper{nil, Meta{false, 505, err.Error()}}
		json.NewEncoder(w).Encode(msg)

		return
	}

	msg := wrapper{val, Meta{true, 200, "successfully completed method : " + r.Method}}
	json.NewEncoder(w).Encode(msg)

}

func AuthMiddleware(h http.Handler) http.Handler {
	return nil
}
