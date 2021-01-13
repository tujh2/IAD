package handlers

import (
	"../models"
	"../utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type StockHtmlPageObject struct {
	Stock     models.Stock
	Details   *[]models.DetailView
	Suppliers *[]models.Supplier
}

func StockHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	w.Header().Set("Content-Type", "text/html")

	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		writeCode(HttpNotFound, w, ctx)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeCode(HttpInternalServerError, w, ctx)
		return
	}

	var pageData StockHtmlPageObject
	query := ctx.DB.Find(&pageData.Stock, id)
	if query.RowsAffected == 0 {
		writeCode(HttpNotFound, w, ctx)
		return
	}

	var details []models.DetailView
	ctx.DB.
		Select("`details`.*, `detail_stocks`.count, `suppliers`.name as supplier_name").
		Model(&models.Detail{}).
		Joins("left join suppliers on suppliers.id = details.supplier").
		Joins("left join detail_stocks on details.id = detail_stocks.detail_id").
		Joins("left join stocks on stocks.id = detail_stocks.stock_id").
		Where("stocks.id = ?", pageData.Stock.ID).
		Limit(1000).
		Scan(&details)

	if len(details) != 0 {
		pageData.Details = &details
	}

	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/stock.html")
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
