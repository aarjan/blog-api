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
	m.StrictSlash(true)
	// GET all items
	m.Handle("/api/v1/posts", AppHandler{app, GetPosts}).Methods("GET")
	m.Handle("/api/v1/tags", AppHandler{app, GetTags}).Methods("GET")
	m.Handle("/api/v1/categories", AppHandler{app, GetCategories}).Methods("GET")

	// GET particular item
	m.Handle("/api/v1/post/{id:[0-9]+}", AppHandler{app, GetPost}).Methods("GET")
	m.Handle("/api/v1/tag/{id:[0-9]+}", AppHandler{app, GetTag}).Methods("GET")
	m.Handle("/api/v1/category/{id:[0-9]+}", AppHandler{app, GetCategory}).Methods("GET")

	// POST
	m.Handle("/api/v1/post", AppHandler{app, InsertPost}).Methods("POST")
	m.Handle("/api/v1/tag", AppHandler{app, InsertTag}).Methods("POST")
	m.Handle("/api/v1/category", AppHandler{app, InsertCategory}).Methods("POST")

	// DELETE
	m.Handle("/api/v1/post/{id:[0-9]+}", AppHandler{app, DeletePost}).Methods("DELETE")
	m.Handle("/api/v1/tag/{id:[0-9]+}", AppHandler{app, DeleteTag}).Methods("DELETE")
	m.Handle("/api/v1/category/{id:[0-9]+}", AppHandler{app, DeleteCategory}).Methods("DELETE")

	// PUT
	m.Handle("/api/v1/tag/{id:[0-9]+}", AppHandler{app, UpdateTag}).Methods("PUT")
	m.Handle("/api/v1/category/{id:[0-9]+}", AppHandler{app, UpdateCategory}).Methods("PUT")
	return m
}
