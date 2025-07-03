package config

import (
	"net/http"

	admin "github.com/hsnylmz1283/goblog/admin/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//tit ADMIN
	r.GET("/admin", admin.Dashboard{}.Index)

	//tit SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
