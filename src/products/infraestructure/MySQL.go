package infraestructure

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}


func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product *entities.Product) error {
	_, err := mysql.conn.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		log.Printf("Error al insertar un producto: %v", err)
		return err
	}
	fmt.Print("[MySQL] - Producto guardado")
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Product, error) {
	fmt.Print("[MySQL] - Lista de productos")
	rows, err := mysql.conn.DB.Query("SELECT id, name, price FROM products")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var products []entities.Product
    for rows.Next() {
        var product entities.Product
        if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
            log.Fatal(err)
        }
        products = append(products, product)
    }
    return products,nil
	
}

func (mysql *MySQL) Edit(id int,updatedProduct *entities.Product) error {
    _, err := mysql.conn.DB.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", updatedProduct.Name, updatedProduct.Price, id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("[MySQL] - Producto actualizado")
	return nil
}
func (mysql *MySQL) Delete(id int) error {
    _, err := mysql.conn.DB.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("[MySQL] - Producto eliminado con id:",id)
	return nil
}
