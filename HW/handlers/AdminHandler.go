package handlers

import (
	"../models"
	"../utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type AdminHtmlPageObject struct {
	UrlParams    []UrlParam
	Details      *[]models.Detail
	Suppliers    *[]models.Supplier
	Stocks       *[]models.Stock
	DetailStocks *[]models.DetailStock
}

type UrlParam struct {
	Token  string
	Table  string
	Offset string
}

const PAGELIMIT = 15

func AdminHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	w.Header().Set("Content-Type", "text/html")
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		writeCode(HttpForbidden, w, ctx)
		return
	}
	tableStr := r.URL.Query().Get("table")
	offsetStr := r.URL.Query().Get("offset")

	var offset, _ = strconv.Atoi(offsetStr)

	var pageData = AdminHtmlPageObject{
		UrlParams: []UrlParam{
			{Token: token, Table: "details", Offset: offsetStr},
			{Token: token, Table: "suppliers", Offset: offsetStr},
			{Token: token, Table: "stocks", Offset: offsetStr},
			{Token: token, Table: "detail_stocks", Offset: offsetStr},
		},
	}

	switch tableStr {
	case "details":
		pageData.Details = &[]models.Detail{}
		ctx.DB.Offset(offset).Limit(PAGELIMIT).Find(pageData.Details)
	case "suppliers":
		pageData.Suppliers = &[]models.Supplier{}
		ctx.DB.Offset(offset).Limit(PAGELIMIT).Find(pageData.Suppliers)
	case "stocks":
		pageData.Stocks = &[]models.Stock{}
		ctx.DB.Offset(offset).Limit(PAGELIMIT).Find(pageData.Stocks)
	case "detail_stocks":
		pageData.DetailStocks = &[]models.DetailStock{}
		ctx.DB.Offset(offset).Limit(PAGELIMIT).Find(pageData.DetailStocks)
	}

	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/admin.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Println(err)
		return
	}

}
