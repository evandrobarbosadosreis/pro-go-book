package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAttend bool
}

//Make -> Initializes the slice defining the initial size and capacity
var responses = make([]*Rsvp, 0, 10)

func main() {
	fmt.Println("TODO: add some features");
}
