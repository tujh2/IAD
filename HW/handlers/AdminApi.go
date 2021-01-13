package handlers

import (
	"HW/models"
	"HW/utils"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	tableStr := r.URL.Query().Get("table")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	switch tableStr {
	case "details":
		ctx.DB.Unscoped().Delete(&models.Detail{}, id)
	case "suppliers":
		ctx.DB.Unscoped().Delete(&models.Supplier{}, id)
	case "stocks":
		ctx.DB.Unscoped().Delete(&models.Stock{}, id)
	case "detail_stocks":
		ctx.DB.Unscoped().Delete(&models.DetailStock{}, id)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		writeHttpCode(http.StatusForbidden, w)
		return
	}

	tableStr := r.URL.Query().Get("table")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		writeHttpCode(http.StatusInternalServerError, w)
		return
	}

	switch tableStr {
	case "details":
		var detail = models.Detail{}
		err := json.Unmarshal(bodyBytes, &detail)
		if err != nil {
			log.Println(err)
			writeHttpCode(http.StatusBadRequest, w)
			return
		}
		ctx.DB.Save(&detail)
	case "suppliers":
		var supplier = models.Supplier{}
		err := json.Unmarshal(bodyBytes, &supplier)
		if err != nil {
			log.Println(err)
			writeHttpCode(http.StatusBadRequest, w)
			return
		}
		ctx.DB.Save(&supplier)
	case "stocks":
		var stock = models.Stock{}
		err := json.Unmarshal(bodyBytes, &stock)
		if err != nil {
			log.Println(err)
			writeHttpCode(http.StatusBadRequest, w)
			return
		}
		ctx.DB.Save(&stock)
	case "detail_stocks":
		var detailStock = models.DetailStock{}
		err := json.Unmarshal(bodyBytes, &detailStock)
		if err != nil {
			log.Println(err)
			writeHttpCode(http.StatusBadRequest, w)
			return
		}
		ctx.DB.Save(&detailStock)
	}

}

func UploadImage(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		writeHttpCode(http.StatusForbidden, w)
		return
	}

	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Println(err)
		writeHttpCode(http.StatusInternalServerError, w)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		writeHttpCode(http.StatusBadRequest, w)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		writeHttpCode(http.StatusInternalServerError, w)
		return
	}

	hasher := sha1.New()
	hasher.Write(fileBytes)
	fileId := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	err = ioutil.WriteFile(ctx.Config.Root+"images/"+fileId, fileBytes, 0644)
	if err != nil {
		log.Println(err)
		writeHttpCode(http.StatusInternalServerError, w)
		return
	}
}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		writeHttpCode(http.StatusForbidden, w)
		return
	}
	fileID := r.URL.Query().Get("fileID")
	if len(fileID) == 0 {
		writeHttpCode(http.StatusBadRequest, w)
		return
	}
	err := os.Remove(ctx.Config.Root + "images/" + fileID)
	if err != nil {
		writeHttpCode(http.StatusInternalServerError, w)
		return
	}

}
