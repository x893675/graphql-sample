package validate

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate(dataSet interface{}) (bool, string) {
	err := v.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//Validation errors occurred
		var errString string

		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errString = "The " + name + " is required"
				break
			case "email":
				errString = "The " + name + " should be a valid email"
				break
			case "eqfield":
				errString = "The " + name + " should be equal to the " + err.Param()
				break
			default:

				errString = "The " + name + " is invalid"
				break
			}
		}
		return false, errString
	}
	return true, ""
}
