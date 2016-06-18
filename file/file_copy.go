package main

import "os"
import "io"
import "fmt"

func main() {

	from_file_name := "from"
	to_file_name := "to"

	from_f, err := os.Open(from_file_name)
	if err != nil {
		panic(err)
	}
	defer from_f.Close()

	if _, err := os.Stat(to_file_name); err == nil {
		fmt.Println(to_file_name + " is alreadly file exist")
		return
	}

	to_f, err := os.Create(to_file_name)
	if err != nil {
		panic(err)
	}
	defer to_f.Close()

	_, err = io.Copy(to_f, from_f)
	if err != nil {
		panic(err)
	}
}
