package validate

import (
	"encoding/json"
	"testing"
	"time"
)

// user is input
type user struct {
	Name    *string `validate:"required,min=2"`
	Age     uint8   `validate:"gte=18,lte=130"`
	Email   string  `validate:"required,email"`
	Address address
}

type address struct {
	Zip uint16 `validate:"required,gte=10000,lte=99999"`
}

// GoliathErrorStruct is a struct of goliath.Error
type GoliathErrorStruct struct {
	Status    int            `json:"status"`
	Message   string         `json:"message"`
	Time      time.Time      `json:"time"`
	LogID     string         `json:"log_id"`
	ErrorCode string         `json:"error_code"`
	ErrorArgs ErrorArguments `json:"error_args"`
	ErrorDev  ErrorDev       `json:"error_dev"`
	err       error
}

type ErrorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

type ErrorArguments struct {
	ValidationErrors []ValidationErrorCase `json:"validation_errors"`
}

type ValidationErrorCase struct {
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
	expectedError GoliathErrorStruct
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
	b, err := json.Marshal(gotErr)
	if err != nil {
		panic(err)
	}
	var goliathError GoliathErrorStruct
	json.Unmarshal(b, &goliathError)

	expectedErrCode := "goliath.validate.Struct.InvalidValidationError"
	if goliathError.ErrorCode != expectedErrCode {
		t.Errorf("Expected %v, got %v", expectedErrCode, goliathError.ErrorCode)
	}
}
