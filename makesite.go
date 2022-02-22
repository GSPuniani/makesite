package main

import (
	// "fmt"
	"io/ioutil"
	"html/template"
	"os"
	"flag"
	"strings"
)

type Page struct {
	MyData string
}


func main() {
	inputFile := flag.String("file", "first-post.txt", ".txt file to parse into HTML")
 	flag.Parse()

 	outputFile := strings.Split(*inputFile, ".")[0] + ".html"
 	fileData, _ := ioutil.ReadFile(*inputFile)

	my_page := Page{MyData: string(fileData)}

	template_parse, err := template.ParseFiles("template.tmpl")
	html_write, _ := os.Create(outputFile)
	err = template_parse.Execute(html_write, my_page)
	if err != nil {
		panic(err)
	}
	

}
