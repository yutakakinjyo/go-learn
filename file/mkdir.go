package main

import "os"
import "fmt"

func main() {
	if err := os.Mkdir("hoge", os.ModeDir|0755); err != nil {
		fmt.Println(err)
	}
}
