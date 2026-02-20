package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

type tagValidator struct {
	ErrorMessage   string
	ValidationFunc func(validator *tagValidator, node *html.Node) bool
	ApplicableTags map[string]bool
}

var dataOriginalValidator = tagValidator{
	ErrorMessage: "Tag should have data-orignal attribute",
	ValidationFunc: func(v *tagValidator, node *html.Node) bool {
		hasDataOrignalAttr := slices.ContainsFunc(node.Attr, func(attr html.Attribute) bool {
			return attr.Key == "data-orignal"
		})
		return hasDataOrignalAttr
	},
	ApplicableTags: map[string]bool{"img": true, "span": true},
}

var validators = []tagValidator{dataOriginalValidator}

func traverse(node *html.Node) {
	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			continue
		}
		fmt.Printf("%v => \n", n.Data)
		for _, validator := range validators {
			if !validator.ApplicableTags[n.Data] {
				break
			}
			hasTagPassedValidation := validator.ValidationFunc(&validator, n)
			if !hasTagPassedValidation {
				fmt.Printf("\t %v \n", validator.ErrorMessage)
			}
		}
	}
}

func main() {
	htmlFileContent, err := os.ReadFile("./test-files/invalid.html")
	if err != nil {
		panic(err)
	}

	htmlAST, err := html.Parse(strings.NewReader(string(htmlFileContent)))
	if err != nil {
		panic(err)
	}

	traverse(htmlAST.FirstChild.FirstChild.NextSibling)
}
