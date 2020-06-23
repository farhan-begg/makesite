package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type content struct {
	Content string
}

// reads filename and returns the contents
func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

// writes data to a file named by filename.
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

func parser() {
	var dir string
	flag.StringVar(&dir, "dir", "", "this is the directory")
	flag.Parse()

	fmt.Println("Directory:", dir)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filenameCheck(file.Name()) == true {
			fmt.Println(file.Name())
			writeTemplateToFile("template.tmpl", file.Name())
		}
	}
}

func filenameCheck(filename string) bool {
	tail := "txt"
	for i := range filename {
		if filename[i] == '.' {
			s := strings.Split(filename, ".")[1]
			// fmt.Println(s)
			if s == tail {
				return true
			}
		}
	}
	return false
}

func main() {
	// arg := os.Args[1]
	parser()
	// renderTemplate("template.tmpl", readFile(arg))
	// writeTemplateToFile("template.tmpl", arg)
}
