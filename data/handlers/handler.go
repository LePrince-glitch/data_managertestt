package handlers

import (
	"data_manager/data"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var site data.Site

	if r.FormValue("site-id") != "" {
		id := r.FormValue("site-id")

		site = data.GetSite(id)

	} else {
		sites := data.GetSites()
		site = sites[len(sites)-1]
	}

	number_id := fmt.Sprint(time.Now().Unix())
	_number := r.FormValue("customer-number")

	number := data.Number{
		Id:      number_id,
		Number:  _number,
		Site_id: site.Id,
	}

	number.Add()
}

func AddSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var customer data.Customer

	if r.FormValue("-id") != "" {
		id := r.FormValue("customer-id")

		customer = data.GetCustomer(id)

	} else {
		customers := data.GetCustomers()
		customer = customers[len(customers)-1]

	}

	id := fmt.Sprint(time.Now().Unix())
	name := r.FormValue("site-name")

	site := data.Site{
		Id:          id,
		Name:        name,
		Customer_id: customer.Id,
	}

	site.Add()
}

func AddRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number

	if r.FormValue("number-id") != "" {
		number = data.GetNumber(r.FormValue("number-id"))
	} else {
		numbers := data.GetNumbers()
		number = numbers[len(numbers)-1]
	}

	id := fmt.Sprint(time.Now().Unix())

	volume64, _ := strconv.ParseInt(r.FormValue("data"), 10, 64)
	volume := int(volume64)
	date_re, _ := time.Parse("2006-01-02", r.FormValue("date-re"))
	date_exp := date_re.AddDate(0, 1, 1)
	auto_re, _ := strconv.ParseBool(r.FormValue("auto-re"))

	recharge := data.Recharge{
		Id:        id,
		Volume:    volume,
		Date_re:   date_re,
		Date_exp:  date_exp,
		Auto_re:   auto_re,
		Number_id: number.Id,
	}

	recharge.Add()
}

func DeleteRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recharge data.Recharge
	recharge.Id = r.FormValue("recharge-id")
	v, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	recharge.Volume = int(v)
	recharge.Date_re, _ = time.Parse("2006-01-02", r.FormValue("date-re"))
	recharge.Date_exp = recharge.Date_re.AddDate(0, 1, 1)

	data.DeleteRecharge(recharge.Id)
}

func UpdateRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recharge data.Recharge
	recharge.Id = r.FormValue("recharge-id")
	v, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	recharge.Volume = int(v)
	recharge.Date_re, _ = time.Parse("2006-01-02", r.FormValue("date-re"))
	recharge.Date_exp = recharge.Date_re.AddDate(0, 1, 1)

	recharge.Update()
}

func DeleteNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number
	number.Id = r.FormValue("number-id")
	number.Number = r.FormValue("customer-number")

	data.DeleteNumber(number.Id)
}

func UpdateNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number data.Number
	number.Id = r.FormValue("number-id")
	number.Number = r.FormValue("customer-number")

	number.Update()
}

func DeleteSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var site data.Site
	site.Id = r.FormValue("site-id")
	site.Name = r.FormValue("site-name")

	data.DeleteSite(site.Id)
}
func UpdateSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var site data.Site
	site.Id = r.FormValue("site-id")
	site.Name = r.FormValue("site-name")

	site.Update()
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("customer-name")
	id := fmt.Sprint(time.Now().Unix())

	customer := data.Customer{
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

	var customer data.Customer
	id := r.FormValue("customer-id")

	customer.Id = id
	fmt.Println(w, customer.Id)
	customer.Name = r.FormValue("customer-name")

	customer.Update()

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var customer data.Customer
	id := r.FormValue("customer-id")
	customer.Id = id

	data.DeleteCustomer(id)

	http.Redirect(w, r, "/show-customer", http.StatusMovedPermanently)
}

func ShowCustomer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/customer.html")

	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	id := vars["id"]

	customer := data.GetCustomer(id)

	t.Execute(w, customer)
}

func ShowCustomers(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/dashboard.html")

	if err != nil {
		panic(err)
	}
	data := data.GetCustomersDetails()

	t.Execute(w, data)
}
