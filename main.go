package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"ruley-linter/internal/config"
	"ruley-linter/internal/validator"
)

var (
	tagValidators     []validator.TagValidator
	textNodeValidator validator.TextNodeValidator
)

func traverse(n *html.Node, cfg *config.Config, depth int) {
	switch n.Type {
	case html.ElementNode:
		if depth > cfg.Meta.MaxNesting {
			fmt.Printf("Depth level exceeded. Max depth level is 5, current depth level => %v\n", depth)
		}
		for i := range tagValidators {
			if tagValidators[i].ApplicableTags[n.Data] {
				isValid := tagValidators[i].ValidationFunc(n)

				if !isValid {
					fmt.Printf("[Validation errror] Tag <%s>: %s\n", n.Data, tagValidators[i].ErrorMessage)
				}
			}
		}

	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) != 0 {
			isValid := textNodeValidator.ValidationFunc(n)

			if !isValid {
				fmt.Println(textNodeValidator.ErrorMessage)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, cfg, depth+1)
	}
}

func main() {
	cfg, err := config.Load(".ruley.config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading config: %v\n", err)
		os.Exit(1)
	}

	tagValidators, textNodeValidator, err = cfg.BuildValidators()
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

	traverse(htmlAST.FirstChild.FirstChild.NextSibling.NextSibling, cfg, 0)
}
