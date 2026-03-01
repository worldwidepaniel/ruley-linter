package textregex

import (
	"fmt"
	"regexp"

	"golang.org/x/net/html"

	"ruley-linter/internal/validator"
)

func NewValidator(regex *regexp.Regexp, customErrorMsg string) validator.TextNodeValidator {
	errMsg := customErrorMsg
	if errMsg == "" {
		errMsg = fmt.Sprintf("[TextRegex] TextNode does not satisfy regex expression %v", regex.String())
	}

	validationLogic := func(node *html.Node) bool {
		return regex.MatchString(node.Data)
	}

	return validator.NewTextNode(errMsg, validationLogic)
}
