package main

import (
	"data_manager/database"
	"reflect"
	"testing"
)

func Test_GetCustomers(t *testing.T) {
	var c []database.Customer
	customers := database.GetCustomers()

	if reflect.TypeOf(customers) != reflect.TypeOf(c) && len(customers) <= 0 {
		t.Error("Incorrect result: expected []Customers, got", reflect.TypeOf(customers))
	}
}
