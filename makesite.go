package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	firstPostContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	type Page struct {
		
	}
	fmt.Println("Hello, world!")
}
