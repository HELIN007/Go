package validatoor

import (
	"GinProject/middleware"
	"GinProject/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uniTrans := ut.New(zh_Hans_CN.New())
	trans, _ := uniTrans.GetTranslator("zh_Hans_CN")

	//注册默认的翻译
	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("====", err)
	}

	//映射label标签
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			middleware.Infof("验证信息错误：%v", v.Translate(trans))
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
