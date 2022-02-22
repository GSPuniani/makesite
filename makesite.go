package main

import (
	"fmt"
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


	inputDir := flag.String("dir", " ", "Directory with which to parse all .txt files")
 	flag.Parse()

 	if *inputDir != " " {
		fmt.Println("Converting all txt files in " + *inputDir + "/")
		files, _ := ioutil.ReadDir(*inputDir)

		for _, file := range files {
			if !file.IsDir() {
				fileNameSplit := strings.Split(file.Name(), ".")
				if fileNameSplit[1] == "txt" {
					outputName := strings.Split(file.Name(), ".")[0] + ".html"
					fileData, _ := ioutil.ReadFile(file.Name())

					myStruct := Page{MyData: string(fileData)}

					// Use a defined template
					parsedTemplate, _ := template.ParseFiles("template.tmpl")

					// Create a file to write to
					newFile, _ := os.Create(outputName)

					// Write to new file using template and data
					err := parsedTemplate.Execute(newFile, myStruct)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	} else {
		fmt.Println("Empty directory argument. No HTML files generated.")
	}

}
