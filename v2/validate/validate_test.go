package validate

import (
	"encoding/json"
	"github.com/golfz/goliath/v2"
	"testing"
	"time"
)

// user is input
type user struct {
	Name    *string `validate:"required,min=2"`
	Age     uint8   `validate:"gte=18,lte=130"`
	Email   string  `validate:"required,email"`
	Address address
	Phone   []phone
}

type address struct {
	Zip uint16 `validate:"required,gte=10000,lte=99999"`
}

type phone struct {
	Number string `validate:"required,min=6"`
}

// goliathErrorStruct is a struct of goliath.Error
type goliathErrorStruct struct {
	Status    int            `json:"status"`
	Message   string         `json:"message"`
	Time      time.Time      `json:"time"`
	LogID     string         `json:"log_id"`
	ErrorCode string         `json:"error_code"`
	ErrorArgs errorArguments `json:"error_args"`
	ErrorDev  errorDev       `json:"error_dev"`
}

type errorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

type errorArguments struct {
	ValidationErrors []validationErrorCase `json:"validation_errors"`
}

type validationErrorCase struct {
	Key           string      `json:"key"`
	ActualValue   interface{} `json:"actual_value"`
	Rule          string      `json:"rule"`
	ExpectedValue interface{} `json:"expected_value"`
	Message       string      `json:"message"`
}

// testCase for table test
type testCase struct {
	name          string
	input         interface{}
	expectedError goliathErrorStruct
	errArgs       int
}

func TestStruct_NoError(t *testing.T) {
	name := "Tom"
	u := user{
		Name:    &name,
		Age:     35,
		Email:   "tom@email.com",
		Address: address{Zip: 20000},
	}

	got := Struct(u)
	if got != nil {
		t.Errorf("expect = %v, got = %v", nil, got)
	}
}

func TestStruct_InvalidValidationError(t *testing.T) {
	gotErr := Struct(nil)
	gError := getGoliathErrorStruct(gotErr)

	expectedErrCode := "goliath.validate.Struct.InvalidValidationError"
	if gError.ErrorCode != expectedErrCode {
		t.Errorf("Expected %v, got %v", expectedErrCode, gError)
	}
}

func TestStruct_ValidationErrors(t *testing.T) {
	u := user{
		Name:    nil,
		Age:     35,
		Email:   "tom@email.com",
		Address: address{Zip: 20000},
	}

	err := Struct(u)
	gErr := getGoliathErrorStruct(err)

	expectedErrCode := "goliath.validate.Struct.ValidationErrors"
	expectedErrItems := 1
	if gErr.ErrorCode != expectedErrCode {
		t.Errorf("expect err code = %v, got = %v", expectedErrCode, gErr.ErrorCode)
	}
	if len(gErr.ErrorArgs.ValidationErrors) != expectedErrItems {
		t.Errorf("expect err items = %v, got = %v", expectedErrItems, len(gErr.ErrorArgs.ValidationErrors))
	}
}

func getGoliathErrorStruct(gotError goliath.Error) goliathErrorStruct {
	b, err := json.Marshal(gotError)
	if err != nil {
		panic(err)
	}
	var goliathError goliathErrorStruct
	json.Unmarshal(b, &goliathError)

	return goliathError
}
