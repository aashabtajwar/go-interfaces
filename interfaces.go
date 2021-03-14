// second demonstration for interfaces
package main

import "fmt"

// declaring an interface
type Shape interface {
	/*
	* Declaring a "Shape" interface which has two methods: Area, and Perimeter
	* Both of them accept no argument and return float64 value.
	* Any type that implements these methods (with exact method signatures)
	* will also implement "Shape" interface
	 */
	Area() float64
	Perimeter() float64
}
type Rect struct {
	// creating Rect struct
	width  float64
	height float64
}

/*
* Defining methods Area and Perimeter which belongs to Rect type
* therefore Rect implemented those methods
* Since these methods are defined by the Shape interface,
* the Rect struct type implements the Shape interface.
* Since we haven't forced Rect to implement the Shape interface,
* it is all happening automatically
* So it is said that interfaces in Go are implicitely implemented
 */

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// interface is a type, just like structs
// so we can create a variable of its type
// in the above case, we can create a variable "s" of type interface "Shape"
func main() {
	var s Shape
	// When a type implements an interface, a variable of that type
	// can also be represented as the type of an interface
	// We can confirm by creating 'nil' interface 's' of type "Shape"
	// and assign a struct of type Rect
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Println("Value of s is", s)
	fmt.Printf("type of s is %T\n", s)
	fmt.Println("s == r is", s == r)
}

// %T - a Go-syntax representation of the type of the value
