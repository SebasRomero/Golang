package main

import "fmt"

type Point struct { //Creating the type with their values
	x        int32
	y        int32
	isOnGrid bool
}

type Circle struct {
	radius float64
	center *Point
}

func changeX(pt *Point) { //This function change the real value using the pointer in the RAM
	pt.x = 100
}

func main() {
	var p1 Point = Point{1, 2, true} //Way to create a Point var type
	var p2 Point = Point{-5, 7, true}
	p3 := Point{9, 5, false} //Another way
	p4 := &Point{x: 12}      //Another way, also using the pointer sign '&'
	fmt.Println(p1.x, p2.x)

	p1.x = 21
	fmt.Println(p1.x, (p2.isOnGrid == false), p3.x, p4) //Calling the differents variables, look the p4, that symbol isn't using it
	fmt.Println("-----")

	changeX(p4)
	fmt.Println(p4) //The value was changed, because we used the real (selecting the pointer)

	(*p4).x = 99 // Here we change the value without use some function, just directly
	fmt.Println(p4)
	p4.x = 33
	fmt.Println(p4) // It doesn't matter if we use the * like we realized
	fmt.Println("Circle: ")
	c1 := Circle{4.41, &(p1)}
	fmt.Println(c1.center.x)
}
