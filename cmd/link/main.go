package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gad/gophercises-link/parser"
)

func main() {
	path := flag.String("path", "ex1.html", "path to file to parse")
	flag.Parse()

	// open the html file to parse
	r, err := os.Open(*path)
	if errors.Is(err, os.ErrNotExist) {
		log.Panicln("file does not exists error")
	}

	defer r.Close()
	
	links, err := parser.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)

}
