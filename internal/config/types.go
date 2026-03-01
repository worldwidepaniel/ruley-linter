package config

import "encoding/json"

type TagValidatorType string

const (
	MaxLength TagValidatorType = "maxLength"
	MinLength TagValidatorType = "minLength"

	RequiredAttribute   TagValidatorType = "requiredAttribute"
	DisallowedAttribute TagValidatorType = "disallowedAttribute"

	AllowedAttributeValues    TagValidatorType = "allowedAttributeValues"
	DisallowedAttributeValues TagValidatorType = "disallowedAttributeValues"

	TextRegex           TagValidatorType = "textRegex"
	AttributeValueRegex TagValidatorType = "attributeValueRegex"

	DisallowedTags TagValidatorType = "disallowedTags"
)

var AllTagValidatorTypes = []TagValidatorType{
	MaxLength,
	MinLength,
	RequiredAttribute,
	DisallowedAttribute,
	AllowedAttributeValues,
	DisallowedAttributeValues,
	TextRegex,
	AttributeValueRegex,
	DisallowedTags,
}

type Meta struct {
	MaxNesting  int `json:"maxNesting"`
	MaxCSSFiles int `json:"maxCssFiles"`
}

type ValidatorValue struct {
	Target string          `json:"target"`
	Value  json.RawMessage `json:"value"`
}

type TagValidatorConfig struct {
	Name         string           `json:"name"`
	Type         TagValidatorType `json:"type"`
	Value        ValidatorValue   `json:"value"`
	AllowedTags  []string         `json:"allowedTags"`
	ErrorMessage string           `json:"errorMessage"`
}

type TextNodeValidatorConfig struct {
	Name         string `json:"name"`
	Value        string `json:"value"`
	ErrorMessage string `json:"errorMessage"`
}

type Config struct {
	Meta              Meta                    `json:"meta"`
	TagValidators     []TagValidatorConfig    `json:"tagValidators"`
	TextNodeValidator TextNodeValidatorConfig `json:"textNodeValidator"`
}
