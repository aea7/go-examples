package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Employee struct {
	id 	   int
	name   string
	score  int
	salary float32
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=aak password=yo dbname=company sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	employees := make([]Employee, 0)
	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(&employee.id, &employee.name, &employee.score, &employee.salary) // order matters
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, employee := range employees {
		// fmt.Println(employee.isbn, employee.title, employee.author, employee.price)
		fmt.Printf("%d, %s, %d, $%.2f\n", employee.id, employee.name, employee.score, employee.salary)
	}
}
