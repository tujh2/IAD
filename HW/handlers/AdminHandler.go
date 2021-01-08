package handlers

import (
	"../utils"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request, ctx *utils.AppContext) {
	w.Header().Set("Content-Type", "text/html")
	token := r.URL.Query().Get("token")
	if token != ctx.Config.AdminToken {
		writeCode(HttpForbidden, w, ctx)
		return
	}

}
