package validator

import (
	"errors"
	"fmt"
	"myapp/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// dateValidation 入力パラメータのフォーマットが日付型(YYYY-MM-DD限定)であるか確認
// Note: 初期値である空文字やnilの場合もfalseを返す。
func dateValidation(fl validator.FieldLevel) bool {
	if _, err := utils.ParseTime(fl.Field().String()); err != nil {
		return false
	}
	return true
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	v := validator.New()
	v.RegisterValidation("is_date", dateValidation)
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)

	if err != nil {
		var messges []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			field := fieldErr.Field()
			value := fieldErr.Value()
			tag := fieldErr.Tag()

			var messge string
			// validationのログ内容をカスタマイズ
			switch tag {
			case "required":
				messge = fieldErr.Error()
			case "is_date":
				messge = fieldErr.Error() + fmt.Sprintf("%s = %v", field, value)
			default:
				messge = fieldErr.Error()
			}
			messges = append(messges, messge)
		}

		err = errors.New(strings.Join(messges[:], " "))
	}

	return err
}
