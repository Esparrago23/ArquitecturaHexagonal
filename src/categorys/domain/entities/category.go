package entities

type Category struct {
	Id          int32
	Name        string
	Description string
}

func NewCategory(Name string, Description string) *Category {
    return &Category{Id: 1, Name: Name, Description: Description}
}

func (c *Category) GetName() string {
    return c.Name
}

func (c *Category) SetName(Name string) {
    c.Name = Name
}
