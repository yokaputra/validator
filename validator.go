package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	enLang          = en.New()
	uni             = ut.New(enLang, enLang)
	translatorMutex sync.Mutex
)

const (
	ErrValidationFailed = "validation failed on field"
)

type IValidator interface {
	Validate(i interface{}) error
}

type ValidationError struct {
	Field    string `json:"field"`
	Tag      string `json:"tag"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

type validatorImpl struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidator() IValidator {
	trans, _ := getTranslator("en")
	v := initValidator(trans)
	return &validatorImpl{
		validate: v,
		trans:    trans,
	}
}

func getTranslator(lang string) (ut.Translator, error) {
	trans, _ := uni.GetTranslator(lang)
	return trans, nil
}

func initValidator(trans ut.Translator) *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	translatorMutex.Lock()
	en_translations.RegisterDefaultTranslations(v, trans)
	translatorMutex.Unlock()

	return v
}

func (v *validatorImpl) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		validationErrors := getValidatorMessage(err)
		errorMessage := formatValidationErrorMessages(validationErrors)
		return fmt.Errorf(errorMessage)
	}

	return nil
}

func formatValidationErrorMessages(validationErrors []ValidationError) string {
	var errorMessages []string
	for _, ve := range validationErrors {
		message := fmt.Sprintf("%s %s: %s", ErrValidationFailed, ve.Field, ve.Tag)
		if ve.Tag == "oneof" {
			message = fmt.Sprintf("%s %s: %s, Expected: %s, Actual: %s", ErrValidationFailed, ve.Field, ve.Tag, ve.Expected, ve.Actual)
		}
		errorMessages = append(errorMessages, message)
	}
	return strings.Join(errorMessages, "\n")
}

func getValidatorMessage(err error) []ValidationError {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return nil
	}

	var errs []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		errorMsg := ValidationError{
			Field:    err.Field(),
			Tag:      err.ActualTag(),
			Expected: err.Param(),
			Actual:   fmt.Sprintf("%v", err.Value()),
		}

		errs = append(errs, errorMsg)
	}

	return errs
}
