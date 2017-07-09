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

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func Places(ctx *fasthttp.RequestCtx) {
	// Read
	places := []Place{}
	db.DB.Find(&places) // find places
	fmt.Println(places)
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

func PlaceById(ctx *fasthttp.RequestCtx) {
	// Read
	var place Place
	// Raw SQL
	db.DB.First(&place, ctx.UserValue("id"))
	// set some headers and status code first
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	p, err := json.Marshal(&place)
	if err == nil {
		fmt.Fprintf(ctx, string(p))
	} else{
		log.Fatal("Cannot encode to JSON ", err)
	}
}

func Countries(ctx *fasthttp.RequestCtx) {
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

func CountryById(ctx *fasthttp.RequestCtx) {
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

func Services(ctx *fasthttp.RequestCtx) {
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

func ServiceById(ctx *fasthttp.RequestCtx) {
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

