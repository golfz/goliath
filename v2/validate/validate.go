package validate

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golfz/goliath/v2"
)

type errorArgs struct {
	Key           string      `json:"key"`
	ActualValue   interface{} `json:"actual_value"`
	Rule          string      `json:"rule"`
	ExpectedValue string      `json:"expected_value"`
	Message       string      `json:"message"`
}

// Struct validate struct data
func Struct(v interface{}) goliath.Error {
	if err := validator.New().Struct(v); err != nil {

		var (
			errStatus  = http.StatusBadRequest
			errCode    string
			errArgs    map[string]interface{}
			errMessage string
			logID      = ""
		)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			errCode = "goliath.validate.Struct.InvalidValidationError"
			errArgs = nil
			errMessage = "invalid validation error"
			return goliath.NewError(errStatus, errCode, errArgs, err, logID, errMessage)
		}

		var errDetails []errorArgs = nil

		for _, err := range err.(validator.ValidationErrors) {
			jsonKey := getInvalidJsonKey(v, err.StructNamespace())
			errDetails = append(errDetails, errorArgs{
				Key:           jsonKey,
				ActualValue:   err.Value(),
				Rule:          err.Tag(),
				ExpectedValue: err.Param(),
				Message:       fmt.Sprintf("Field validation for '%s' failed on the '%s' rule (%s=%s)", jsonKey, err.Tag(), err.Tag(), err.Param()),
			})

		}

		if errDetails != nil {
			errCode = "goliath.validate.Struct.ValidationErrors"
			errArgs = map[string]interface{}{"validation_errors": errDetails}
			errMessage = "validation errors"
			return goliath.NewError(errStatus, errCode, errArgs, err, logID, errMessage)
		}

	}
	return nil
}

func getInvalidJsonKey(datasets interface{}, structNameSpace string) string {
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
		jsonText = getInvalidJsonKey(item.Interface(), nameSpace)
		output = fmt.Sprintf("%s.%s", output, jsonText)
	} else {
		jsonText = getInvalidJsonKey(v.FieldByName(currentNameSpace).Interface(), nameSpace)
		output = fmt.Sprintf("%s.%s", output, jsonText)
	}

	return output
}
