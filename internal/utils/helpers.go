package utils

import (
	"ArepasSA/internal/models"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Custom validators
var ValidatePhone validator.Func = func(fl validator.FieldLevel) bool {
	phone, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	// Validación simple: solo números y al menos 8 dígitos
	if len(phone) < 8 {
		return false
	}
	for _, c := range phone {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// TranslateValidationErrors traduce errores de validación a mensajes legibles
func TranslateValidationErrors(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s es requerido", field))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("%s debe ser un email válido", field))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s debe tener al menos %s caracteres", field, e.Param()))
			case "gt":
				errorMessages = append(errorMessages, fmt.Sprintf("%s debe ser mayor que %s", field, e.Param()))
			case "gte":
				errorMessages = append(errorMessages, fmt.Sprintf("%s debe ser mayor o igual a %s", field, e.Param()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("validación falló para %s", field))
			}
		}
		return fmt.Errorf("%s", strings.Join(errorMessages, "; "))
	}
	return err
}

// FormatCurrency formatea números como moneda
func FormatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// Round redondea números a 2 decimales
func Round(value float64) float64 {
	return math.Round(value*100) / 100
}

// ParseDate parsea fechas en formato YYYY-MM-DD
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// Contains verifica si un slice contiene un elemento
func Contains(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

// GenerateBarcode genera un código de barras ficticio (en producción usaría una librería real)
func GenerateBarcode() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()%10000000000000)
}

// CalculateComboPrice calcula el precio total de un combo basado en sus productos
func CalculateComboPrice(items []models.ComboItem, products []models.Product) float64 {
	total := 0.0
	for _, item := range items {
		for _, product := range products {
			if product.ID == item.ProductID {
				total += product.SellPrice * item.Quantity
				break
			}
		}
	}
	return Round(total)
}
