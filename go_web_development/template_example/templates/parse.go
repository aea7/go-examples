package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil{
		log.Fatalln(err)
	}


	new_file, err := os.Create("new_file")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer new_file.Close()

	template, err = template.ParseFiles("new_file", "templates/index.html")

	sages := []string{"MLK", "JEZOS"}
	err = template.ExecuteTemplate(new_file, "index.html", sages)

	template, err = template.ParseGlob("templates/*")
	template, err = template.ParseGlob("templates/*.html")

}
