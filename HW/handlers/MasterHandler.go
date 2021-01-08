package handlers

import (
	"../models"
	"../utils"
	"html/template"
	"log"
	"net/http"
)

func MasterHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/master.html")
	if err != nil {
		log.Println(err)
		writeCode(HttpInternalServerError, w, ctx)
		return
	}
	var details []models.Detail
	ctx.DB.Find(&details)

	err = tmpl.Execute(w, details)
	if err != nil {
		log.Println(err)
		writeCode(HttpInternalServerError, w, ctx)
		return
	}
}
