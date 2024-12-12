package validatorx

import (
	"reflect"
	"strings"

	goValidator "github.com/go-playground/validator/v10"
)

// ------------------- Validator -------------------

type Validator interface {
	Validate(out any) error
}

// ------------------- New -------------------

type validator struct {
	validator *goValidator.Validate
}

func NewValidatorx() ValidatorSetter {
	validate := goValidator.New()

	return &validator{
		validator: validate,
	}
}

// ------------------- Setter -------------------

type ValidatorSetter interface {
	Init() Validator
	AddPasswordAtLeastOneCharNumValidation(tag string) ValidatorSetter
	AddCheckMustFieldIfIdFieldExistValidation(tagName string) ValidatorSetter
	AddPhoneNumValidation(tagName string) ValidatorSetter
	AddUrlValidation(tagName string) ValidatorSetter
	SetExtractTagName() ValidatorSetter
}

func (v *validator) AddPasswordAtLeastOneCharNumValidation(tagName string) ValidatorSetter {
	if err := v.validator.RegisterValidation(tagName, validatePasswordAtLeastOneCharNum); err != nil {
		panic(err)
	}
	return v
}

func (v *validator) AddPhoneNumValidation(tagName string) ValidatorSetter {
	if err := v.validator.RegisterValidation(tagName, validatePhoneNum); err != nil {
		panic(err)
	}
	return v
}

func (v *validator) AddUrlValidation(tagName string) ValidatorSetter {
	if err := v.validator.RegisterValidation(tagName, validateUrl); err != nil {
		panic(err)
	}
	return v
}

func (v *validator) AddCheckMustFieldIfIdFieldExistValidation(tagName string) ValidatorSetter {
	if err := v.validator.RegisterValidation(tagName, validateCheckMustFieldIfIdFieldExist, true); err != nil {
		panic(err)
	}
	return v
}

func (v *validator) SetExtractTagName() ValidatorSetter {
	v.validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonName := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		queryName := strings.SplitN(field.Tag.Get("query"), ",", 2)[0]
		paramName := strings.SplitN(field.Tag.Get("params"), ",", 2)[0]

		if jsonName != "" && jsonName != "-" {
			return jsonName
		}

		if queryName != "" && queryName != "-" {
			return queryName
		}

		if paramName != "" && paramName != "-" {
			return paramName
		}

		return ""
	})
	return v
}

func validateCheckMustFieldIfIdFieldExist(fl goValidator.FieldLevel) bool {
	fls := fl.Parent()
	if !fls.FieldByName("Id").IsNil() {
		return true
	} else {
		if fl.Field().IsZero() {
			return false
		}
	}
	return true
}

func validatePhoneNum(fl goValidator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return ValidateRegex(PhoneRegex, value)
	}
	return true
}

func validateUrl(fl goValidator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return ValidateRegex(UrlRegex, value)
	}
	return true
}

func validatePasswordAtLeastOneCharNum(fl goValidator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return ValidateRegex(ForbiddenSpecialCharRegex, value) && ValidateRegex(AtLeastOneCharOneNumRegex, value)
	}
	return true
}

func (v *validator) Init() Validator {
	return v
}

// ------------------- Validator Method -------------------
type ErrorResponse struct {
	FailedField        string
	FailedFieldTagName string
	Tag                string
	Value              string
}

func (v *validator) Validate(out any) error {
	err := v.validator.Struct(out)
	if err != nil {
		return err
	}
	return nil
}
