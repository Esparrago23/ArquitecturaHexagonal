package entities

type Product struct {
	Id int32
	Name string
	Price float32
}

func NewProduct(Name string,price float32) *Product{
	return &Product{Id:1,Name: Name, Price: price}
}

func (p *Product) GetName() string{
	return p.Name
}

func (p *Product) SetName(Name string) {
	p.Name = Name
}