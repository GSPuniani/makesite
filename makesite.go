package main

import (
	"fmt"
	"io/ioutil"
	"html/template"
	"os"
)

type Page struct {
	MyData string
}


func main() {
	firstPostContents, err := ioutil.ReadFile("first-post.txt")
	fmt.Println(string(firstPostContents))
	my_page := Page{MyData: string(firstPostContents)}
	fmt.Println(my_page.MyData)
	if err != nil {
		panic(err)
	}

	// t := template.Must(template.New("template.tmpl").ParseFiles("first-post.html"))
	template_parse, err := template.ParseFiles("template.tmpl")
	html_write, _ := os.Create("first-post.html")
	err = template_parse.Execute(html_write, my_page)
	// err = t.Execute(os.Stdout, my_page)
	if err != nil {
		panic(err)
	}
	

}
