package handlers

import (
	"data_manager/cust"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func AddCommercial(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	commercial := cust.Commercial{
		Id:         0,
		Last_name:  r.FormValue("last-name"),
		First_name: r.FormValue("first-name"),
		Role:       r.FormValue("role"),
		Address:    r.FormValue("address"),
	}

	commercial.Add()

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func AddProject(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	project := cust.Project{
		Id:            0,
		Name:          r.FormValue("project-name"),
		Description:   r.FormValue("project-description"),
		Customer_id:   r.FormValue("customer-id"),
		Commercial_id: r.FormValue("commercial-id"),
		Service_id:    r.FormValue("service-id"),
	}

	project.Add()

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
func AddCustomer(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	customer := cust.Customer{
		Id:            0,
		Logo:          r.FormValue("customer-logo"),
		Name:          r.FormValue("customer-name"),
		Address:       r.FormValue("customer-address"),
		Contact:       r.FormValue("customer-contact"),
		Branch:        r.FormValue("customer-branch"),
		Commercial_id: r.FormValue("commercial-id"),
	}
	customer.Add()

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func AddService(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	service := cust.Service{
		Id:   0,
		Name: r.FormValue("service-name"),
	}

	service.Add()

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user := cust.User{
		Id:            0,
		Email:         r.FormValue("user-email"),
		Password:      r.FormValue("user-password"),
		Commercial_id: r.FormValue("commercial-id"),
	}

	user.Add()

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("user-id"))), 10, 64)
	id := int(id64)

	cust.DeleteUser(id)
}

func DeleteCommercial(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	id := int(id64)

	cust.DeleteCommercial(id)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)
	id := int(id64)

	cust.DeleteService(id)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("project-id"))), 10, 64)
	id := int(id64)

	cust.DeleteProject(id)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("customer-id"))), 10, 64)
	id := int(id64)

	cust.DeleteCustomer(id)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("customer-id"))), 10, 64)
	id := int(id64)

	customer := cust.Customer{
		Id:            id,
		Logo:          r.FormValue("customer-logo"),
		Name:          r.FormValue("customer-name"),
		Address:       r.FormValue("customer-address"),
		Contact:       r.FormValue("customer-contact"),
		Branch:        r.FormValue("customer-branch"),
		Commercial_id: r.FormValue("commercial-id"),
	}

	customer.Update()

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("user-id"))), 10, 64)
	id := int(id64)

	user := cust.User{
		Id:       id,
		Email:    r.FormValue("user-email"),
		Password: r.FormValue("user-password"),
	}

	user.Update()

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("project-id"))), 10, 64)
	id := int(id64)

	project := cust.Project{
		Id:            id,
		Name:          r.FormValue("project-name"),
		Description:   r.FormValue("project-description"),
		Customer_id:   r.FormValue("customer-id"),
		Commercial_id: r.FormValue("commercial-id"),
		Service_id:    r.FormValue("service-id"),
	}

	fmt.Fprintln(w, project, id)
	//project.Update()

}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("service-id"))), 10, 64)
	id := int(id64)

	service := cust.Service{
		Id:   id,
		Name: r.FormValue("service-name"),
	}

	service.Update()

}

func UpdateCommercial(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id64, _ := strconv.ParseInt(strings.TrimSpace((r.FormValue("commercial-id"))), 10, 64)
	id := int(id64)
	fmt.Print(id)

	commercial := cust.Commercial{
		Id:         id,
		Last_name:  r.FormValue("commercial-last-name"),
		First_name: r.FormValue("commercial-first-name"),
		Role:       r.FormValue("commercial-role"),
		Address:    r.FormValue("commercial-address"),
	}

	commercial.Update()

}
