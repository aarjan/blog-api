package models

import (
	"database/sql"
)

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts,omitempty"`
}

func (t *Tag) GetTag(db *sql.DB) error {
	const sqlQuery = `SELECT id,name FROM tags WHERE id=$1;`
	err := db.QueryRow(sqlQuery, t.ID).Scan(&t.ID, &t.Name)
	if err != nil {
		return err
	}

	// Get all the posts
	const postQuery = `select posts.* ` +
		`FROM tags ` +
		`inner join tags_posts on tags.id=tags_posts.tag_id ` +
		`inner join posts  on tags_posts.post_id=posts.id ` +
		`where tags.id=$1;`
	query, err := db.Query(postQuery, t.ID)
	defer query.Close()
	if err != nil {
		return err
	}
	p := Post{}
	for query.Next() {
		query.Scan(&p.ID, &p.Name, &p.Content, &p.CategoryID)
		t.Posts = append(t.Posts, p)
	}
	return nil
}

func (t *Tag) CreateTag(db *sql.DB) error {
	const sqlQuery = `INSERT INTO tags (name` +
		`)VALUES(` +
		`$1);`
	_, err := db.Exec(sqlQuery, t.Name)
	return err
}

func GetTags(db *sql.DB) ([]Tag, error) {
	const sqlQuery = `SELECT * FROM tags;`
	query, err := db.Query(sqlQuery)
	defer query.Close()
	if err != nil {
		return nil, err
	}
	var tags []Tag
	t := Tag{}
	for query.Next() {
		query.Scan(&t.ID, &t.Name)
		tags = append(tags, t)
	}
	return tags, nil
}
