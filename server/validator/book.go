package validator

import (
	"strconv"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var bookISBN validator.Func = func(fl validator.FieldLevel) bool {
	isbn, ok := fl.Field().Interface().(int64)
	if ok {
		str := strconv.FormatInt(isbn, 10)
		l := len(str)
		return l == 10 || l == 13
	}
	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookISBN", bookISBN)
	}
}
