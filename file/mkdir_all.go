package main

import "os"
import "fmt"

func main() {
	if err := os.MkdirAll("hoge/fuga", os.ModeDir|0755); err != nil {
		if !os.IsExist(err) {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
