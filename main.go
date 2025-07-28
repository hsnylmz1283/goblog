package main

import (
	"net/http"

	admin_models "github.com/hsnylmz1283/goblog/admin/models"
	"github.com/hsnylmz1283/goblog/config"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
