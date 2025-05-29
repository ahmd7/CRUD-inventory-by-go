package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT id, name, quantity, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func getAProduct(db *sql.DB, key int) (*product, error) {
	query := "SELECT name, quantity, price FROM products WHERE id =?"
	row := db.QueryRow(query, key)
	var p product
	p.ID = key
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func addProduct(db *sql.DB, p *product) error {
	query := fmt.Sprintf("insert into products(name, quantity,price	) values('%v',%v,%v)", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}
func updateAProduct(db *sql.DB, p *product) error {
	query := fmt.Sprintf("update products set name='%v', quantity=%v, price=%v where id=%v", p.Name, p.Quantity, p.Price, p.ID)
	_, err := db.Exec(query)
	return err
}
func deleteAProduct(db *sql.DB, id int) error {
	query := fmt.Sprintf("DELETE FROM products WHERE id = %v", id)
	result, err := db.Exec(query)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no such row exists")
	}
	return err
}
