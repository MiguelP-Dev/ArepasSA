package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidators registra validadores personalizados
func RegisterCustomValidators(v *validator.Validate) {
	_ = v.RegisterValidation("phone", ValidatePhone)
	_ = v.RegisterValidation("notzero", ValidateNotZero)
}

// ValidateNotZero valida que un número no sea cero
var ValidateNotZero validator.Func = func(fl validator.FieldLevel) bool {
	switch fl.Field().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fl.Field().Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fl.Field().Uint() != 0
	case reflect.Float32, reflect.Float64:
		return fl.Field().Float() != 0
	default:
		return false
	}
}

// ValidateStock valida que haya suficiente stock
func ValidateStock(fl validator.FieldLevel) bool {
	// Esta validación requeriría acceso a la base de datos
	// Se implementaría de manera diferente en la práctica
	return true
}
