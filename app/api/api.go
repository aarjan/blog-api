package api

import (
	"database/sql"
)

type AppServicer interface {
}

type AppService struct {
	DB *sql.DB
}
