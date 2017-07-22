package dataobjects

type Product struct {
	Id			int		`gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	Name		string	`gorm:"size:255" json:"name"`
	Alias		string	`gorm:"size:250" json:"alias"`
	Country		Country	`gorm:"ForeignKey:Id;AssociationForeignKey:CountryId" json:"country"` //ISO "ALPHA-2 Code"
	CountryId	string	`gorm:"size:2" json:"countryId,omitempty"`
	FileName	string	`gorm:"size:250" json:"file"`
}

type Country struct {
	Id		string	`gorm:"primary_key;size:2" json:"id"`
	Name	string	`gorm:"size:150" json:"name"`
}


type Service struct {
	Id		string	`gorm:"primary_key;size:2" json:"id"` //Two letter key
	Name	string	`gorm:"size:150"`
}

type Place struct {
	Id			int 		`gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	Name		string		`gorm:"size:150" json:"name"`
	Address		string		`gorm:"size:255" json:"address"`
	Latitude	string		`gorm:"size:10" json:"lat"`
	Longitude	string		`gorm:"size:10" json:"lng"`
	Services	Service		`gorm:"ForeignKey:Id;AssociationForeignKey:ServicesId" json:"services"`
	ServicesId	string		`gorm:"size:2" json:"servicesId,omitempty"`
	Products	[]Product	`gorm:"many2many:place_products;AssociationForeignKey:Id;ForeignKey:Id" json:"products"`
}

