package validation

import "testing"

type RequiredStrict struct {
	Name string `validate:"required"`
}

type RequiredButEmptyOK struct {
	MiddleName *string `validate:"required"`
}

func TestRequired(t *testing.T) {
	empty := RequiredStrict{
		Name: "",
	}
	if err := Validate(empty); err == nil {
		t.Error(err)
	}

	filled := RequiredStrict{
		Name: "Tom",
	}

	if err := Validate(filled); err != nil {
		t.Error(err)
	}

	emptyStr := ""
	empty2 := RequiredButEmptyOK{
		MiddleName: &emptyStr,
	}
	if err := Validate(empty2); err != nil {
		t.Error(err)
	}

	notPresent := RequiredButEmptyOK{
		MiddleName: nil,
	}
	if err := Validate(notPresent); err == nil {
		t.Error(err)
	}
}

type OneOfString struct {
	Weekend string `validate:"oneOf=Saturday|Sunday"`
}

type OneOfOptional struct {
	Weekend *string `validate:"oneOf=Saturday|Sunday"`
}

func TestOneOf(t *testing.T) {
	oneOf := OneOfString{
		Weekend: "Saturday",
	}
	if err := Validate(oneOf); err != nil {
		t.Error(err)
	}

	notOneOf := OneOfString{
		Weekend: "Tuesday",
	}
	if err := Validate(notOneOf); err == nil {
		t.Error(err)
	}

	weekendVal := "Saturday"
	oneOfOpt := OneOfOptional{
		Weekend: &weekendVal,
	}
	if err := Validate(oneOfOpt); err != nil {
		t.Error(err)
	}

	weekdayVal := "Tuesday"
	notOneOfOpt := OneOfOptional{
		Weekend: &weekdayVal,
	}
	if err := Validate(notOneOfOpt); err == nil {
		t.Error(err)
	}

	oneOfNotPresent := OneOfOptional{
		Weekend: nil,
	}
	if err := Validate(oneOfNotPresent); err == nil {
		t.Error(err)
	}
}

type RequiredAndOneOf struct {
	MiddleName *string `validate:"required,oneOf=Albert|Dawson"`
}

func TestMultipleRules(t *testing.T) {
	mName := "Dawson"
	multiple := RequiredAndOneOf{
		MiddleName: &mName,
	}
	if err := Validate(multiple); err != nil {
		t.Error(err)
	}
}
