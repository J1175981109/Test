package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./Test/student.txt" )
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}