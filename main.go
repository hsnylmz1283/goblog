package main

import (
	"net/http"

	admin_models "github.com/hsnylmz1283/goblog/admin/models"
	"github.com/hsnylmz1283/goblog/config"
)

func main() {
	admin_models.Post{}.Migrate()
	//todo Proje her çalıştığında aynı nesne tabloya eklenir. Bu yüzden eğitmen bu kısmı sildi.
	// admin_models.Post{
	// 	(Title)//parantezi sil ve öyle şalıştır.: "Go ile Web Programlama",
	// 	Slug:  "go-ile-web-programlama",
	// }.Add()

	//tit veritabanından tek bir veri çekme işlemi
	// post := admin_models.Post{}.Get("description = ?", "deneme")
	// fmt.Println(post.Title)

	//todo birden fazla veri çekme işlemi
	// fmt.Println(admin_models.Post{}.GetAll())

	//tit tekli update işlemi
	post := admin_models.Post{}.Get(1)
	// post.Update("slug", "web")

	//todo çoklu update işlemi
	// post.Updates(admin_models.Post{Title: "Python ile web programlama", Description: "test"})

	//tit Delete işlemi
	post.Delete()
	http.ListenAndServe(":8080", config.Routes())
}
