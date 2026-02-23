package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"ruley-linter/internal/config"
	"ruley-linter/internal/validator"
)

var validators []validator.TagValidator

func traverse(n *html.Node, depth int) {
	switch n.Type {
	case html.ElementNode:
		if depth > 5 {
			fmt.Printf("Depth level exceeded. Max depth level is 5, current depth level => %v\n", depth)
		}
		for i := range validators {
			if validators[i].ApplicableTags[n.Data] {
				isValid := validators[i].ValidationFunc(n)

				if !isValid {
					fmt.Printf("[Validation errror] Tag <%s>: %s\n", n.Data, validators[i].ErrorMessage)
				}
			}
		}

	case html.TextNode:
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, depth+1)
	}
}

func main() {
	cfg, err := config.Load(".ruley.config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading config: %v\n", err)
		os.Exit(1)
	}

	validators, err = cfg.BuildValidators()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while building validators: %v\n", err)
		os.Exit(1)
	}

	htmlFileContent, err := os.ReadFile("./test-files/valid.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading html file: %v\n", err)
		os.Exit(1)
	}

	htmlAST, err := html.Parse(strings.NewReader(string(htmlFileContent)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while linting htmlAST: %v\n", err)
		os.Exit(1)
	}

	traverse(htmlAST.FirstChild.FirstChild.NextSibling.NextSibling, 0)
}
