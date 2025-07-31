# GoBlog 📝

**GoBlog**, Golang ile geliştirilmiş, yönetici paneli üzerinden dinamik olarak blog gönderileri ekleyebileceğiniz temel bir blog sistemidir.

## 🎯 Amacı

Bu proje, [bu eğitim serisi](https://www.youtube.com/playlist?list=PLs98OhOT86fgpH6IOa6I-OUyKG-LUzeUY) üzerinden ilerlenerek öğrenme amacıyla geliştirilmiştir. Proje boyunca amaç, Go dilini ve web geliştirme süreçlerini daha iyi kavramaktır.

## 🚀 Özellikler

- Admin paneli üzerinden:
  - Yeni blog gönderisi ekleme
  - Gönderi düzenleme
  - Gönderi silme
- Dinamik içerik yönetimi
- Temel kullanıcı giriş sistemi (Login)
- Statik dosya sunumu (`/admin/assets` ve `/uploads`)

## 🛠️ Kullanılan Teknolojiler

- Go (Golang)
- HTML
- CSS
- JavaScript
- Bootstrap

## 🖥️ Kurulum ve Çalıştırma

1. Repoyu Klonla:
- git clone https://github.com/htnylmz1283/goblog.git
- cd goblog
2. Modülleri indir:
-go mod tidy
3. Projeyi başlat:
- go run main.go
4. Tarayıcıda aç:
- http://localhost:8080/admin

## 🙏 Teşekkür

Bu proje, [sschrs](https://github.com/sschrs)'nin [goblog-demo](https://github.com/sschrs/goblog-demo) projesinden öğrenilerek oluşturulmuştur. Eğitim serisi için kendisine teşekkür ederim.