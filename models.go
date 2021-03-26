package main

import (
	"time"
)

type IndexPage struct {
	AllCustomers []Customer
}

type CustomerPage struct {
	TargetCustomer Customer
}

type Customer struct {
	ID        int64
	FirstName string
	LastName  string
	BirthDate time.Time
	Gender    string
	Email     string
	Address   string
}

func (c Customer) BirthDateStr() string {
	return c.BirthDate.Format("02 January 2006")
}

type ErrorPage struct {
	ErrorMsg string
}
