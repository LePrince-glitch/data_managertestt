package main

import (
	"data_manager/cust"
	"data_manager/data"
	"net/http"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	file, _ := template.ParseFiles("views/dashboard.html")

	CustomersDetails := data.GetCustomersDetails()
	DataCustomers := data.GetCustomers()
	Users := cust.GetUsers()
	Commercials := cust.Getcommercials()
	Services := cust.GetServices()
	Customers := cust.GetCustomers()
	Projects := cust.GetProjectByAll()
	inter := "interconnexion"

	file.Execute(w, map[string]any{
		"CustomersDetails": CustomersDetails,
		"DataCustomers":    DataCustomers,
		"Users":            Users,
		"Commercials":      Commercials,
		"Services":         Services,
		"Customers":        Customers,
		"Projects":         Projects,
		"inter":            inter,
	})
}
