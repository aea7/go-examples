package main

import (
	"html/template"
	"os"
	"log"
)

//contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

type item struct {
	Price int
	Name  string
}
type Menu struct {
	Breakfast []item
	Lunch     []item
	Dinner    []item
}

type restaurant struct {
	Menu []Menu
	Name string
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.html"))
}

func main() {

	restaurants := []restaurant{
		{
			Name: "Ramadan Bingöl",
			Menu: []Menu{
				Menu{
					Breakfast: []item{
						item{
							50,
							"serpme",
						},
						item{
							10000,
							"paha bicilmez",
						},
					},
					Lunch: []item{
						item{
							100,
							"tavuk",
						},
					},
					Dinner: []item{
						item{
							150,
							"spago",
						},
					},
				},
			},
		},
		{
			Name: "Kahve Dehası",
			Menu: []Menu{
				Menu{
					Breakfast: []item{
						item{
							8,
							"apo",
						},
						item{
							10,
							"mocha",
						},
					},
				},
			},
		},
	}

	new_file, err := os.Create("new_file")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer new_file.Close()
	err = tpl.Execute(new_file, restaurants)
}
