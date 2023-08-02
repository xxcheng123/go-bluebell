package controllers

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitTrans() error {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		return nil
	} else {
		return errors.New("get gin Validator failed")
	}
}
