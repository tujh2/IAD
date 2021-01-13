package handlers

import (
	"HW/models"
	"HW/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type SupplierHtmlPageObject struct {
	Supplier models.Supplier
	Details  *[]models.Detail
	Stocks   *[]models.Stock
}

func SupplierHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
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

	var pageData SupplierHtmlPageObject
	query := ctx.DB.Find(&pageData.Supplier, id)
	if query.RowsAffected == 0 {
		writeCode(HttpNotFound, w, ctx)
		return
	}

	var details []models.Detail
	ctx.DB.Find(&details).Where("supplier = ?", pageData.Supplier.ID)
	if len(details) != 0 {
		pageData.Details = &details
	}

	var stocks []models.Stock
	ctx.DB.
		Model(&models.Stock{}).
		Distinct().
		Joins("left join detail_stocks on stocks.id = detail_stocks.stock_id").
		Joins("left join details on details.id = detail_stocks.detail_id").
		Where("details.supplier = ?", pageData.Supplier.ID).
		Limit(1000).
		Scan(&stocks)

	if len(stocks) != 0 {
		pageData.Stocks = &stocks
	}

	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/supplier.html")
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
