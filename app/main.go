package main

import (
	"log"
	"net/http"

	"github.com/aarjan/blog/app/api"
	h "github.com/aarjan/blog/app/handlers"
	"github.com/aarjan/blog/app/shared"
)

func main() {

	db := shared.NewDBConn()

	defer db.Close()
	app := api.AppService{db}
	log.Fatal(http.ListenAndServe(":3000", h.Handlers(app)))
}
