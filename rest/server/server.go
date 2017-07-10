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
	routes.GET("/places", handler.GetAllPlaces)
	routes.GET("/places/:id", handler.GetPlace)
	routes.GET("/countries", handler.GetAllCountries)
	routes.GET("/countries/:id", handler.GetCountry)
	routes.GET("/services", handler.GetAllServices)
	routes.GET("/services/:id", handler.GetService)
	routes.GET("/products", handler.GetAllProducts)
	routes.GET("/products/:id", handler.GetProduct)

	//Routes Handler
	log.Fatal(fasthttp.ListenAndServe(":3000", routes.Handler))
}
