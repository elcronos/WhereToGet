package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB	*gorm.DB
)

func Init(){
	var err error

	if DB, err = OpenConnection(); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to database, but got err=%+v", err))
	}
}

func OpenConnection() (db *gorm.DB, err error){

	fmt.Println("Testing postgres...")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable port=32768")

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	return db, err
}

type Place struct {
	Id			int 		`gorm:"primary_key; AUTO_INCREMENT"`
	Name		string		`gorm:"size:150"`
	Address		string		`gorm:"size:255"`
	Latitude	string		`gorm:"size:10"`
	Longitude	string		`gorm:"size:10"`
	Services	string		`gorm:"size:255"`
	Products	string		`gorm:"size:255"`
	CountryCode	string		`gorm:"size:2"`		//ISO "ALPHA-2 Code"
}


func IndexHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func HelloHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func PlacesHandler(ctx *fasthttp.RequestCtx) {
	// Read
	places := []Place{}
	DB.Find(&places) // find places
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
	DB.First(&place, ctx.UserValue("id"))
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

func main() {
	Init()
	//Database
	//Migrate the schema
	DB.AutoMigrate(&Place{})

	router := fasthttprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/hello/:name", HelloHandler)
	router.GET("/places", PlacesHandler)
	router.GET("/places/:id", PlaceById)
	log.Fatal(fasthttp.ListenAndServe(":3000", router.Handler))
}