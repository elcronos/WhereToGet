package server

import (
	"log"
	handler "../handlers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func handleCORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
}

func init(){
	finish := make(chan bool)
	//FastHttpRouter
	routes := fasthttprouter.New()
	// Box File Server
	box := rice.MustFindBox("static").HTTPBox()
	http.Handle("/", handleCORS(http.FileServer(box)))
	go func() {
		http.ListenAndServe(":3001", nil)
	}()
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
	routes.POST("/emagil", handler.SendEmail)
	routes.OPTIONS("/email", handler.OptionsResponse)

	//Routes Handler
	go func() {
		log.Fatal(fasthttp.ListenAndServe(":3000", routes.Handler))
	}()

	<-finish
}
