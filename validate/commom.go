package validate

import (
	"fmt"
	"strings"
	"web/store"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Run() {
	db := store.ConnectDB()
	v, check := binding.Validator.Engine().(*validator.Validate)

	if check {
		v.RegisterValidation("uniqueInDB", UniqueInDB(db))
	}
}

func UniqueInDB(db *gorm.DB) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		fmt.Printf("\n====%+v===", strings.Split(fl.Param(), " "))
		fmt.Printf("\n====%+v===", fl.Field())
		fmt.Printf("\n====%+v===", fl.FieldName())
		fmt.Printf("\n====%+v===", fl.StructFieldName())
		fmt.Printf("\n====%+v===", fl.GetTag())
		return false
	}
}

func ResponseErrorValidate(ve []validator.FieldError) map[string]string {
	messageError := make(map[string]string, len(ve))

	for _, fe := range ve {
		messageError[fe.Field()] = GetMessageError(fe)
	}

	return messageError
}

func GetMessageError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"

	case "min":
		return "the field is require min 8 characters"
	}

	return "Unknown field"
}
