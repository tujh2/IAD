package main

import (
	"./handlers"
	"./utils"
	"log"
	"net/http"
)

func main() {
	conf := utils.ReadConfig()
	applicationContext := utils.AppContext{
		Config: conf,
		DB: utils.InitDatabase(conf),
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.MasterHandler(writer, request, &applicationContext)
	})
	http.HandleFunc("/detail", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.DetailHandler(writer, request, &applicationContext)
	})
	http.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.AdminHandler(writer, request, &applicationContext)
	})

	http.HandleFunc("/api/delete", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.DeleteHandler(writer, request, &applicationContext)
	})

	http.HandleFunc("/api/update", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.UpdateHandler(writer, request, &applicationContext)
	})

	http.HandleFunc("/api/upload", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlers.UploadImage(writer, request, &applicationContext)
	})

	static := http.FileServer(http.Dir(conf.Root + "static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	images := http.FileServer(http.Dir(conf.Root + "images"))
	http.Handle("/img/", http.StripPrefix("/img/", images))

	err := http.ListenAndServe(applicationContext.Config.Host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
