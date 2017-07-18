package models

import (
	"database/sql"
)

type Post struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id,omitempty"`

	// Extra
	Tags []Tag    `json:"tags,omitempty"`
	Cat  Category `json:"category,omitempty"`
}

func (p *Post) CreatePost(db *sql.DB) error {
	const sqlQuery = `insert into public.posts (` +
		`name,content,category_id` +
		`)values(` +
		`$1,$2,$3);`
	_, err := db.Exec(sqlQuery, p.Name, p.Content, p.CategoryID)
	return err
}

func (p *Post) GetPost(db *sql.DB) error {
	const sqlQuery = `SELECT id,name,content,category_id FROM public.posts WHERE id=$1;`
	err := db.QueryRow(sqlQuery, p.ID).Scan(&p.ID, &p.Name, &p.Content, &p.CategoryID)
	if err != nil {
		return err
	}

	// Get the tags associated with the post
	const tagsQuery = `SELECT tags.* ` +
		`FROM posts ` +
		`INNER JOIN tags_posts on tags_posts.post_id = posts.id ` +
		`INNER JOIN tags on tags.id = tags_posts.tag_id ` +
		`WHERE posts.id=$1`

	query, err := db.Query(tagsQuery, p.ID)
	defer query.Close()
	if err != nil {
		return err
	}
	t := Tag{}
	for query.Next() {
		query.Scan(&t.ID, &t.Name)
		p.Tags = append(p.Tags, t)
	}

	// Get the category from category_id
	const catQuery = `SELECT id,name FROM categories WHERE id=$1`
	err = db.QueryRow(catQuery, p.CategoryID).Scan(&p.Cat.ID, &p.Cat.Name)
	return err
}

func GetPosts(db *sql.DB) ([]Post, error) {
	const sqlQuery = `SELECT * FROM posts;`
	query, err := db.Query(sqlQuery)
	defer query.Close()
	if err != nil {
		return nil, err
	}
	var posts []Post
	p := Post{}
	for query.Next() {
		query.Scan(&p.ID, &p.Name, &p.Content)
		posts = append(posts, p)
	}
	return posts, nil
}
