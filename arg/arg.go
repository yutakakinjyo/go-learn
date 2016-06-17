package main

import "os"
import "fmt"

func main() {
	args := os.Args
	if len(args) >= 2 {
		fmt.Println(args[1])
	}
}
