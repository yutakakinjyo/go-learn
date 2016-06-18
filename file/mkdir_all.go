package main

import "os"

func main() {
	if err := os.MkdirAll("hoge/fuga", os.ModeDir|0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
}
