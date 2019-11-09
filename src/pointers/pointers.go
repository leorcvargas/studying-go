package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name     string
	health   int
	position position
}

func addOne(num *int) {
	*num++
}

func whereIsBadGuy(guy *badGuy) {
	x := guy.position.x
	y := guy.position.y
	fmt.Printf("( x: %f, y: %f2 )", x, y)
}

func main() {
	x := 5
	xPointer := &x

	fmt.Println("x value: ", x)
	fmt.Println("x address: ", xPointer)

	addOne(xPointer)
	fmt.Println("x value after addOne: ", x)

	thanos := badGuy{"Thanos", 958, position{10.5, 90}}
	whereIsBadGuy(&thanos)
}
