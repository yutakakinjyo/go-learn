package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v *Vertex) add() int {
	return v.x + v.y
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v.add())

}
