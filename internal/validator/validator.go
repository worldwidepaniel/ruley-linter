package validator

import "golang.org/x/net/html"

type TagValidator struct {
	ErrorMessage   string
	ValidationFunc func(node *html.Node) bool
	ApplicableTags map[string]bool
}

func New(tags []string, errorMessage string, validationFunc func(*html.Node) bool) TagValidator {
	applicableTagsMap := make(map[string]bool, len(tags))
	for _, tag := range tags {
		applicableTagsMap[tag] = true
	}

	return TagValidator{
		ErrorMessage:   errorMessage,
		ValidationFunc: validationFunc,
		ApplicableTags: applicableTagsMap,
	}
}
