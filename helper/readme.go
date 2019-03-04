package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
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

func (e byNumber) Len() int {
	return len(e)
}

func (e byNumber) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e byNumber) Less(i, j int) bool {
	return e[i].Number < e[j].Number
}

func readExcersices(dir os.FileInfo) ([]excercise, error) {
	chNum, err := strconv.Atoi(chRe.FindString(dir.Name()))

	if err != nil {
		return []excercise{}, err
	}

	ch := chapter{
		Number: chNum,
	}
	fmt.Println(ch)

	exFiles, err := ioutil.ReadDir(dir.Name())

	if err != nil {
		return []excercise{}, err
	}

	excercises := []excercise{}

	for _, ef := range exFiles {
		if ef.IsDir() {
			efName := ef.Name()
			efNum := exRe.FindStringSubmatch(efName)

			if len(efNum) == 2 {
				fmt.Printf("%v\n", efNum[1])

				efNumInt, _ := strconv.Atoi(efNum[1])

				ex := excercise{
					&ch,
					efNumInt,
					"",
				}

				excercises = append(excercises, ex)
			}

		}
	}

	return excercises, nil
}

func main() {
	const readme = `
# The Go
{{range .}}
ch {{ .Chapter.Number }} n {{ .Number }}
{{end}}
	`

	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	excercises := []excercise{}

	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), "ch") {
			chExcs, err := readExcersices(f)
			if err != nil {
				log.Fatal(err)
			}
			excercises = append(excercises, chExcs...)
		}
	}

	t := template.Must(template.New("readme").Parse(readme))

	sort.Sort(byNumber(excercises))

	err = t.Execute(os.Stdout, excercises)
	if err != nil {
		log.Println("template:", err)
	}
}
