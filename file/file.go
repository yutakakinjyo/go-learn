package main

import "os"
import "fmt"

func main() {

	if _, err := os.Stat("hoge/fuga/piyo"); err == nil {
		fmt.Println("already file is exist")
		os.Exit(0)
	}

	if _, err := os.Create("hoge/fuga/piyo"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
