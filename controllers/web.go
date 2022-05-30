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

	var routes map[string]string
	if err := json.Unmarshal(ctx.Request.Body(), &routes); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	w.storageManager.SetConfig(routes)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (w *WebController) AddUser(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	var user storage.User

	if err := json.Unmarshal(ctx.Request.Body(), &user.Headers); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	u := w.storageManager.AddUser(user)
	jsonresonse, _ := json.Marshal(u)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(jsonresonse)
}

func (w *WebController) EditUser(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	var user storage.User

	if err := json.Unmarshal(ctx.Request.Body(), &user.Headers); err != nil {
		jsonBody, _ := json.Marshal(map[string]string{"error": "Form error"})
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(jsonBody)
		return
	}

	id := ctx.UserValue("id").(string)
	u := w.storageManager.EditUser(id, user)
	jsonresonse, _ := json.Marshal(u)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(jsonresonse)
}

func (w *WebController) DeleteUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	w.storageManager.DeleteUser(id)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (w *WebController) ActivateUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	w.storageManager.ActivateUser(id)
	ctx.SetStatusCode(fasthttp.StatusOK)
}