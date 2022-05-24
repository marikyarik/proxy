package main

//go:generate qtc -dir=templates

import (
	"api/controllers"
	"api/services/storage"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)


func main() {

	log.Println("Init services")
	storageManager := storage.NewStorageManager()

	log.Println("Init controllers")
	w := controllers.NewWebController(storageManager)
	p := controllers.NewProxyController(storageManager)

	log.Println("Init router")
	r := router.New()

	r.GET("/proxy-ui/", w.Dashboard)
	r.POST("/proxy-ui/config", w.SetConfig)
	r.POST("/proxy-ui/user", w.AddUser)
	r.POST("/proxy-ui/user/{id}", w.EditUser)
	r.DELETE("/proxy-ui/user/{id}", w.DeleteUser)
	r.PATCH("/proxy-ui/user/{id}", w.ActivateUser)
	r.ANY("/{path:*}", p.Proxy)

	r.PanicHandler  = func (ctx *fasthttp.RequestCtx, data interface{})  {
		log.Println(data)
		ctx.SetStatusCode(500)
		ctx.SetBodyString("Internal server error")
	}

	log.Println("Starting webserver")

	ln, err := reuseport.Listen("tcp4", ":80")
	if err != nil {
		log.Fatalf("error in reuseport listener: %s", err.Error())
	}

	if err = fasthttp.Serve(ln, r.Handler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err.Error())
	}
}
