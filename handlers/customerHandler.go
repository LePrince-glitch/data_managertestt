package handlers

import (
	"data_manager/database"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("customer-name")
	id := fmt.Sprint(time.Now().Unix())

	customer := database.Customer{
		Id:   id,
		Name: name,
	}
	customer.Add()
}

func CustomerForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/customer_form.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var customer database.Customer
	id := r.FormValue("customer-id")

	customer.Id = id
	customer.Name = r.FormValue("customer-name")

	customer.Update()

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(r.FormValue("customer-id"), 10, 62)

	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	database.DeleteCustomer(id)
	database.DeleteCustomer(id)

	http.Redirect(w, r, "/show-customer", http.StatusMovedPermanently)
}

func ShowCustomer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/customer.html")

	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	id := vars["id"]

	customer := database.GetCustomer(id)

	t.Execute(w, customer)
}

func ShowCustomers(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/dashboard.html")

	if err != nil {
		panic(err)
	}
	data := database.GetCustomersDetails()

	t.Execute(w, data)
}
