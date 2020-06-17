package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type content struct {
	Content string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

func writeFile(name string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func renderTemplate(filename string, data string) {

	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	var err error
	err = t.Execute(os.Stdout, content{Content: data})
	if err != nil {
		panic(err)
	}
}
func filterInput(input string) string {
	ext := ".html"
	return strings.Split(input, ".")[0] + ext

}

func writeTemplateToFile(templateName string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	filter := filterInput(data)
	f, err := os.Create(filter)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, content{Content: readFile(data)})
	if err != nil {
		panic(err)
	}

}

func main() {
	fileParse := flag.String("file", "", "txt file will be converted to html file")
	flag.Parse()
	if *fileParse != "" {
		renderTemplate("template.tmpl", readFile(*fileParse))
		writeTemplateToFile("template.tmpl", *fileParse)
	} else {
		renderTemplate("template.tmpl", readFile("first-post.txt"))
		writeTemplateToFile("template.tmpl", "test3.html")
	}
}
