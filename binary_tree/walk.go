package main

import "code.google.com/p/go-tour/tree"
import "fmt"

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for val1 := range ch1 {
		val2 := <-ch2
		if val1 != val2 {
			return false
		}
	}

	return true
}

func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

func _Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	_Walk(t.Left, ch)
	ch <- t.Value
	_Walk(t.Right, ch)
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Same!")
	} else {
		fmt.Println("not Same!")
	}
}
