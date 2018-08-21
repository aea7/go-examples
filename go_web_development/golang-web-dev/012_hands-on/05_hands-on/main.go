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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.html"))
}

func main() {

	restaurant := Menu{
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
	}

	new_file, err := os.Create("new_file")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer new_file.Close()
	err = tpl.Execute(new_file, restaurant)
}
