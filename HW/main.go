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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.MasterHandler(w, r, &applicationContext)
	})
	http.HandleFunc("/detail", func(writer http.ResponseWriter, request *http.Request) {
		handlers.DetailHandler(writer, request, &applicationContext)
	})
	http.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		handlers.AdminHandler(writer, request, &applicationContext)
	})

	static := http.FileServer(http.Dir(conf.Root + "static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	err := http.ListenAndServe(applicationContext.Config.Host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
