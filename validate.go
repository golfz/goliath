package goliath

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golfz/goliath/x/data/output"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func Validate(v interface{}) output.GoliathError {
	if err := validator.New().Struct(v); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			stack := string(debug.Stack())
			errorStr := fmt.Sprint(err)
			return &output.Error{
				Status:  http.StatusBadRequest,
				Time:    time.Now(),
				Type:    "validation",
				Code:    fmt.Sprintf("validation-error"),
				Error:   fmt.Sprintf("Validation error"),
				Message: fmt.Sprintf("Cannot validate data. The input may not be struct (The outermost input must not be an array)."),
				ErrorDev: output.ErrorDev{
					Error:      errorStr,
					Stacktrace: stack,
				},
			}
		}

		for _, err := range err.(validator.ValidationErrors) {
			stack := string(debug.Stack())
			errorStr := fmt.Sprint(err)
			jsons := getJsonTag(v, err.StructNamespace())

			if err.Kind() == reflect.String {
				v := reflect.ValueOf(err.Value()).String()
				if v == "" && err.ActualTag() != "required" {
					continue
				}
			}

			return &output.Error{
				Status:  http.StatusBadRequest,
				Time:    time.Now(),
				Type:    "validation",
				Code:    fmt.Sprintf("validation-%s-%s", jsons, err.Tag()),
				Error:   fmt.Sprintf("Invalid '%s': %#v", jsons, err.Value()),
				Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' rule (%s=%s)", jsons, err.Tag(), err.Tag(), err.Param()),
				ErrorDev: output.ErrorDev{
					Error:      errorStr,
					Stacktrace: stack,
				},
			}
		}
	}
	return nil
}

func getJsonTag(datasets interface{}, structNameSpace string) string {
	output := ""

	firstDotIndex := strings.Index(structNameSpace, ".")
	nameSpace := structNameSpace[(firstDotIndex + 1):]

	arrNameSpace := strings.Split(nameSpace, ".")
	level := len(arrNameSpace)

	currentNameSpace := arrNameSpace[0]

	isArray := false
	arrIndex := -1

	startBracketIndex := strings.Index(currentNameSpace, "[")
	endBracketIndex := strings.Index(currentNameSpace, "]")
	isArray = (startBracketIndex != -1)
	if isArray {
		strIndex := currentNameSpace[(startBracketIndex + 1):endBracketIndex]
		arrIndex, _ = strconv.Atoi(strIndex)
		currentNameSpace = currentNameSpace[:startBracketIndex]
	}

	v := reflect.ValueOf(datasets)
	field, _ := v.Type().FieldByName(currentNameSpace)
	jsonText := field.Tag.Get("json")

	if jsonText == "" {
		jsonText = currentNameSpace
	}

	if isArray {
		output = fmt.Sprintf("%s[%d]", jsonText, arrIndex)
	} else {
		output = jsonText
	}

	if level == 1 {
		return output
	}

	if isArray {
		r := v.FieldByName(currentNameSpace).Interface()
		items := reflect.ValueOf(r)
		item := items.Index(arrIndex)
		jsonText = getJsonTag(item.Interface(), nameSpace)
		output = fmt.Sprintf("%s.%s", output, jsonText)
	} else {
		jsonText = getJsonTag(v.FieldByName(currentNameSpace).Interface(), nameSpace)
		output = fmt.Sprintf("%s.%s", output, jsonText)
	}

	return output
}
