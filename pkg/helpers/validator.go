package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateStruct valida un struct basándose en tags
func ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Validar tag "validate"
		if tag := field.Tag.Get("validate"); tag != "" {
			tags := strings.Split(tag, ",")
			for _, t := range tags {
				// required
				if t == "required" && isZero(value) {
					errors = append(errors, ValidationError{
						Field:   field.Name,
						Message: "field is required",
					})
				}

				// email
				if t == "email" && value.Kind() == reflect.String {
					if !ValidateEmail(value.String()) {
						errors = append(errors, ValidationError{
							Field:   field.Name,
							Message: "invalid email format",
						})
					}
				}

				// min=N
				if strings.HasPrefix(t, "min=") {
					minStr := strings.TrimPrefix(t, "min=")
					var min int
					fmt.Sscanf(minStr, "%d", &min)
					if value.Kind() == reflect.Int && int(value.Int()) < min {
						errors = append(errors, ValidationError{
							Field:   field.Name,
							Message: fmt.Sprintf("value must be at least %d", min),
						})
					}
				}

				// max=N
				if strings.HasPrefix(t, "max=") {
					maxStr := strings.TrimPrefix(t, "max=")
					var max int
					fmt.Sscanf(maxStr, "%d", &max)
					if value.Kind() == reflect.Int && int(value.Int()) > max {
						errors = append(errors, ValidationError{
							Field:   field.Name,
							Message: fmt.Sprintf("value must be at most %d", max),
						})
					}
				}
			}
		}
	}

	return errors
}

// isZero verifica si un valor es el zero value
func isZero(v reflect.Value) bool {
	return v.IsZero()
}

// ValidateEmail valida formato de email básico
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// ValidateURL valida formato URL básico
func ValidateURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// ValidatePhone valida formato de teléfono básico (solo números y guiones)
func ValidatePhone(phone string) bool {
	for _, char := range phone {
		if !((char >= '0' && char <= '9') || char == '-' || char == '+' || char == ' ') {
			return false
		}
	}
	return len(phone) >= 7
}

// ValidateLength valida longitud de string
func ValidateLength(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}

// ValidateRequired valida campo requerido
func ValidateRequired(value interface{}) bool {
	if value == nil {
		return false
	}

	v := reflect.ValueOf(value)
	return !v.IsZero()
}

// ValidateNumeric valida valor numérico
func ValidateNumeric(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// ValidateMin valida valor mínimo
func ValidateMin(value interface{}, min float64) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()) >= min
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()) >= min
	case reflect.Float32, reflect.Float64:
		return v.Float() >= min
	}
	return false
}

// ValidateMax valida valor máximo
func ValidateMax(value interface{}, max float64) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()) <= max
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()) <= max
	case reflect.Float32, reflect.Float64:
		return v.Float() <= max
	}
	return false
}
