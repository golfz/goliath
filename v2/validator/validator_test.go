package validator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type user struct {
	Name  *string `validate:"required,min=1"`
	Age   uint8  `validate:"gte=18,lte=130"`
	Email string `validate:"required,email"`
	Address Address
}

type Address struct {
	Zip uint16 `validate:"required,gte=10000,lte=99999"`
}

func TestStruct(t *testing.T) {
	name := "Marc"
	u := user{
		Name:  &name,
		Age:   250,
		Email: "abc@gmail.com",
		Address: Address{Zip: 20},
	}

	err := Struct(u)
	if err != nil {
		fmt.Printf("%#v \n", err)

		b, err := json.MarshalIndent(err, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}
