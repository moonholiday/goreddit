package main

import (
	"github.com/moonholiday/goreddit/postgres"
	"github.com/moonholiday/goreddit/web"
	"log"
	"net/http"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:Postgres@7881@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
