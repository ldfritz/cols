package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: cols FILENAME [COLUMNS ...]")
		return
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	c := csv.NewReader(f)
	records, err := c.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 {
		fmt.Println(strings.Join(records[0], "\n"))
		return
	}

	columns := make(map[string]int)
	for i, v := range records[0] {
		columns[v] = i
	}

	var notFound []string
	var header []string
	var indexes []int
	for _, name := range os.Args[2:] {
		index, ok := columns[name]
		if !ok {
			notFound = append(notFound, name)
		}
		header = append(header, name)
		indexes = append(indexes, index)
	}

	if len(notFound) != 0 {
		fmt.Println("column(s) not found:", strings.Join(notFound, "\n"))
		fmt.Println("")
		fmt.Println(strings.Join(records[0], "\n"))
		return
	}

	out := csv.NewWriter(os.Stdout)
	if err = out.Write(header); err != nil {
		log.Fatal(err)
	}

	for _, r := range records[1:] {
		var row []string
		for _, i := range indexes {
			row = append(row, r[i])
		}
		if err = out.Write(row); err != nil {
			log.Fatal(err)
		}
	}

	out.Flush()

	if err = out.Error(); err != nil {
		log.Fatal(err)
	}
}
