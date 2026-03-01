package attributevalueregex

import (
	"fmt"
	"regexp"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(attributeName string, regex *regexp.Regexp, tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[AttributeValueRegex] '%s' does not match given regex '%v'", attributeName, regex.String())
	}

	validationLogic := func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == attributeName {
				return regex.MatchString(attr.Val)
			}
		}

		return true
	}

	return validator.New(tags, errMsg, validationLogic)
}
