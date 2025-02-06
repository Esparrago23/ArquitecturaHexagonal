package infraestructure

import (
    "demo/src/core"
    "demo/src/categorys/domain/entities"
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

func (mysql *MySQL) Save(category *entities.Category) error {
    _, err := mysql.conn.DB.Exec("INSERT INTO categories (name, description) VALUES (?, ?)", category.Name, category.Description)
    if err != nil {
        log.Printf("Error al insertar una categoría: %v", err)
        return err
    }
    fmt.Print("[MySQL] - Categoría guardada")
    return nil
}

func (mysql *MySQL) GetAll() ([]entities.Category, error) {
    fmt.Print("[MySQL] - Lista de categorías")
    rows, err := mysql.conn.DB.Query("SELECT id, name, description FROM categories")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var categories []entities.Category
    for rows.Next() {
        var category entities.Category
        if err := rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
            log.Fatal(err)
        }
        categories = append(categories, category)
    }
    return categories, nil
}

func (mysql *MySQL) Edit(id int, updatedCategory *entities.Category) error {
    _, err := mysql.conn.DB.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?", updatedCategory.Name, updatedCategory.Description, id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("[MySQL] - Categoría actualizada")
    return nil
}

func (mysql *MySQL) Delete(id int) error {
    _, err := mysql.conn.DB.Exec("DELETE FROM categories WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("[MySQL] - Categoría eliminada con id:", id)
    return nil
}