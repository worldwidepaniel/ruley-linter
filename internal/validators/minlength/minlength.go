package minlength

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(target string, value int, tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[MinLength] Attribute '%s' value should exceed %d characters", target, value)
	}

	validationLogic := func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == target {
				charCount := utf8.RuneCountInString(attr.Val)
				return charCount >= value
			}
		}

		return true
	}

	return validator.New(tags, errMsg, validationLogic)
}
