package models

import "database/sql"

type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts,omitempty"`
}

// GetCategory provides all the required category, with all the posts related with it.
func (c *Category) GetCategory(db *sql.DB) error {
	const sqlQuery = `SELECT id,name FROM categories WHERE id=$1;`
	err := db.QueryRow(sqlQuery, c.ID).Scan(&c.ID, &c.Name)
	if err != nil {
		return err
	}

	const postQuery = `SELECT id,name,content FROM public.posts WHERE category_id=$1`
	query, err := db.Query(postQuery, c.ID)
	defer query.Close()
	if err != nil {
		return err
	}
	p := Post{}
	for query.Next() {
		query.Scan(&p.ID, &p.Name, &p.Content)
		c.Posts = append(c.Posts, p)
	}

	return nil
}

func (c *Category) CreateCategory(db *sql.DB) error {
	const sqlQuery = `INSERT INTO categories (name` +
		`)VALUES(` +
		`$1) RETURNING id;`
	return db.QueryRow(sqlQuery, &c.Name).Scan(&c.ID)
}

func GetCategories(db *sql.DB) ([]Category, error) {
	const sqlQuery = `SELECT * FROM categories;`
	query, err := db.Query(sqlQuery)
	defer query.Close()
	if err != nil {
		return nil, err
	}
	var categories []Category
	c := Category{}
	for query.Next() {
		query.Scan(&c.ID, &c.Name)
		categories = append(categories, c)
	}
	return categories, nil
}

func (c *Category) Delete(db *sql.DB) error {
	const sqlQuery = `DELETE FROM public.categories WHERE id=$1`
	_, err := db.Exec(sqlQuery, c.ID)
	return err
}

func (c *Category) GetCategoryByName(db *sql.DB) error {
	const sqlQuery = `SELECT id FROM public.categories WHERE name=$1;`
	return db.QueryRow(sqlQuery, c.Name).Scan(&c.ID)
}

func (c *Category) GetCategoryByID(db *sql.DB) error {
	const sqlQuery = `SELECT id FROM public.categories WHERE id=$1;`
	// TODO: Make it less hacky
	c1 := Category{}
	err := db.QueryRow(sqlQuery, c.ID).Scan(&c1.ID)
	c.ID = c1.ID
	return err
}

func (c *Category) UpdateCategory(db *sql.DB) error {
	const sqlQuery = `UPDATE categories SET name=$1 WHERE id=$2 RETURNING id,name`
	return db.QueryRow(sqlQuery, c.Name, c.ID).Scan(&c.ID, &c.Name)
}
