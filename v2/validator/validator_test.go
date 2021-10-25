package validator

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type user struct {
	Name    *string `validate:"required,min=1"`
	Age     uint8   `validate:"gte=18,lte=130"`
	Email   string  `validate:"required,email"`
	Address address
}

type address struct {
	Zip uint16 `validate:"required,gte=10000,lte=99999"`
}

type validateError struct {
	Status    int          `json:"status"`
	Message   string       `json:"message"`
	Time      time.Time    `json:"time"`
	LogID     string       `json:"log_id"`
	ErrorCode string       `json:"error_code"`
	ErrorArgs errArguments `json:"error_args"`
	ErrorDev  errorDev     `json:"error_dev"`
	err       error
}

type errArguments struct {
	ValidationErrors []validationError `json:"validation_errors"`
}

type validationError struct {
	Key           string      `json:"key"`
	ActualValue   interface{} `json:"actual_value"`
	Rule          string      `json:"rule"`
	ExpectedValue interface{} `json:"expected_value"`
	Message       string      `json:"message"`
}

type errorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

func TestStruct(t *testing.T) {
	name := "Marc"
	u := user{
		Name:    &name,
		Age:     250,
		Email:   "abc@gmail.com",
		Address: address{Zip: 20},
	}

	err := Struct(u)
	if err != nil {
		b, err := json.MarshalIndent(err, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

		var errStruct validateError
		json.Unmarshal(b, &errStruct)

		fmt.Println("------------------------------")
		fmt.Printf("%#v", errStruct)
	}
}
