package handlers

import (
	"errors"
	"net/http"

	"github.com/aarjan/blog/app/api"

	"strconv"

	"github.com/aarjan/blog/app/models"
	"github.com/gorilla/mux"
)

// GET
func GetPosts(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	return models.GetPosts(env.DB)
}

// GetPost retrieves a particluar post
func GetPost(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	post := &models.Post{ID: id}

	err := post.GetPost(env.DB)
	return post, err
}

// InsertPost inserts a record in the posts table
func InsertPost(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	r.ParseForm()
	name := r.FormValue("name")
	content := r.FormValue("content")
	// categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
	env := app.(api.AppService)

	post := &models.Post{
		Name:       name,
		Content:    content,
		CategoryID: 1,
	}

	// Check if post exists
	post.GetPostByName(env.DB)
	if post.ID != 0 {
		// Exists
		return nil, errors.New("Record already exists")
	}

	// Create new post
	err := post.CreatePost(env.DB)

	return post, err
}

func DeletePost(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	post := models.Post{ID: id}

	// Check if record exists
	post.GetPostByID(env.DB)
	if post.ID == 0 {
		return nil, errors.New("Record not found")
	}

	// Delete record
	err := post.Delete(env.DB)
	return nil, err
}

func GetTags(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	return models.GetTags(env.DB)
}

func GetTag(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tag := &models.Tag{ID: id}

	err := tag.GetTag(env.DB)
	return tag, err
}

// InsertTag inserts a record in the posts table
func InsertTag(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	r.ParseForm()
	name := r.FormValue("name")
	env := app.(api.AppService)

	tag := &models.Tag{
		Name: name,
	}

	// Check if post exists
	tag.GetTagByName(env.DB)
	if tag.ID != 0 {
		// Exists
		return nil, errors.New("Record already exists")
	}

	// Create new tag
	err := tag.CreateTag(env.DB)

	return tag, err
}

func DeleteTag(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tag := models.Tag{ID: id}

	// Check if record exists
	tag.GetTagByID(env.DB)
	if tag.ID == 0 {
		return nil, errors.New("Record not found")
	}

	// Delete record
	err := tag.Delete(env.DB)
	return nil, err
}

func UpdateTag(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	r.ParseForm()
	name := r.FormValue("name")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tag := models.Tag{
		ID:   id,
		Name: name,
	}
	err := tag.UpdateTag(env.DB)
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

func InsertCategory(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	r.ParseForm()
	name := r.FormValue("name")
	env := app.(api.AppService)

	category := &models.Category{
		Name: name,
	}

	// Check if post exists
	category.GetCategoryByName(env.DB)
	if category.ID != 0 {
		// Exists
		return nil, errors.New("Record already exists")
	}

	// Create new category
	err := category.CreateCategory(env.DB)

	return category, err
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	category := models.Category{ID: id}

	// Check if record exists
	category.GetCategoryByID(env.DB)
	if category.ID == 0 {
		return nil, errors.New("Record not found")
	}

	// Delete record
	err := category.Delete(env.DB)
	return nil, err
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, app api.AppServicer) (interface{}, error) {
	env := app.(api.AppService)
	r.ParseForm()
	name := r.FormValue("name")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	category := models.Category{
		ID:   id,
		Name: name,
	}
	err := category.UpdateCategory(env.DB)
	return category, err

}
