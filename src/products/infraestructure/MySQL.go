package infraestructure

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
	"fmt"
	"log"
    "time"
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
	_, err := mysql.conn.DB.Exec("INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		log.Printf("Error al insertar un producto: %v", err)
		return err
	}
	fmt.Print("[MySQL] - Producto guardado")
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Product, error) {
	fmt.Print("[MySQL] - Lista de productos")
	rows, err := mysql.conn.DB.Query("SELECT id, name, price, quantity, created_at FROM products")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var products []entities.Product
    for rows.Next() {
        var createdAtStr string
        var product entities.Product
        if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &createdAtStr); err != nil {
            log.Fatal(err)
        }
        product.Created_at, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
    if err != nil {
        log.Fatal(err)
    }
        products = append(products, product)
    }
    return products,nil
	
}
func (mysql *MySQL) CheckMissingProducts() ([]entities.Product, error) {
	fmt.Print("[MySQL] - Lista de productos")
	rows, err := mysql.conn.DB.Query("SELECT id, name, price, quantity FROM products where quantity < 5")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var products []entities.Product
    for rows.Next() {
        var product entities.Product
        if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity); err != nil {
            log.Fatal(err)
        }
        products = append(products, product)
    }
    return products,nil
	
}

func (mysql *MySQL) Edit(id int,updatedProduct *entities.Product) error {
    _, err := mysql.conn.DB.Exec("UPDATE products SET name = ?, price = ?, quantity = ? WHERE id = ?", updatedProduct.Name, updatedProduct.Price, updatedProduct.Quantity, id)
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

//long polling
func (mysql *MySQL) CheckNewProducts() ([]entities.Product, error) {
	fmt.Print("[MySQL] - Lista de nuevos productos")
	rows, err := mysql.conn.DB.Query("SELECT id, name, price, quantity, created_at FROM products WHERE created_at >= DATE_SUB(NOW(), INTERVAL 2 WEEK)ORDER BY created_at DESC")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var products []entities.Product
    for rows.Next() {
        var createdAtStr string
        var product entities.Product
        if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &createdAtStr); err != nil {
            log.Fatal(err)
        }
       product.Created_at, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
    if err != nil {
        log.Fatal(err)
    }
        products = append(products, product)
    }
    return products,nil
	
}