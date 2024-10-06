package handlers

import (
	"data_manager/database"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	unix := time.Now().Unix()
	fmt.Fprintln(w, unix)
}

func loggingHandler(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("views/home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Execute(w, make(map[int]string))
}

func AddNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var site database.Site

	if r.FormValue("site-id") != "" {
		id := r.FormValue("site-id")

		site = database.GetSite(id)

	} else {
		sites := database.GetSites()
		site = sites[len(sites)-1]
	}

	number_id := fmt.Sprint(time.Now().Unix())
	customer_number := r.FormValue("customer-number")

	number := database.Number{
		Id:      number_id,
		Number:  customer_number,
		Site_id: site.Id,
	}

	number.Add()
}

func AddSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var customer database.Customer

	if r.FormValue("customer-id") != "" {
		id := r.FormValue("customer-id")

		customer = database.GetCustomer(id)

	} else {
		customers := database.GetCustomers()
		customer = customers[len(customers)-1]

	}

	id := fmt.Sprint(time.Now().Unix())
	name := r.FormValue("site-name")

	site := database.Site{
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

	var number database.Number

	if r.FormValue("number-id") != "" {
		number = database.GetNumber(r.FormValue("number-id"))
	} else {
		numbers := database.GetNumbers()
		number = numbers[len(numbers)-1]
	}

	id := fmt.Sprint(time.Now().Unix())

	volume64, _ := strconv.ParseInt(r.FormValue("data"), 10, 64)
	volume := int(volume64)
	date_re, _ := time.Parse("2006-01-02", r.FormValue("date-re"))
	date_exp := date_re.AddDate(0, 1, 1)
	auto_re, _ := strconv.ParseBool(r.FormValue("auto-re"))

	recharge := database.Recharge{
		Id:        id,
		Volume:    volume,
		Date_re:   date_re,
		Date_exp:  date_exp,
		Auto_re:   auto_re,
		Number_id: number.Id,
	}

	recharge.Add()
}

func UpdateRecharge(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var recharge database.Recharge
	recharge.Id = r.FormValue("recharge-id")
	v, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	recharge.Volume = int(v)
	recharge.Date_re, _ = time.Parse("2006-01-02", r.FormValue("date-re"))
	recharge.Date_exp = recharge.Date_re.AddDate(0, 1, 1)

	recharge.Update()
}

func UpdateNumber(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var number database.Number
	number.Id = r.FormValue("number-id")
	number.Number = r.FormValue("customer-number")

	number.Update()
}

func UpdateSite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var site database.Site
	site.Id = r.FormValue("site-id")
	site.Name = r.FormValue("site-name")

	site.Update()
}
