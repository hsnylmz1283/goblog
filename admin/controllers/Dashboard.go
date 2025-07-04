package controllers

import (
	"fmt"
	"net/http"      //web sunucusu oluşturma, response gönderme gibi işlemleri içerir.
	"text/template" //HTML Şablonu. .html dosyalarına veri yerleştirip, dinamik içerik üretmeyi sağlar.

	"github.com/hsnylmz1283/goblog/admin/helpers" //El emeği göz nuru, yani BENİM PAKETİM. Dinamik olarak .html dosyalarını çeken yardımcı fonksiyon içeriyor.
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
