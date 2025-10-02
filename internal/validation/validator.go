package validation

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// ValidateAndDecode unmarshals JSON and validates struct tags
func ValidateAndDecode(data []byte, v any) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return Validate(v)
}

// Validate checks struct validate tags
func Validate(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := range val.NumField() {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "" {
			continue
		}

		rules := strings.SplitSeq(tag, ",")
		for rule := range rules {
			if err := validateRule(field, fieldType.Name, rule); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateRule(field reflect.Value, fieldName, rule string) error {
	log.Printf("validating field: %s, rule: %s", fieldName, rule)
	switch rule {
	case "required":
		if field.IsZero() {
			return fmt.Errorf("%s is required", fieldName)
		}
	case "oneOf":
		if field.Kind() != reflect.String {
			return fmt.Errorf("field %s must be a string to use 'oneOf' validator ", fieldName)
		}
		options := strings.SplitSeq(rule, ",")
		for option := range options {
			if option == field.String() {
				return nil
			}
		}
		return fmt.Errorf("%s value is not oneOf %v", fieldName, options)
	default:
		log.Printf("uncaught rule, field: %s, rule: %s", fieldName, rule)
	}
	return nil
}
