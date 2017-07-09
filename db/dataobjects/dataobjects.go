package dataobjects

type Product struct {
	Id		int	`gorm:"primary_key; AUTO_INCREMENT"`
	Name	string	`gorm:"size:255"`
	Alias	string	`gorm:"size:250"`
}

type Country struct {
	Id		string	`gorm:"primary_key;size:2"`
	Name	string	`gorm:"size:150""`
}

type Service struct {
	Id		string	`gorm:"primary_key;size:2"` //Two letter key
	Name	string	`gorm:"size:150""`
}

type Place struct {
	Id		int 		`gorm:"primary_key; AUTO_INCREMENT"`
	Name		string		`gorm:"size:150"`
	Address		string		`gorm:"size:255"`
	Latitude	string		`gorm:"size:10"`
	Longitude	string		`gorm:"size:10"`
	Services	Service
	Products	[]Product	`gorm:"many2many:place_products;"`
	Country	Country		//ISO "ALPHA-2 Code"
}

