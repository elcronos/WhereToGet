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
		insertProducts(db)
		/*
			PLACES
		*/
		insertPlaces(db)
	}
}

func checkDatabaseEmpty(db *gorm.DB){
	countPlaces, countCountries, countServices, countProducts := 0,0,0,0
	db.Find(&places).Count(&countPlaces)
	db.Find(&countries).Count(&countCountries)
	db.Find(&services).Count(&countServices)
	db.Find(&product).Count(&countProducts)
	db.Find(&places).Count(&countPlaces)
	if countPlaces == 0 && countCountries == 0 &&
		countServices == 0 && countProducts == 0 && countPlaces == 0 {
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
	db.Exec(query,"BT","BOTTLE SHOP")
}

func insertProducts(db *gorm.DB){
	var query = "INSERT INTO products VALUES(?,?,?,?)"
	/*
		TYPE OF PRODUCTS
	 */
	db.Exec(query,1,"AGUARDIENTE ANTIOQUEÑO","","CO")
	db.Exec(query,2,"AJIACO CONGELADO","","CO")
	db.Exec(query,3,"AREPAS","","CO")
	db.Exec(query,4,"AREQUIPE","","CO")
	db.Exec(query,5,"BOCADILLO","","CO")
	db.Exec(query,6,"BOM BOM BUM","","CO")
	db.Exec(query,7,"CAFECITO","","CO")
	db.Exec(query,8,"CHOCOLATINA JET","","CO")
	db.Exec(query,9,"CHORIZO","","CO")
	db.Exec(query,10,"CROQUETAS DE YUCA","","CO")
	db.Exec(query,11,"HARINA PAN","","CO")
	db.Exec(query,12,"MASA DE BUÑUELOS","","CO")
	db.Exec(query,13,"PAN DE BONO","","CO")
	db.Exec(query,14,"PANELA","","CO")
	db.Exec(query,15,"PAPAS CRIOLLAS","","CO")
	db.Exec(query,16,"PONY MALTA","","CO")
	db.Exec(query,17,"PULPAS DE FRUTA","","CO")
	db.Exec(query,18,"RON MEDELLIN","","CO")
	db.Exec(query,19,"SANCOCHO CONGELADO","","CO")
	db.Exec(query,20,"TAJADAS DE PLATANO","","CO")
	db.Exec(query,21,"TOSTONES","","CO")
	db.Exec(query,22,"TRIGUISAR","","CO")
}

func insertPlaces(db *gorm.DB){
	var query = "INSERT INTO places VALUES(?,?,?,?,?,?)"
	/*
		TYPE OF PRODUCTS
	 */
	db.Exec(query,1,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,2,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,3,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,4,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,5,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,6,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,7,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
	db.Exec(query,8,"Carlton Hotel Bottle Shop","247 Hay street, East Perth","0","0","BT")
}