package handlers

import (
	. "../../db/dataobjects"
	db "../../db/database"
	ex "../errors"
	"github.com/valyala/fasthttp"
	"github.com/jordan-wright/email"
	"fmt"
	"strings"
	"encoding/json"
	"log"
	"net/smtp"
)

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
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetPlace(ctx *fasthttp.RequestCtx) {
	// Read
	var place Place
	var service Service
	var products []Product
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

		} else {
			log.Fatal("Cannot encode to JSON ", err)
			ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
		}
	}	else{
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusNoContent, "")
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetAllCountries(ctx *fasthttp.RequestCtx) {
	// Read
	countries := []Country{}
	db.DB.Find(&countries) // find places
	list, err := json.Marshal(&countries)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetCountry(ctx *fasthttp.RequestCtx) {
	// Read
	var country Country
	// Raw SQL
	db.DB.Where("id = ?", strings.ToUpper(ctx.UserValue("id").(string))).First(&country)
	p, err := json.Marshal(&country)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
	} else{
		log.Fatal("Cannot encode to JSON ", err)
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetAllServices(ctx *fasthttp.RequestCtx) {
	// Read
	services := []Service{}
	db.DB.Find(&services) // find places
	list, err := json.Marshal(&services)
	if err == nil {
		for _, item := range list {
			// process here
			fmt.Fprintf(ctx, string(item))
		}
	} else{
		log.Fatal("Cannot encode to JSON ", err)
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetService(ctx *fasthttp.RequestCtx) {
	// Read
	var service Service
	// Raw SQL
	db.DB.Where("id = ?", strings.ToUpper(ctx.UserValue("id").(string))).First(&service)
	ctx.SetStatusCode(fasthttp.StatusOK)
	p, err := json.Marshal(&service)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
	} else{
		log.Fatal("Cannot encode to JSON ", err)
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
}

func GetAllProducts(ctx *fasthttp.RequestCtx){
	// Variables
	var country Country
	products := []Product{}
	//Read Parameters /products?name:camilo  or /products?country:colombia
	nameParam := ctx.QueryArgs().Peek("name")
	countryParam := ctx.QueryArgs().Peek("country")
	//If a parameter was sent for Products
	if nameParam != nil || countryParam != nil {
		if nameParam != nil {
			//Find product by name
			db.DB.Where("name LIKE ?", "%"+strings.ToUpper(string(nameParam))+"%").Find(&products) // find products
		}else{
			//Find Country by name
			db.DB.Where("name LIKE ?", "%"+strings.ToUpper(string(countryParam))+"%").First(&country)
			if country.Id != "" {
				//Find product by country
				db.DB.Where("country_id = ?", country.Id).Find(&products)
			}
		}
	}else {
		db.DB.Find(&products) // find products
	}
	//There are some values
	if len(products) > 0 {
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
		} else {
			log.Fatal("Cannot encode to JSON ", err)
			ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
		}
	}else{
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusNoContent, "")
	}
	// set some headers and status code first
	ctx = AllowCORS(ctx)
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
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
	ctx = AllowCORS(ctx)
}

func SendEmail(ctx *fasthttp.RequestCtx) {
	ctx = AllowCORS(ctx)
	var m Email
	byteArray := ctx.Request.Body()
	err := json.Unmarshal(byteArray, &m)
	if err == nil {
		e := email.NewEmail()
		e.From = m.Name+ "<"+m.Email+">"
		e.To = []string{"capcarde@gmail.com"}
		e.Subject = "WTG Contact Us Email from"
		e.Text = []byte(m.Message)
		e.HTML = []byte("<p>"+m.Message+"</p>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "capcarde@gmail.com", "PASSWORD", "smtp.gmail.com"))
		ctx.SetStatusCode(fasthttp.StatusOK)
	}else {
		log.Fatal("Error while sending email", err)
		ctx = ex.ErrorHandler(ctx, fasthttp.StatusInternalServerError, err.Error())
	}
}

func OptionsResponse(ctx *fasthttp.RequestCtx){
	ctx.Response.Header.Add("Allow", "GET,POST,OPTIONS")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	ctx.Response.Header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers")
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func AllowCORS(ctx *fasthttp.RequestCtx) *fasthttp.RequestCtx{
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	ctx.Response.Header.Add("Content-Type", "application/json")

	return ctx
}