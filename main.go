package main

import (
	"net/http"

	"github.com/hsnylmz1283/goblog/config"
)

func main() {
	http.ListenAndServe(":8080", config.Routes())
}
