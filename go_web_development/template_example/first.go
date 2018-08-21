package template_example

import "fmt"

type human interface {
	speak()

}

type person struct {
	fname string
	lname string
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("ananen")
}


func main(){
	p1 := person {
		"anan",
		"baban",
	}

	saySomething(p1)
	p1.speak()

}