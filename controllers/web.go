package controllers

import (
	"api/services/storage"
	"api/templates"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type WebController struct {
	storageManager *storage.StorageManager
}

func NewWebController(storageManager *storage.StorageManager) *WebController {
	return &WebController{storageManager: storageManager}
}

func (w *WebController) Dashboard(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("text/html")

	c, err := json.Marshal(w.storageManager)
	if err != nil {
		ctx.Error("Config error", fasthttp.StatusInternalServerError)
		return
	}

	templates.WriteDashboard(ctx, c)
}

func (w *WebController) SetConfig(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	if err := json.Unmarshal(ctx.Request.Body(), &w.storageManager); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	w.storageManager.Save()
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (w *WebController) AddUser(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	var user interface{}

	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	hash := w.storageManager.AddUser(user)
	jsonresonse, _ := json.Marshal(map[string]string{"hash": hash})
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(jsonresonse)
}

func (w *WebController) EditUser(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	var user interface{}

	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	hash := ctx.UserValue("hash").(string)
	w.storageManager.EditUser(hash, user)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (w *WebController) DeleteUser(ctx *fasthttp.RequestCtx) {
	hash := ctx.UserValue("hash").(string)
	w.storageManager.DeleteUser(hash)
	ctx.SetStatusCode(fasthttp.StatusOK)
}