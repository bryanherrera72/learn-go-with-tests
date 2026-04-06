package structsmethinterface

import "math"

// we now have a shape type with an Area function. (Methods to be exact)
// This is like "class Rectangle implements Shape(){}" in Java.
// Difference here is that it is implicit. If Rectangle has an Area function, it is inherently a shape.
// This decouples the helper, from the concrete types.
type Shape interface{
	Area() float64
}

// A Struct is like a data class / java POJO.
// only contains fields, and doesn't have getters/ setters unless needed. 
type Rectangle struct{
	Width float64
	Height float64
}

// Method = function + receiver.
// This is similar to creating static methods in java. 
// The value of "this" is the receiver, in this method it is r
// this is how you define methods on the struct itself.
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

type Circle struct{
	Radius float64
}

func(c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct{
	Base float64
	Height float64
}

func(t Triangle) Area() float64{
	return .5 * t.Base * t.Height
}

func Perimeter(rect Rectangle) float64{
	return 2 * (rect.Width + rect.Height)
}

// func Area(rect Rectangle) float64{
// 	return rect.Width * rect.Height
// }