package server

import (
	"log"
	handler "../handlers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func init(){
	//FastHttpRouter
	routes := fasthttprouter.New()
	//All Routes
	routes.GET("/", handler.Index)
	routes.GET("/hello/:name", handler.Hello)
	routes.GET("/places", handler.Places)
	routes.GET("/places/:id", handler.PlaceById)
	routes.GET("/countries", handler.Countries)
	routes.GET("/countries/:id", handler.CountryById)
	routes.GET("/services", handler.Services)
	routes.GET("/services/:id", handler.ServiceById)

	//Routes Handler
	log.Fatal(fasthttp.ListenAndServe(":3000", routes.Handler))
}
