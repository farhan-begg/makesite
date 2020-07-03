package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"github.com/fatih/color"

)

type content struct {
	Content string
}


func readFile(templateName string) string {
	fileContents, err := ioutil.ReadFile(templateName)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
        // value that we don’t know how to (or want to) handle. This example
        // panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	return string(fileContents)
}



func renderTemplate(filename string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , content{Content: data})
	if err != nil {
		panic(err)
	}
}



func addExtHTML(filename string) string {
	ext := ".html"
	return strings.Split(filename, ".")[0] + ext
}



func writeTemplateToFile(tmplName string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(tmplName))

	file := addExtHTML(data)
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, content{Content: readFile(data)})
	if err != nil {
		panic(err)
	}
}



func isTxtFile(filename string) bool {
	if strings.Contains(filename, ".") {
		return strings.Split(filename, ".")[1] == "txt"
	}
	return false
}



func main() {

	filePtr := flag.String("file", "", "name of txt file to be converted to html file")
	dirPtr := flag.String("dir", "", "name of directory to search")
	
	flag.Parse()
	if *dirPtr != "" {
		files, err := ioutil.ReadDir(*dirPtr)
		if err != nil{
			panic(err)
		}

		for _, file := range files {
			name := file.Name()
			if isTxtFile(name) == true {
				renderTemplate("template.tmpl", readFile(name))
				writeTemplateToFile("template.tmpl", name)
				fmt.Println(file.Name())

				minion := color.New(color.FgBlack).Add(color.BgYellow).Add(color.Bold)
				minion.Println("Minion says: banana!!!!!!")
			 
				m := minion.PrintlnFunc()
				m("I want another banana!!!!!")
			 
				slantedRed := color.New(color.FgRed, color.BgWhite, color.Italic).SprintFunc()
				fmt.Println("I've made a huge", slantedRed("mistake"))
		
			}
		}
	}

	if *filePtr != "" {
		renderTemplate("template.tmpl", readFile(*filePtr))
		writeTemplateToFile("template.tmpl", *filePtr)
	}
}