package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	file := flag.String("file", "", "the file path")
	flag.Parse()
	content, err := readPdf(*file)
	if err != nil {
		fmt.Println(content)
		panic(err)
	}
}

func readPdf(path string) (string, error) {
	r, err := Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		fmt.Println(">>>>")
		fmt.Println("page: ", pageIndex)
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			for _, word := range row.Content {
				s := strings.TrimSpace(word.S)
				length := len([]rune(s))
				if length > 0 {
					fmt.Println(s)
				}
			}
		}
		fmt.Println("<<<<")
	}
	return "", nil
}
