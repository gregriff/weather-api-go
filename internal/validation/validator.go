package validation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// DecodeAndValidate unmarshals JSON from body into obj, then validates obj using its struct tags
func DecodeAndValidate(body io.ReadCloser, obj any) error {
	if err := json.NewDecoder(body).Decode(obj); err != nil {
		return fmt.Errorf("json decode error: %w", err)
	}
	return Validate(obj)
}

// ValidateAndEncode validates struct tags and then marshals it into JSON to create a http response body
func ValidateAndEncode(body http.ResponseWriter, obj any) error {
	if vErr := Validate(obj); vErr != nil {
		return vErr
	}

	if err := json.NewEncoder(body).Encode(obj); err != nil {
		return fmt.Errorf("json encode error: %w", err)
	}
	return nil
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
				return fmt.Errorf("validation error: %w", err)
			}
		}
	}
	return nil
}

func validateRule(field reflect.Value, fieldName, rule string) error {
	// log.Printf("validating field: %s, rule: %s", fieldName, rule)
	switch rule {
	case "required":
		if field.IsZero() {
			return fmt.Errorf("%s is required", fieldName)
		}
	default:
		if strings.HasPrefix(rule, "oneOf") {
			if !strings.Contains(rule, "=") {
				return fmt.Errorf("invalid oneOf rule: %s. ensure rule is of the format `oneOf=%%s|%%s...`", rule)
			}
			kind := field.Kind()
			if kind != reflect.Pointer && kind != reflect.String {
				return fmt.Errorf("invalid oneOf rule: only string fields are supported")
			}
			optionPart := strings.Split(rule, "=")[1]

			var value string
			if kind == reflect.Pointer {
				value = field.Elem().String()
			} else {
				value = field.String()
			}

			options := strings.SplitSeq(optionPart, "|")
			for option := range options {
				if option == value {
					return nil
				}
			}
			return fmt.Errorf("%s value is not oneOf %v", fieldName, optionPart)
		}
		log.Printf("uncaught rule, field: %s, rule: %s", fieldName, rule)
	}
	return nil
}
