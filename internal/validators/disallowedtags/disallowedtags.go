package disallowedtags

import (
	"fmt"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(tags []string, customErrorMsg string) validator.TagValidator {
	errMsg := customErrorMsg

	validationLogic := func(node *html.Node) bool {
		if errMsg == "" {
			errMsg = fmt.Sprintf("[DisallowedTags] Tag %s is not allowed", node.Data)
		}
		return false
	}

	return validator.New(tags, errMsg, validationLogic)
}
