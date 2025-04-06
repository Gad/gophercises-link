package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	path := flag.String("path", "ex1.html", "file to parse")
	flag.Parse()
	r, err := os.Open(*path)
	if errors.Is(err, os.ErrNotExist) {
		log.Panicln("file not exists error")
	}

	defer r.Close()
	type Link struct{
		href string
		text string
	}

	links := make([]Link,0)

	treeTop, _ := html.Parse(r)
	for node := range treeTop.Descendants() {
		link :=Link{}
		if node.Type == html.ElementNode && node.Data == "a" {
			link.href = node.Attr[0].Val
			for node1 := range node.Descendants() {
				if node1.Type == html.TextNode {
					link.text += strings.TrimSpace(node1.Data)
				}
			}
			links = append(links, link)

		}

	}

	fmt.Printf("%+v\n", links)
	/* t, err := io.ReadAll(r)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(t)) */
}
