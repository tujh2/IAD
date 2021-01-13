package handlers

import (
	"HW/models"
	"HW/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type DetailHtmlPageObject struct {
	Detail   models.Detail
	Supplier *models.Supplier
	Stocks   *[]models.StockView
}

func DetailHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
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

	var pageData DetailHtmlPageObject
	query := ctx.DB.Find(&pageData.Detail, id)
	if query.RowsAffected == 0 {
		writeCode(HttpNotFound, w, ctx)
		return
	}

	pageData.Supplier = &models.Supplier{}
	query = ctx.DB.Find(pageData.Supplier, pageData.Detail.Supplier)
	if query.RowsAffected == 0 {
		pageData.Supplier = nil
	}

	var stocks []models.StockView
	ctx.DB.
		Select("`stocks`.*, `detail_stocks`.count").
		Model(&models.Stock{}).
		Joins("left join detail_stocks on stocks.id = detail_stocks.stock_id").
		Where("detail_stocks.detail_id = ?", pageData.Detail.ID).
		Limit(1000).
		Scan(&stocks)
	if len(stocks) != 0 {
		pageData.Stocks = &stocks
	}

	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/detail.html")
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
