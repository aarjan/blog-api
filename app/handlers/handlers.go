package handlers

import (
	"net/http"

	"github.com/aarjan/blog/app/api"

	"strconv"

	"github.com/aarjan/blog/app/models"
	"github.com/gorilla/mux"
)

// GET
func GetPosts(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	posts, err := models.GetPosts(env.DB)
	return posts, err
}

// GET
func GetPost(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	post := &models.Post{ID: id}

	err := post.GetPost(env.DB)
	return post, err
}

func GetTags(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	tags, err := models.GetTags(env.DB)
	return tags, err
}

func GetTag(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tag := &models.Tag{ID: id}

	err := tag.GetTag(env.DB)
	return tag, err
}

func GetCategories(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	categories, err := models.GetCategories(env.DB)
	return categories, err
}

func GetCategory(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category := &models.Category{ID: id}

	err := category.GetCategory(env.DB)
	return category, err
}
