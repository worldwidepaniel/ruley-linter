package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"ruley-linter/internal/validator"
	"ruley-linter/internal/validators/allowedattributevalues"
	"ruley-linter/internal/validators/attributevalueregex"
	"ruley-linter/internal/validators/disallowedattribute"
	"ruley-linter/internal/validators/disallowedattributevalues"
	"ruley-linter/internal/validators/disallowedtags"
	"ruley-linter/internal/validators/maxlength"
	"ruley-linter/internal/validators/minlength"
	"ruley-linter/internal/validators/requiredattribute"
	"ruley-linter/internal/validators/textregex"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading config file %q: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error while parsing config file %q: %w", path, err)
	}

	return &cfg, nil
}

func (c *Config) BuildValidators() ([]validator.TagValidator, validator.TextNodeValidator, error) {
	result := make([]validator.TagValidator, 0, len(c.TagValidators))

	for _, vc := range c.TagValidators {
		v, err := buildValidator(vc)
		if err != nil {
			return nil, validator.TextNodeValidator{}, fmt.Errorf("error while building validator %q: %w", vc.Name, err)
		}
		result = append(result, v)
	}

	textNodeValidator, err := buildTextNodeValidator(c.TextNodeValidator)
	if err != nil {
		return nil, validator.TextNodeValidator{}, fmt.Errorf("error while building text node validator: %w", err)
	}

	return result, textNodeValidator, nil
}

func buildTextNodeValidator(vc TextNodeValidatorConfig) (validator.TextNodeValidator, error) {
	regexp, err := regexp.Compile(vc.Value)
	if err != nil {
		return validator.TextNodeValidator{}, fmt.Errorf("error while compiling regexString to regexp: %w", err)
	}

	return textregex.NewValidator(regexp, vc.ErrorMessage), nil
}

func buildValidator(vc TagValidatorConfig) (validator.TagValidator, error) {
	switch vc.Type {
	case MaxLength:
		var limit int
		if err := json.Unmarshal(vc.Value.Value, &limit); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling maxLength value (expected integer): %w", err)
		}
		return maxlength.NewValidator(vc.Value.Target, limit, vc.AllowedTags, vc.ErrorMessage), nil

	case MinLength:
		var limit int
		if err := json.Unmarshal(vc.Value.Value, &limit); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling maxLength value (expected integer): %w", err)
		}
		return minlength.NewValidator(vc.Value.Target, limit, vc.AllowedTags, vc.ErrorMessage), nil

	case RequiredAttribute:
		var attrName string
		if err := json.Unmarshal(vc.Value.Value, &attrName); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling requiredAttribute value (expected string): %w", err)
		}
		return requiredattribute.NewValidator(attrName, vc.AllowedTags, vc.ErrorMessage), nil

	case DisallowedAttribute:
		var attrName string
		if err := json.Unmarshal(vc.Value.Value, &attrName); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling disallowedAttribute value (expected string): %w", err)
		}
		return disallowedattribute.NewValidator(attrName, vc.AllowedTags, vc.ErrorMessage), nil

	case AllowedAttributeValues:
		var values []string
		if err := json.Unmarshal(vc.Value.Value, &values); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling allowedAttributeValues value (expected array of strings): %w", err)
		}
		return allowedattributevalues.NewValidator(vc.Value.Target, values, vc.AllowedTags, vc.ErrorMessage), nil

	case DisallowedAttributeValues:
		var values []string
		if err := json.Unmarshal(vc.Value.Value, &values); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling disallowedAttributeValues value (expected array of strings): %w", err)
		}
		return disallowedattributevalues.NewValidator(vc.Value.Target, values, vc.AllowedTags, vc.ErrorMessage), nil

	case AttributeValueRegex:
		var regexString string
		if err := json.Unmarshal(vc.Value.Value, &regexString); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling textRegex value (expected regex in form of a string): %w", err)
		}
		regexp, err := regexp.Compile(regexString)
		if err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while compiling regexString to regexp: %w", err)
		}

		return attributevalueregex.NewValidator(vc.Value.Target, regexp, vc.AllowedTags, vc.ErrorMessage), nil

	case DisallowedTags:
		var values []string
		if err := json.Unmarshal(vc.Value.Value, &values); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling disallowedTags value (expected array of strings): %w", err)
		}
		return disallowedtags.NewValidator(vc.AllowedTags, vc.ErrorMessage), nil

	default:
		return validator.TagValidator{}, fmt.Errorf(
			"unknown type %q — valid types: %v",
			vc.Type,
			AllTagValidatorTypes,
		)
	}
}
