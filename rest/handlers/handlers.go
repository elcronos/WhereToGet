package handlers

import (
	. "../../db/dataobjects"
	db "../../db/database"
	"github.com/valyala/fasthttp"
	"fmt"
	"strings"
	"encoding/json"
	"log"
)

type errorMessage struct {
	Status		int
	Message		string
}

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func GetAllPlaces(ctx *fasthttp.RequestCtx) {
	// Read
	var place Place
	var service Service
	var products []Product
	var index int
	places := []Place{}
	db.DB.Find(&places) // find places
	for index , place = range places {
		//Find Service
		db.DB.Where("id = ?", place.ServicesId).First(&service)
		places[index].Services = service
		//Find Products
		db.DB.Model(&place).Association("Products").Find(&products)
		places[index].Products = products
	}

	list, err := json.Marshal(&places)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func GetPlace(ctx *fasthttp.RequestCtx) {
	// Read
	var place Place
	var service Service
	var products []Product
	var error errorMessage
	// Raw SQL
	db.DB.First(&place, ctx.UserValue("id"))
	if place.Id != 0 {
		//Find Service
		db.DB.Where("id = ?", place.ServicesId).First(&service)
		place.Services = service
		//Find Products
		db.DB.Model(&place).Association("Products").Find(&products)
		place.Products = products
		p, err := json.Marshal(&place)
		if err == nil {
			fmt.Fprintf(ctx, string(p))
			ctx.SetStatusCode(fasthttp.StatusOK)
		} else {
			log.Fatal("Cannot encode to JSON ", err)
			error.Status = fasthttp.StatusInternalServerError
			error.Message = "Error in handler GetPlace"
			errorJson, _ := json.Marshal(&error)
			fmt.Fprintf(ctx, string(errorJson))
			ctx.SetStatusCode(error.Status)
		}
	}	else{
		error.Status = fasthttp.StatusNoContent
		error.Message = "There is not content this request"
		errorJson, _ := json.Marshal(&error)
		fmt.Fprintf(ctx, string(errorJson))
		ctx.SetStatusCode(error.Status)
	}
	// set some headers and status code first
	ctx.SetContentType("application/json")

}

func GetAllCountries(ctx *fasthttp.RequestCtx) {
	// Read
	countries := []Country{}
	db.DB.Find(&countries) // find places
	fmt.Println(countries)
	list, err := json.Marshal(&countries)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func GetCountry(ctx *fasthttp.RequestCtx) {
	// Read
	var country Country
	// Raw SQL
	db.DB.Where("id = ?", strings.ToUpper(ctx.UserValue("id").(string))).First(&country)
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	p, err := json.Marshal(&country)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
}

func GetAllServices(ctx *fasthttp.RequestCtx) {
	// Read
	services := []Service{}
	db.DB.Find(&services) // find places
	fmt.Println(services)
	list, err := json.Marshal(&services)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func GetService(ctx *fasthttp.RequestCtx) {
	// Read
	var service Service
	// Raw SQL
	db.DB.Where("id = ?", strings.ToUpper(ctx.UserValue("id").(string))).First(&service)
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	p, err := json.Marshal(&service)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
}

func GetAllProducts(ctx *fasthttp.RequestCtx) {
	// Read
	var country Country
	products := []Product{}
	db.DB.Find(&products) // find products
	for index, product := range products {
		//Find Country
		db.DB.Where("id = ?", product.CountryId).First(&country)
		products[index].Country = country
	}
	list, err := json.Marshal(&products)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func GetProduct(ctx *fasthttp.RequestCtx) {
	// Read
	var product Product
	var country Country
	// Raw SQL
	db.DB.Where("id = ?", strings.ToUpper(ctx.UserValue("id").(string))).First(&product)
	//Find Country
	db.DB.Where("id = ?", product.CountryId).First(&country)
	product.Country = country
	p, err := json.Marshal(&product)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
		ctx.SetStatusCode(fasthttp.StatusOK)
	} else{
		log.Fatal("Cannot encode to JSON", err)
	}
	ctx.SetContentType("application/json")
}
