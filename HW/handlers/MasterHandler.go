package handlers

import (
	"../models"
	"../utils"
	"html/template"
	"log"
	"net/http"
)

type MasterHtmlPageOject struct {
	Details   []models.Detail
	Suppliers []models.Supplier
	Stocks    []models.Stock
}

func MasterHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/master.html")
	if err != nil {
		log.Println(err)
		writeCode(HttpInternalServerError, w, ctx)
		return
	}
	var pageData MasterHtmlPageOject
	ctx.DB.Find(&pageData.Details).Limit(1000)
	ctx.DB.Find(&pageData.Suppliers).Limit(1000)
	ctx.DB.Find(&pageData.Stocks).Limit(1000)

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Println(err)
		writeCode(HttpInternalServerError, w, ctx)
		return
	}
}
