package model

type Product struct {
	Persistence
	Name string `bson:"Name"`
}

func NewProduct() *Product {

}
