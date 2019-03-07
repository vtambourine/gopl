package main

import (
	"bufio"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type chapter struct {
	Number int
	Title  string
}

type excercise struct {
	Chapter *chapter
	Number  int
	Source  string
}

type byNumber []excercise

var chRe = regexp.MustCompile(`\d+`)
var exRe = regexp.MustCompile(`\w+\d+\.(\d+)`)

var chIndex []string

func (a byNumber) Len() int {
	return len(a)
}

func (a byNumber) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byNumber) Less(i, j int) bool {
	return a[i].Number < a[j].Number
}

func readChapter(dir os.FileInfo) ([]excercise, error) {
	chNum, err := strconv.Atoi(chRe.FindString(dir.Name()))
	excercises := []excercise{}

	if err != nil {
		return excercises, err
	}

	ch := chapter{
		Number: chNum,
		Title:  chIndex[chNum-1],
	}

	exFiles, err := ioutil.ReadDir(dir.Name())

	if err != nil {
		return excercises, err
	}

	for _, ef := range exFiles {
		if ef.IsDir() {
			// ex, err := readExcercise(ef)

			efName := ef.Name()
			efNum := exRe.FindStringSubmatch(efName)

			if len(efNum) == 2 {
				efNumInt, _ := strconv.Atoi(efNum[1])

				ex := excercise{
					&ch,
					efNumInt,
					path.Join(dir.Name(), ef.Name()),
				}

				excercises = append(excercises, ex)
			}
		}
	}

	return excercises, nil
}

func main() {
	const readme = `
{{- define "chapter" -}}
### Chapter {{ .Chapter.Number }}: {{ .Chapter.Title }}
{{- end}}
{{- define "excersise" -}}
[Excersise {{ .Chapter.Number }}.{{ .Number }}]({{ .Source }}){{ " " }}
{{- end}}
{{- $c := 0 -}}
# The Go Programming Language
Coding notes on [The Go Programming Language](http://www.gopl.io) book.
{{- range $k, $v := .}}
{{- if gt $v.Chapter.Number $c}}
{{- $c = $v.Chapter.Number}}

{{template "chapter" $v}}
{{end -}}
{{- template "excersise" . -}} 
{{end -}}
	`
	// {{if $v.Chapter.Number gt $c}}{{$c = $v.Chapter.Number}}{{end}}

	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	chpaterTitles, err := os.Open("helper/chapters.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer chpaterTitles.Close()

	scanner := bufio.NewScanner(chpaterTitles)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		chIndex = append(chIndex, scanner.Text())
	}

	excercises := []excercise{}

	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), "ch") {
			chExcs, err := readChapter(f)
			if err != nil {
				log.Fatal(err)
			}
			excercises = append(excercises, chExcs...)
		}
	}

	t := template.Must(template.New("readme").Parse(readme))

	sort.Sort(byNumber(excercises))

	// err = t.Execute(os.Stdout, excercises)
	// if err != nil {
	// 	log.Println("template:", err)
	// }

	err = t.ExecuteTemplate(os.Stdout, "readme", excercises)
	if err != nil {
		log.Println("template:", err)
	}
}
