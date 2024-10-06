package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Commercial struct {
	Id         int
	Last_name  string
	First_name string
	Role       string
	Address    string
}

func DeleteCommercial(id int) sql.Result {
	db := InitConn()
	defer db.Close()

	query := `DELETE FROM commercial WHERE "id" = $1`
	r, err := db.Exec(query, id)

	checkIfError(err)

	return r
}

func Getcommercials() []Commercial {
	db := InitConn()
	defer db.Close()

	var (
		commercial  Commercial
		commercials []Commercial
	)

	query := "SELECT * FROM commercial"
	rows, err := db.Query(query)
	checkIfError(err)

	for rows.Next() {
		rows.Scan(&commercial.Id, &commercial.Last_name, &commercial.First_name, &commercial.Role, &commercial.Address)
		commercials = append(commercials, commercial)
	}

	return commercials
}

func Getcommercial(id int) Commercial {
	db := InitConn()
	defer db.Close()

	query := `SELECT * FROM commercial  WHERE "id" = $1`
	row, err := db.Query(query, id)
	checkIfError(err)

	var commercial Commercial

	if row.Next() {
		row.Scan(&commercial.Id, &commercial.Last_name, &commercial.First_name, &commercial.Role, &commercial.Address)
	}

	return commercial
}

func (commercial Commercial) Add() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO commercial VALUES (DEFAULT, $1, $2, $3, $4)`)
	r, err := db.Exec(query, commercial.Last_name, commercial.First_name, commercial.Role, commercial.Address)
	checkIfError(err)

	return r
}

func (commercial Commercial) Update() sql.Result {
	db := InitConn()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE commercial SET "last_name" = $2, "first_name" = $3, "role" = $4, "address" = $5 WHERE "id" = $1`)
	r, err := db.Exec(query, commercial.Id, commercial.Last_name, commercial.First_name, commercial.Role, commercial.Address)
	checkIfError(err)

	return r
}
