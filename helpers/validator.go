package helpers

import (
	"fmt"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	esTranslations "github.com/go-playground/validator/v10/translations/es"
	"reflect"
	"strconv"
	"strings"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

// Inicializar el validador y la traducción
func init() {
	validate = validator.New()

	// Configurar el traductor al español
	spanish := es.New()
	uni = ut.New(spanish, spanish)

	var found bool
	trans, found = uni.GetTranslator("es")
	if !found {
		fmt.Println("Traductor al español no encontrado")
		return
	}

	// Registrar las traducciones al español
	esTranslations.RegisterDefaultTranslations(validate, trans)
}

// Prevalidar y convertir los datos antes de validarlos
func PrevalidateStruct(model interface{}) error {
	// Usamos reflexión para inspeccionar los campos de la estructura
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Obtener el valor real si el modelo es un puntero
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Si el campo es un string, intentamos convertirlo a su tipo correspondiente
		if field.Kind() == reflect.String {
			strVal := field.String()

			// Verificamos si el tipo esperado es numérico y convertimos el string
			switch fieldType.Type.Kind() {
			case reflect.Float64:
				if strVal != "" {
					convertedValue, err := strconv.ParseFloat(strVal, 64)
					if err == nil {
						// Asignar el valor convertido al campo
						field.SetFloat(convertedValue)
					} else {
						return fmt.Errorf("el campo %s debe ser un número válido", fieldType.Name)
					}
				}
			case reflect.Float32:
				if strVal != "" {
					convertedValue, err := strconv.ParseFloat(strVal, 32)
					if err == nil {
						// Asignar el valor convertido al campo
						field.SetFloat(convertedValue)
					} else {
						return fmt.Errorf("el campo %s debe ser un número válido (float32)", fieldType.Name)
					}
				}
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if strVal != "" {
					convertedValue, err := strconv.ParseInt(strVal, 10, 64)
					if err == nil {
						// Asignar el valor convertido al campo
						field.SetInt(convertedValue)
					} else {
						return fmt.Errorf("el campo %s debe ser un número entero válido", fieldType.Name)
					}
				}
			}
		}
	}

	return nil
}

// ValidateStruct valida una estructura usando un mapa de alias para los campos
func ValidateStruct(model interface{}, aliases map[string]string) error {
	// Prevalidación de los datos antes de validarlos
	if err := PrevalidateStruct(model); err != nil {
		return err // Si falla la prevalidación, retornamos el error
	}

	// Registrar el nombre del campo con el alias proporcionado
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		fieldName := field.Name
		if alias, ok := aliases[fieldName]; ok {
			return alias // Utilizar el alias si está definido
		}
		return fieldName // Usar el nombre del campo si no hay alias
	})

	// Validar la estructura
	if err := validate.Struct(model); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		for _, e := range validationErrors {
			// Utilizar el traductor para obtener el mensaje de error
			msg := e.Translate(trans)
			errorMessages = append(errorMessages, msg)
		}

		// Devolver los mensajes de error concatenados
		return fmt.Errorf("Errores de validación: [%s]", strings.Join(errorMessages, ", "))
	}
	return nil
}
