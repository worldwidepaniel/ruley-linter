package requiredattribute

import (
	"fmt"
	"slices"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(attributeName string, tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[RequiredAttribute] Tag should have %s attribute", attributeName)
	}

	validationLogic := func(node *html.Node) bool {
		return slices.ContainsFunc(node.Attr, func(attr html.Attribute) bool {
			return attr.Key == attributeName
		})
	}

	return validator.New(tags, errMsg, validationLogic)
}
