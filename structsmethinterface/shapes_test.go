package structsmethinterface

import "testing"

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want{
		t.Errorf("got %.2f want %.2f", got, want) //f = float64 .2 = 2 decimal places
	}
}

func TestArea(t *testing.T){
	//"table driven tests". This creates a list of shapes and their wants as a single struct.
	// Structured kind of like a table that we can loop through
	areaTests := []struct {
		name string 
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12,Height:6}, hasArea:72.0}, //we can name fields here to make them easy to read
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12,Height: 6},hasArea: 36.0 },
	}

	for _, tt := range areaTests{
		t.Run(tt.name, func(t *testing.T){ // t.Run will allow us to print the name of the test if one fails
			got := tt.shape.Area()
			if got != tt.hasArea{
				t.Errorf("%#v got %g want %g",tt.shape, got, tt.hasArea) // %#v lets us print the struct for more information.
			}
		})
	}



}
//*** These are the original tests. I placed them here to create table driven tests seen above ^

	// func TestArea(t *testing.T){

	// 	checkArea := func(t testing.TB, shape Shape, want float64){
	// 		t.Helper()
	// 		got := shape.Area()
	// 		if got != want {
	// 			t.Errorf("got %g want %g", got, want)
	// 		}
	// 	}
	// 	t.Run("rectangles", func(t *testing.T){
	// 		rectangle := Rectangle{12.0, 6.0}
	// 		checkArea(t, rectangle, 72.0)
	// 	})

	// 	t.Run("Circles", func(t *testing.T){
	// 		circle := Circle{10}
	// 		checkArea(t, circle, 314.1592653589793)
	// 	})

