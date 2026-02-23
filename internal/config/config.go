package config

import (
	"encoding/json"
	"fmt"
	"os"

	"ruley-linter/internal/validator"
	"ruley-linter/internal/validators/allowedvalues"
	"ruley-linter/internal/validators/maxlength"
	"ruley-linter/internal/validators/requiredattribute"
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

func (c *Config) BuildValidators() ([]validator.TagValidator, error) {
	result := make([]validator.TagValidator, 0, len(c.TagValidators))

	for _, vc := range c.TagValidators {
		v, err := buildValidator(vc)
		if err != nil {
			return nil, fmt.Errorf("error while building validator %q: %w", vc.Name, err)
		}
		result = append(result, v)
	}

	return result, nil
}

func buildValidator(vc TagValidatorConfig) (validator.TagValidator, error) {
	switch vc.Type {

	case "requiredAttribute":
		var attrName string
		if err := json.Unmarshal(vc.Value.Value, &attrName); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling requiredAttribute value (expected string): %w", err)
		}
		return requiredattribute.NewValidator(attrName, vc.AllowedTags, vc.ErrorMessage), nil

	case "maxLength":
		var limit int
		if err := json.Unmarshal(vc.Value.Value, &limit); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling maxLength value (expected integer): %w", err)
		}
		return maxlength.NewValidator(vc.Value.Target, limit, vc.AllowedTags, vc.ErrorMessage), nil

	case "allowedValues":
		var values []string
		if err := json.Unmarshal(vc.Value.Value, &values); err != nil {
			return validator.TagValidator{}, fmt.Errorf("error while unmarshaling allowedValues value (expected array of strings): %w", err)
		}
		return allowedvalues.NewValidator(vc.Value.Target, values, vc.AllowedTags, vc.ErrorMessage), nil

	default:
		return validator.TagValidator{}, fmt.Errorf(
			"unknown type %q — valid types: requiredAttribute, maxLength, allowedValues",
			vc.Type,
		)
	}
}
