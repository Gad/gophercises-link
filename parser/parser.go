package parser

import (
	"io"
	"strings"
	"golang.org/x/net/html"
)

type Link struct{
	href string
	text string
}

// accept an html doc and yield a slice of {href, text}
// for each <a href="href"> text </a>
func Parse(r io.Reader) ([]Link, error) {

	links := make([]Link,0)

	treeTop, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	//not an efficient implementation
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
	return links,nil

}