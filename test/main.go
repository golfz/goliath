package test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golfz/goliath/cleanarch/data/viewmodel"
	"github.com/golfz/goliath/utils/structmapper"
	"github.com/golfz/goliath/utils/validator"
)

type TestStruct struct {
	ID    int          `json:"id"     validate:"min=1,max=10"`
	Email string       `json:"email"  validate:"required,email"`
	My    NameStruct   `json:"my"     validate:"required"`
	Name  []NameStruct `json:"name"   validate:"required,dive,required"`
}

type NameStruct struct {
	FirstName string         `json:"first_name" validate:"min=5,max=10"`
	LastName  LastNameStruct `json:"last_name"  validate:"min=5,max=10"`
}

type LastNameStruct struct {
	FamilyName string `json:"family_name" validate:"min=5,max=10"`
	OwnName    string `json:"own_name" validate:"min=5,max=10"`
}

func main() {
	s := TestStruct{
		ID:    9,
		Email: "fasdf@asdf.asfd",
		My: NameStruct{
			FirstName: "asdfsdff",
			LastName: LastNameStruct{
				FamilyName: "asdfsdff",
				OwnName:    "asdfsdff",
			},
		},
		Name: []NameStruct{
			{
				FirstName: "asdfsdff",
				LastName: LastNameStruct{
					FamilyName: "asdfsdff",
					OwnName:    "asdfsdff",
				},
			},
			{
				FirstName: "asdfsdff",
				LastName: LastNameStruct{
					FamilyName: "",
					OwnName:    "asdfsdff",
				},
			},
		},
	}

	if err := validator.Validate(s); err != nil {
		e := viewmodel.Error{}
		if err := structmapper.NewStructMapper().From(err).To(&e); err != nil {
			fmt.Println(err)
		}
		prettyJSON, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", prettyJSON)
	}
}
