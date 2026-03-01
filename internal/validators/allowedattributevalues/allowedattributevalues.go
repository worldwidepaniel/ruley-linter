package allowedattributevalues

import (
	"fmt"
	"slices"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(attributeName string, allowedValues []string, tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[AllowedValues] '%s' must be one of allowed values: %v", attributeName, allowedValues)
	}

	validationLogic := func(node *html.Node) bool {
		for _, attr := range node.Attr {
			if attr.Key == attributeName {
				return slices.Contains(allowedValues, attr.Val)
			}
		}

		return true
	}

	return validator.New(tags, errMsg, validationLogic)
}
