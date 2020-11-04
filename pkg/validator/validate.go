/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:20
 */
package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type Validator struct {
	Trans ut.Translator
}

type Register struct {
	Func func(validator.FieldLevel) bool
	Msg  string
}

func (va *Validator) GetParamError(err error) []string {
	errors, ok := err.(validator.ValidationErrors)
	res := make([]string, 0)
	if ok {
		for _, err := range errors.Translate(va.Trans) {
			//fmt.Println(field[strings.Index(field, ".")+1:])
			res = append(res, err)
		}
	}
	return res
}

func (va *Validator) InitTrans(locale string, register map[string]*Register) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			var name string
			switch locale {
			case "en":
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			case "zh":
				name = fld.Tag.Get("label")
			default:
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		va.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, va.Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, va.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, va.Trans)
		}

		// 自定义翻译方法
		for key, value := range register {
			if err := v.RegisterValidation(key, value.Func); err != nil {
				return err
			}

			if err := v.RegisterTranslation(
				key,
				va.Trans,
				registerTranslator(key, value.Msg),
				tFunc,
			); err != nil {
				return err
			}
		}
	}
	return
}

func New() *Validator {
	return new(Validator)
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func tFunc(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}