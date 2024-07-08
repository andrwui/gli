package main

import (
	"fmt"

	list "github.com/andrwui/gli/list"
)

type Person struct {
	Name     string
	LastName string
	Birthday string
}

func main() {

	l := list.NewList[Person]()

	addElements := func(l *list.List[Person]) {
		var name string
		var lastName string
		var birthday string

		fmt.Println("Enter a name:")
		fmt.Scanln(&name)
		l.EraseLines(2)
		fmt.Println("Enter a last name:")
		fmt.Scanln(&lastName)
		l.EraseLines(2)
		fmt.Println("Enter a birthday:")
		fmt.Scanln(&birthday)
		l.EraseLines(2)

		a := &Person{
			Name:     name,
			LastName: lastName,
			Birthday: birthday,
		}
		l.AddItem(a)
	}

	printD := func(l *list.List[Person]) {
		l.Exit()
	}

	l.ShowFieldHeaders(true)
	l.SetFieldSize(30)

	l.ShowStructField("Name")
	l.ShowStructField("LastName")
	l.ShowStructField("Birthday")

	l.OnKey('a', addElements)
	l.OnKey('d', printD)

	l.Display()

}
