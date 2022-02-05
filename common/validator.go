package common

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

func InitValidator() {
	zh := zh.New()
	uni := ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) string {
	result := ""

	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result += err.Translate(trans) + "|"
	}
	return strings.TrimSuffix(result, "|")
}
