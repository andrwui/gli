package selectableList

import (
	"fmt"
	"reflect"

	cli "github.com/andrwui/gli/internal"
)

type SelectableList[T any] struct {
	// The position of the cursor in the list.
	//
	// Can be used to access the currently indicated item.
	CursorPos int

	// The slice of struct values used to render the list.
	//
	// Each item added to this slice will be listed.
	Items []*T

	// The slice of the selected struct values rendered in the list.
	SelectedItems []*T

	// The slice containing all the names for the fields that are going to be shown in the list
	shownFields []string

	// The spacing (in characters) bewtween the fields in the list
	fieldSpacing int

	// Controls if the header with the field names is shown.
	isHeaderShown bool

	// Controls if the user did exit the list.
	exited bool

	// Stores the function to execute when the user exits the list.
	onExitFunc func(cb *SelectableList[T])

	// Stores all the functions to execute when the user presses the predefined key.
	keyBindings map[byte]func(*SelectableList[T])
}

// Creates a new selectable list.
//
// A struct type must be specified to list the items.
func NewSelectableList[T any]() *SelectableList[T] {
	return &SelectableList[T]{
		CursorPos:   0,
		Items:       make([]*T, 0),
		keyBindings: make(map[byte]func(*SelectableList[T])),
	}
}

// Adds a struct value as an item to the list.
func (l *SelectableList[T]) AddItem(item *T) {
	l.Items = append(l.Items, item)
}

// Shows the struct field that has the entered name in the list.
func (l *SelectableList[T]) ShowStructField(propName string) {
	l.shownFields = append(l.shownFields, propName)
}

// Sets the minimum character lenght each field in the list will have.
func (l *SelectableList[T]) SetFieldSize(space int) {
	l.fieldSpacing = space
}

// Sets the visibility with the name of the fields at the top of the list.
func (l *SelectableList[T]) ShowFieldHeaders(t bool) {
	l.isHeaderShown = t
}

// Stores the callback function that will be called when the return key is pressed.
func (l *SelectableList[T]) OnReturn(cb func(l *SelectableList[T])) {
	l.keyBindings[13] = cb
}

// Stores the callback function that will be called when the entered key (rune) is pressed.
func (l *SelectableList[T]) OnKey(key rune, cb func(l *SelectableList[T])) {
	l.keyBindings[byte(key)] = cb
}

// Toggles the selection for the currently indicated item.
func (l *SelectableList[T]) ToggleSelected(item *T) {
	if l.IsItemSelected(item) {
		l.DeselectItem(item)
		l.render()
	} else {
		l.SelectItem(item)
		l.render()
	}
}

// Selects the indicated item.
func (l *SelectableList[T]) SelectItem(item *T) {
	l.SelectedItems = append(l.SelectedItems, item)
}

// Deselects the indicated item.
func (l *SelectableList[T]) DeselectItem(item *T) {
	for i, sItem := range l.SelectedItems {
		if sItem == item {
			l.SelectedItems = append(l.SelectedItems[:i], l.SelectedItems[i+1:]...)
		}
	}
}

// Returns if the passed item is selected on the list.
func (l *SelectableList[T]) IsItemSelected(item *T) bool {
	for _, sItem := range l.SelectedItems {
		if sItem == item {
			return true
		}
	}
	return false
}

// Executes the callback for the entered key.
func (l *SelectableList[T]) execKeybinding(b byte) {

	fmt.Print(b)

	if l.keyBindings[b] != nil {
		l.keyBindings[b](l)
	}
	l.render()
}

// Changes the exited state to true, therefore ending the instance of the list.
func (l *SelectableList[T]) Exit() {
	if l.onExitFunc != nil {
		l.onExitFunc(l)
	}
	l.exited = true
}

// Sets the function to be called when the application exits.
func (l *SelectableList[T]) OnExit(cb func(l *SelectableList[T])) {
	l.onExitFunc = cb
}

// Erases the entered quantity of lines from the bottom up.
func (l *SelectableList[T]) EraseLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[A")
		fmt.Print("\033[2K")
	}
	fmt.Print("\r")
}

// Displays the list, and checks the user input continuously to execute the respective functions.
func (l *SelectableList[T]) Display() {

	for {

		if l.exited {
			fmt.Print("\033[H\033[2J")
			return
		}

		if len(l.Items) < 1 {
			fmt.Printf("%-v\n", "   No items")

			keyPressed := cli.GetRawInput()

			if keyPressed == 27 || keyPressed == 3 {
				l.Exit()
			} else {
				l.execKeybinding(keyPressed)
			}

		} else {

			currentItem := l.Items[l.CursorPos]

			keyPressed := cli.GetRawInput()

			// fmt.Println(keyPressed)

			if keyPressed == 27 || keyPressed == 3 {
				l.Exit()
			} else if keyPressed == 65 || keyPressed == 66 {

				switch keyPressed {

				case 65:
					l.CursorPos = (l.CursorPos + len(l.Items) - 1) % len(l.Items)
					l.render()

				case 66:
					l.CursorPos = (l.CursorPos + 1) % len(l.Items)
					l.render()

				}

			} else if keyPressed == 32 {
				l.ToggleSelected(currentItem)
			} else {
				l.execKeybinding(keyPressed)
			}
		}
	}
}

// Function to refresh the screen and update the changes into the console.
func (l *SelectableList[T]) render() {
	fmt.Print("\033[H\033[2J")

	if l.isHeaderShown {
		fmt.Printf("   ")
		fmt.Printf("%-*s", l.fieldSpacing, "Selection")
		for _, propName := range l.shownFields {
			fmt.Printf("%-*s", l.fieldSpacing, propName)
		}
		fmt.Printf("\n")
	}

	for i, item := range l.Items {
		var cursor string

		if l.CursorPos == i {
			cursor = ">  "
		} else {
			cursor = "   "
		}

		var selection string
		if l.IsItemSelected(item) {
			selection = "[ ]"
		} else {
			selection = "[x]"
		}

		fmt.Printf("%s", cursor)
		fmt.Printf("%-*s", l.fieldSpacing, selection)

		for _, propName := range l.shownFields {
			itemValue := reflect.ValueOf(item).Elem()
			nameField := itemValue.FieldByName(propName)

			fmt.Printf("%-*v", l.fieldSpacing, nameField)
		}

		fmt.Print("\n")

	}

}
