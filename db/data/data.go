package db

import (
	"github.com/jinzhu/gorm"
	. "../dataobjects"
)

/*
	Data that will be used to initialise the database in the AutoMigration
 */
 var (
	 places  []Place
	 countries  []Country
	 services []Service
	 product []Product
	 dbInitialised bool = false
 )

func InitialiseDB(db *gorm.DB) {
	//Test if there is data in the database
	checkDatabaseEmpty(db)
	if dbInitialised {
		/*
			COUNTRIES
	 	*/
		insertCountries(db)
		/*
			SERVICES
		*/
		insertServices(db)
		/*
			PRODUCTS
		*/
		//insertProducts(db)
		/*
			PLACES
		*/
	}
}

func checkDatabaseEmpty(db *gorm.DB){
	countPlaces, countCountries, countServices, countProducts := 0,0,0,0
	db.Find(&places).Count(&countPlaces)
	db.Find(&countries).Count(&countCountries)
	db.Find(&services).Count(&countServices)
	db.Find(&product).Count(&countProducts)
	if countPlaces == 0 && countCountries == 0 && countServices == 0 && countProducts == 0 {
		dbInitialised = true
	}
}

func insertCountries(db *gorm.DB){
	var query = "INSERT INTO countries VALUES(?,?)"
	/*
		SOUTH AMERICA
	 */
	db.Exec(query,"AR","ARGENTINA")
	db.Exec(query,"BO","BOLIVIA")
	db.Exec(query,"BR","BRAZIL")
	db.Exec(query,"CL","CHILE")
	db.Exec(query,"CO","COLOMBIA")
	db.Exec(query,"EC","ECUADOR")
	db.Exec(query,"GF","FRENCH GUIANA")
	db.Exec(query,"GY","GUYANA")
	db.Exec(query,"PY","PARAGUAY")
	db.Exec(query,"PE","PERU")
	db.Exec(query,"SR","SURINAME")
	db.Exec(query,"UY","URUGUAY")
	db.Exec(query,"VE","VENEZUELA")
	/*
		CENTRAL AMERICA AND NORTH AMERICA
	 */
	db.Exec(query,"BZ","BELIZE")
	db.Exec(query,"CR","COSTA RICA")
	db.Exec(query,"SV","EL SALVADOR")
	db.Exec(query,"GT","GUATEMALA")
	db.Exec(query,"HN","HONDURAS")
	db.Exec(query,"MX","MEXICO")
	db.Exec(query,"NI","NICARAGUA")
	db.Exec(query,"PA","PANAMA")
}

func insertServices(db *gorm.DB){
	var query = "INSERT INTO services VALUES(?,?)"
	/*
		TYPE OF SERVICES
	 */
	db.Exec(query,"ST","STORE")
	db.Exec(query,"SC","SHOPPING CENTER")
	db.Exec(query,"MK","MARKET")
	db.Exec(query,"CO","COMPANY")
	db.Exec(query,"PE","PEOPLE")
}

func insertProducts(db *gorm.DB){
	var query = "INSERT INTO products VALUES(?,?)"
	/*
		TYPE OF PRODUCTS
	 */
	db.Exec(query,"AGUARDIENTE ANTIOQUEÑO","")
	db.Exec(query,"AJIACO CONGELADO","")
	db.Exec(query,"AREPAS","")
	db.Exec(query,"AREQUIPE","")
	db.Exec(query,"BOCADILLO","")
	db.Exec(query,"BOM BOM BUM","")
	db.Exec(query,"CAFECITO","")
	db.Exec(query,"CHOCOLATINA JET","")
	db.Exec(query,"CHORIZO","")
	db.Exec(query,"CROQUETAS DE YUCA","")
	db.Exec(query,"HARINA PAN","")
	db.Exec(query,"MASA DE BUÑUELOS","")
	db.Exec(query,"PAN DE BONO","")
	db.Exec(query,"PANELA","")
	db.Exec(query,"PAPAS CRIOLLAS","")
	db.Exec(query,"PONY MALTA","")
	db.Exec(query,"PULPAS DE FRUTA","")
	db.Exec(query,"RON MEDELLIN","")
	db.Exec(query,"SANCOCHO CONGELADO","")
	db.Exec(query,"TAJADAS DE PLATANO","")
	db.Exec(query,"TOSTONES","")
	db.Exec(query,"TRIGUISAR","")
}