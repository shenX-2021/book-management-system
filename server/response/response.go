package response

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 获取校验的错误信息
func getValidateErrMsg(err error, dto interface{}) string {
	dtoType := reflect.TypeOf(dto)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			if f, ok := dtoType.Elem().FieldByName(err.Field()); ok {
				fmt.Println(dtoType.Elem(), f.Tag, dtoType.Elem())
				return f.Tag.Get("validateErrMsg")
			}
		}
	}

	return err.Error()
}

// 参数校验识别返回
func ValidateError(c *gin.Context, err error, dto interface{}) {
	c.String(http.StatusBadRequest, getValidateErrMsg(err, dto))
}

func BadRequest(c *gin.Context, str string) {
	c.String(http.StatusBadRequest, str)
}
