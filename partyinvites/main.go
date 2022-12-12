package main

import (
	"fmt"
	"html/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// Make -> Initializes the slice defining the initial size and capacity
var responses = make([]*Rsvp, 0, 10)

var templates = make(map[string]*template.Template, 3)

func loadTemaplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for index, name := range templateNames {

		t, err := template.ParseFiles("layout.html", name+".html")

		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}

}

func main() {
	loadTemaplates()
}
