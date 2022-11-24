package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Verify(data interface{}) error {
	trans := GetTrans()

	Trans(trans)

	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, e := range err.(validator.ValidationErrors) {
			return e
		}
	}

	return nil
}

func GetTrans() ut.Translator {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

func Trans(trans ut.Translator) {
	validate.RegisterTranslation("required", trans, func(ut1 ut.Translator) error {
		return ut1.Add("required", "{0} must have a value!", true)
	}, func(ut2 ut.Translator, fe validator.FieldError) string {
		t, _ := ut2.T("required", fe.Field())

		return t
	})
}
