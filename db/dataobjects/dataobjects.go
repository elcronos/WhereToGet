package dataobjects

type Product struct {
	Id			int		`gorm:"primary_key; AUTO_INCREMENT"`
	Name		string	`gorm:"size:255"`
	Alias		string	`gorm:"size:250"`
	Country		Country	`gorm:"ForeignKey:Id;AssociationForeignKey:CountryId"` //ISO "ALPHA-2 Code"
	CountryId	string	`gorm:"primary_key;size:2"`
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
	Id			int 		`gorm:"primary_key; AUTO_INCREMENT"`
	Name		string		`gorm:"size:150"`
	Address		string		`gorm:"size:255"`
	Latitude	string		`gorm:"size:10"`
	Longitude	string		`gorm:"size:10"`
	Services	Service		`gorm:"ForeignKey:Id;AssociationForeignKey:ServicesId"`
	ServicesId	string		`gorm:"primary_key;size:2"`
	Products	[]Product	`gorm:"many2many:place_products;AssociationForeignKey:Id;ForeignKey:Id"`
}

