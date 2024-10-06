package main

import (
	"data_manager/handlers"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/customer-form", handlers.CustomerForm)
	http.HandleFunc("/add-customer", func(w http.ResponseWriter, r *http.Request) {

		handlers.AddCustomer(w, r)
		handlers.AddSite(w, r)
		handlers.AddNumber(w, r)
		defer handlers.AddRecharge(w, r)

		http.Redirect(w, r, "/show-customers", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/show-customers", handlers.ShowCustomers)
	http.HandleFunc("/show-customer/{id}", handlers.ShowCustomer)
	http.HandleFunc("/update-customer", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		handlers.UpdateCustomer(w, r)
		handlers.UpdateSite(w, r)
		handlers.UpdateNumber(w, r)
		handlers.UpdateRecharge(w, r)

		http.Redirect(w, r, "/show-customers", http.StatusPermanentRedirect)
	})
	http.HandleFunc("/delete-customer/{id}", handlers.DeleteCustomer)

	http.ListenAndServe(":5000", nil)
}
