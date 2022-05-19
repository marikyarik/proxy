package controllers

import (
	"api/services/storage"
	"time"

	"github.com/valyala/fasthttp"
)

type ProxyController struct {
	storageManager *storage.StorageManager
}

func NewProxyController(storageManager *storage.StorageManager) *ProxyController {
	return &ProxyController{storageManager: storageManager}
}

func (p *ProxyController) Proxy(ctx *fasthttp.RequestCtx) {
	ctx.Request.SetHost(p.storageManager.ProxyUrl) 
	ctx.Request.Header.Set("X-User", p.storageManager.GetActiveUser())
	err := fasthttp.DoTimeout(&ctx.Request, &ctx.Response, time.Second * 30)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
}