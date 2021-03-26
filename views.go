package main

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Views struct {
	dtbs *sql.DB
}

func (v *Views) GetCustomer(CustomerID int) (Customer, error) {

	res := Customer{}

	var id int64
	var firstname string
	var lastname string
	var birthdate pq.NullTime
	var gender string
	var email string
	var homeaddress string

	err := v.dtbs.QueryRow(`SELECT id, firstname, lastname, birthdate, gender, email, homeaddress FROM customers where id = $1`, CustomerID).Scan(&id, &firstname, &lastname, &birthdate, &gender, &email, &homeaddress)
	if err == nil {
		res = Customer{ID: id, FirstName: firstname, LastName: lastname, BirthDate: birthdate.Time, Gender: gender, Email: email, Address: homeaddress}
	}

	return res, err
}

func (v *Views) AllCustomers() ([]Customer, error) {

	customers := []Customer{}

	rows, err := v.dtbs.Query(`SELECT id, firstname, lastname, birthdate, gender, email, homeaddress FROM customers order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var firstname string
		var lastname string
		var birthdate pq.NullTime
		var gender string
		var email string
		var homeaddress string

		err = rows.Scan(&id, &firstname, &lastname, &birthdate, &gender, &email, &homeaddress)
		if err != nil {
			return customers, err
		}

		currentCustomer := Customer{ID: id, FirstName: firstname, LastName: lastname, BirthDate: birthdate.Time, Gender: gender, Email: email, Address: homeaddress}
		if birthdate.Valid {
			currentCustomer.BirthDate = birthdate.Time
		}

		customers = append(customers, currentCustomer)
	}

	return customers, err
}

func (v *Views) CreateCustomer(firstname string, lastname string, birthdate time.Time, gender string, email string, homeaddress string) (int, error) {

	var customerID int
	err := v.dtbs.QueryRow(`INSERT INTO customers(firstname, lastname, birthdate, gender, email, homeaddress) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, firstname, lastname, birthdate, gender, email, homeaddress).Scan(&customerID)

	if err != nil {
		return 0, err
	}

	return customerID, err
}

func (v *Views) SearchCustomers(name string) ([]Customer, error) {

	customers := []Customer{}

	rows, err := v.dtbs.Query(`SELECT id, firstname, lastname, birthdate, gender, email, homeaddress FROM customers WHERE firstname=$1 OR lastname=$1`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var firstname string
		var lastname string
		var birthdate pq.NullTime
		var gender string
		var email string
		var homeaddress string

		err = rows.Scan(&id, &firstname, &lastname, &birthdate, &gender, &email, &homeaddress)
		if err != nil {
			return customers, err
		}

		currentCustomer := Customer{ID: id, FirstName: firstname, LastName: lastname, Gender: gender, Email: email, Address: homeaddress}
		if birthdate.Valid {
			currentCustomer.BirthDate = birthdate.Time
		}

		customers = append(customers, currentCustomer)
	}

	return customers, err
}

func (v *Views) EditCustomer(id int, firstname string, lastname string, birthdate time.Time, gender string, email string, homeaddress string) (int, error) {

	res, err := v.dtbs.Exec(`UPDATE customers set firstname=$1, lastname=$2, birthdate=$3, gender=$4, email=$5, homeaddress=$6 where id=$7 RETURNING id`, firstname, lastname, birthdate, gender, email, homeaddress, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func (v *Views) DeleteCustomer(customerID int) (int, error) {

	res, err := v.dtbs.Exec(`delete from customers where id = $1`, customerID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
