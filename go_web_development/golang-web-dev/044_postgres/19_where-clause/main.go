package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"encoding/json"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=aak password=ara port=5432 sslmode=disable dbname=company")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

type Employee struct {
	id 	   int
	name   string
	score  int
	salary float32
}

func main() {
	http.HandleFunc("/employees", employees)
	http.HandleFunc("/employees/show", employeesShow)
	http.ListenAndServe(":8080", nil)
}

func employees(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	employees := make([]Employee, 0)
	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(&employee.id, &employee.name, &employee.score, &employee.salary) // order matters
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, employee := range employees {
		fmt.Fprintf(w, "%d, %s, %d, $%.2f\n", employee.id, employee.name, employee.score, employee.salary)
	}
}

func employeesShow(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	fmt.Println(id)
	decoder := json.NewDecoder(r.Body)

	t := struct {
		Id int
	}{}

	a := decoder.Decode(&t)

	if a != nil{
		panic(a)
	}
	fmt.Println(t.Id)
	if t.Id == -1 {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM employees WHERE id = $1", t.Id)

	employee := Employee{}
	err := row.Scan(&employee.id, &employee.name, &employee.score, &employee.salary)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%d, %s, %d, $%.2f\n", employee.id, employee.name, employee.score, employee.salary)
}

