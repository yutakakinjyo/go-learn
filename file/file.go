package main

import "os"
import "fmt"

func main() {

	if _, err := os.Stat("hoge/fuga/piyo"); err == nil {
		fmt.Println("already file is exist")
		return
	}

	f, err := os.Create("hoge/fuga/piyo")
	if err != nil {
		panic(err)
	}

	defer f.Close()

}
