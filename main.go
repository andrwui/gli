package main

import (
	"fmt"

	selectableList "github.com/andrwui/gli/selectableList"
)

type Person struct {
	Name     string
	LastName string
	Birthday string
}

func main() {

	l := selectableList.NewSelectableList[Person]()

	addElements := func(l *selectableList.SelectableList[Person]) {
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

	printD := func(l *selectableList.SelectableList[Person]) {
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
