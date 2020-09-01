package main

import (
	"fmt" 
	"math")

// Shape Interface Definition
type Shape interface{
	area() float64
}

// Circle Struct Definition
type Circle struct{
	x,y,radius float64

}

// Rectangle Struct Definition
type Rectangle struct{
	length,breadth float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64{
	return r.length * r.breadth
}

func getArea(s Shape) float64{
	return s.area()
}



func main() {
	circle := Circle{x:0,y:0,radius:5}
	rectangle := Rectangle{length:10,breadth:20}
	fmt.Printf("Circle Area : %f\n",getArea(circle))
	fmt.Printf("Rectangle Area : %f\n",getArea(rectangle))
}
