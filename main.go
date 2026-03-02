package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"ruley-linter/internal/config"
	"ruley-linter/internal/flags"
	"ruley-linter/internal/validator"
)

var (
	tagValidators     []validator.TagValidator
	textNodeValidator validator.TextNodeValidator
	validationErrors  []string
)

func traverse(n *html.Node, cfg *config.Config, depth int) {
	switch n.Type {
	case html.ElementNode:
		if depth > cfg.Meta.MaxNesting {
			validationErrors = append(validationErrors, fmt.Sprintf("Depth level exceeded. Max depth level is 5, current depth level => %v\n", depth))
		}
		for i := range tagValidators {
			if tagValidators[i].ApplicableTags[n.Data] {
				isValid := tagValidators[i].ValidationFunc(n)

				if !isValid {
					validationErrors = append(validationErrors, fmt.Sprintf("[Validation errror] Tag <%s>: %s\n", n.Data, tagValidators[i].ErrorMessage))
				}
			}
		}

	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) != 0 {
			isValid := textNodeValidator.ValidationFunc(n)

			if !isValid {
				validationErrors = append(validationErrors, textNodeValidator.ErrorMessage)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, cfg, depth+1)
	}
}

func main() {
	configPath, filePath := flags.Init()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading config: %v\n", err)
		os.Exit(1)
	}

	tagValidators, textNodeValidator, err = cfg.BuildValidators()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while building validators: %v\n", err)
		os.Exit(1)
	}

	htmlFileContent, err := os.ReadFile(*filePath)
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
	if len(validationErrors) != 0 {
		fmt.Fprintf(os.Stderr, "Encountered following errors while linting: \n")
		for _, error := range validationErrors {
			fmt.Fprintf(os.Stderr, "\t - %v\n", error)
		}
		os.Exit(1)
	}
}
