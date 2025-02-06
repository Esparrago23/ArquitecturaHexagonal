package entities
import "time"

type Product struct {
	Id int32
	Name string
	Price float32
	Quantity int32
	Created_at     time.Time
}

func NewProduct(Name string,price float32,Quantity int32, Created_at time.Time) *Product{
	return &Product{Id:1,Name: Name, Price: price, Quantity: Quantity, Created_at: Created_at}
}

func (p *Product) GetName() string{
	return p.Name
}

func (p *Product) SetName(Name string) {
	p.Name = Name
}