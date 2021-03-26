package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your_password"
	dbname   = "customersdb"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt1 := `insert into "customers"("firstname", "lastname", "birthdate", "gender", "email", "homeaddress") values('Ivan', 'Ivanov', '01 January 1991', 'Male', 'ivan@gmail.com', 'AddressOfIvan')`
	_, e := db.Exec(insertStmt1)
	CheckError(e)

	insertDynStmt := `insert into "customers"("firstname", "lastname", "birthdate", "gender", "email", "homeaddress") values($1, $2, $3, $4, $5, $6)`
	_, e = db.Exec(insertDynStmt, "Petr", "Petrov", "01 January 1992", "Male", "petr@gmail.com", "AddressOfPetr")
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
