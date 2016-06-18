package main

import "os"

func main() {
	if err := os.Mkdir("hoge", os.ModeDir|0755); err != nil {
		panic(err)
	}
}
