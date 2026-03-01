package disallowedattributevalues

import (
	"fmt"
	"slices"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(attributeName string, disallowedValues []string, tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[DisallowedValues] '%s' must not be one of: %v", attributeName, disallowedValues)
	}

	validationLogic := func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == attributeName {
				return !slices.Contains(disallowedValues, attr.Val)
			}
		}

		return true
	}

	return validator.New(tags, errMsg, validationLogic)
}
