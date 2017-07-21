package models

import (
	"database/sql"
)

type User struct {
	ID                 int    `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	AccessToken        string `json:"access_token"`
	VerificationToken  string `json:"verification_token,omitempty"`
	VerificationStatus bool   `json:"verification_status,omitempty"`
}

func (u *User) CreateUser(db *sql.DB) error {
	const sqlQuery = `INSERT into users( ` +
		`id,username,password,access_token,verification_token,verification_status ` +
		`)VALUES( ` +
		`$1,$2,$3,$4,$5,$6);`
	_, err := db.Exec(sqlQuery, u.ID, u.Username, u.Password, u.AccessToken, u.VerificationToken, u.VerificationStatus)
	return err
}

func GetUsers(db *sql.DB) ([]User, error) {
	const sqlQuery = `SELECT * FROM users`
	query, err := db.Query(sqlQuery)
	defer query.Close()
	if err != nil {
		return nil, err
	}
	res := []User{}
	u := User{}
	for query.Next() {
		query.Scan(&u.ID, &u.Username, &u.Password, &u.AccessToken, &u.VerificationToken, &u.VerificationStatus)
		res = append(res, u)
	}
	return res, nil
}
