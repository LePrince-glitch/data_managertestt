package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id   string
	Name string
}

type table struct {
	Customer_id   string
	Customer_name string
	Site_id       string
	Site_name     string
	Number_id     string
	Number_name   string
	Recharge_id   string
	Volume        int
	Date_re       string
	Date_exp      string
}

func DeleteCustomer(id int64) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM customer WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func GetCustomers() []Customer {
	db := InitConn()
	defer db.Close()

	var (
		customer  Customer
		customers []Customer
	)

	query := "SELECT * FROM customer"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&customer.Id, &customer.Name)
		customers = append(customers, customer)
	}

	return customers
}

func GetCustomer(id string) Customer {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM customer WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var customer Customer

	if row.Next() {
		row.Scan(&customer.Id, &customer.Name)
	}

	return customer
}

func (customer Customer) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO customer VALUES ($1, $2)`)
	r, err := db.Exec(query, customer.Id, customer.Name)
	checkIfError(err)

	return r
}

func (customer Customer) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE customer SET "name" = $1 WHERE "id" = $2`)
	r, err := db.Exec(query, customer.Name, customer.Id)
	checkIfError(err)

	return r
}

func GetCustomersDetails() []table {

	var row table
	var tables []table

	db := InitConn()
	defer db.Close()

	query := fmt.Sprint(`SELECT c.id, c.name, s.id, s.name, n.id, n.number, r.id, r.volume, r.date_re, r.date_exp FROM customer c, site s, number n, recharge r
	WHERE c.id = s.customer_id and s.id = n.site_id and n.id = r.number_id`)

	r, err := db.Query(query)
	checkIfError(err)

	for r.Next() {

		r.Scan(&row.Customer_id, &row.Customer_name, &row.Site_id, &row.Site_name, &row.Number_id, &row.Number_name, &row.Recharge_id, &row.Volume, &row.Date_re, &row.Date_exp)

		//row.Date_re, _ = time.Parse("2006-01-02", row.Date_re.Format("2006-01-02"))
		//row.Date_exp, _ = time.Parse("2006-01-02", row.Date_exp.Format("2006-01-02"))

		//fmt.Println(time.Parse("2006-01-02", row.Date_exp.Format("2006-01-02")))

		row.Date_re = row.Date_re[:10]
		row.Date_exp = row.Date_exp[:10]

		tables = append(tables, row)

	}

	return tables
}
