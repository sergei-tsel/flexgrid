package main

import (
	"flexgrid/internal/db"
	"flexgrid/internal/router"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db.Init()

	r := router.NewRouter()

	http.ListenAndServe(":8080", r)
}
