package config

import "encoding/json"

type Meta struct {
	MaxNesting                int  `json:"maxNesting"`
	MaxCSSFiles               int  `json:"maxCssFiles"`
	AllowCdnFonts             bool `json:"allowCdnFonts"`
	AllowInlineStyles         bool `json:"allowInlineStyles"`
	AllowStyleTags            bool `json:"allowStyleTags"`
	AllowExternalLinks        bool `json:"allowExternalLinks"`
	RequireLazyClassForImages bool `json:"requireLazyClassForImages"`
}

type ValidatorValue struct {
	Target string          `json:"target"`
	Value  json.RawMessage `json:"value"`
}

type TagValidatorConfig struct {
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Value        ValidatorValue `json:"value"`
	AllowedTags  []string       `json:"allowedTags"`
	ErrorMessage string         `json:"errorMessage"`
}

type Config struct {
	Meta          Meta                 `json:"meta"`
	TagValidators []TagValidatorConfig `json:"tagValidators"`
}
