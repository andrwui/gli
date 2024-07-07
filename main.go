package main

import (
	list "github.com/andrwui/gli/list"
)

type object struct {
	Name     int
	Surname  bool
	Birthday float32
}

func main() {

	l := list.NewList[object]()

	addElements := func(l *list.List[object]) {
		l.AddItem(&object{
			Name:     1,
			Surname:  false,
			Birthday: float32(1),
		})
	}

	printD := func(l *list.List[object]) {
		l.Exit()
	}

	l.ShowFieldHeaders(true)
	l.SetFieldSize(30)

	l.ShowStructField("Name")
	l.ShowStructField("Surname")
	l.ShowStructField("Birthday")

	l.OnKey('a', addElements)
	l.OnKey('d', printD)

	l.Display()

}
