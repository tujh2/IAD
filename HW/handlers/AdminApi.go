package handlers

import (
	"../models"
	"../utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("403 Access denied."))
		return
	}

	tableStr := r.URL.Query().Get("table")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	switch tableStr {
	case "details":
		ctx.DB.Delete(&models.Detail{}, id)
	case "suppliers":
		ctx.DB.Delete(&models.Supplier{}, id)
	case "stocks":
		ctx.DB.Delete(&models.Stock{}, id)
	case "detail_stocks":
		ctx.DB.Delete(&models.DetailStock{}, id)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("403 Access denied."))
		return
	}

	tableStr := r.URL.Query().Get("table")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500"))
		return
	}


	switch tableStr {
	case "details":
		var detail = models.Detail{}
		err := json.Unmarshal(bodyBytes, &detail)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500"))
			return
		}
		ctx.DB.Save(&detail)
	case "suppliers":
		var supplier = models.Supplier{}
		err := json.Unmarshal(bodyBytes, &supplier)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500"))
			return
		}
		ctx.DB.Save(&supplier)
	case "stocks":
		var stock = models.Stock{}
		err := json.Unmarshal(bodyBytes, &stock)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500"))
			return
		}
		ctx.DB.Save(&stock)
	case "detail_stocks":
		var detailStock = models.DetailStock{}
		err := json.Unmarshal(bodyBytes, &detailStock)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500"))
			return
		}
		ctx.DB.Save(&detailStock)
	}

}
