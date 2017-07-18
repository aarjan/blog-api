package handlers

import (
	"github.com/aarjan/blog/app/api"
	"github.com/gorilla/mux"
)

type cat struct {
	name string
}

func Handlers(app api.AppServicer) *mux.Router {
	m := mux.NewRouter()
	// GET all items
	m.Handle("/api/v1/posts", AppHandler{app, GetPosts}).Methods("GET")
	m.Handle("/api/v1/tags", AppHandler{app, GetTags}).Methods("GET")
	m.Handle("/api/v1/categories", AppHandler{app, GetCategories}).Methods("GET")

	// GET particular item
	m.Handle("/api/v1/post/{id:[0-9]+}", AppHandler{app, GetPost}).Methods("GET")
	m.Handle("/api/v1/tag/{id:[0-9]+}", AppHandler{app, GetTag}).Methods("GET")
	m.Handle("/api/v1/category/{id:[0-9]+}", AppHandler{app, GetCategory}).Methods("GET")
	return m
}
