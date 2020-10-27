package test

type TestStruct struct {
	Id    int          `json:"id"     validate:"min=1,max=10"`
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

//func main() {
//	s := TestStruct{
//		Id:    9,
//		Email: "fasdf@asdf.asfd",
//		My: NameStruct{
//			FirstName: "asdfsdff",
//			LastName: LastNameStruct{
//				FamilyName: "asdfsdff",
//				OwnName:    "asdfsdff",
//			},
//		},
//		Name: []NameStruct{
//			{
//				FirstName: "asdfsdff",
//				LastName: LastNameStruct{
//					FamilyName: "asdfsdff",
//					OwnName:    "asdfsdff",
//				},
//			},
//			{
//				FirstName: "asdfsdff",
//				LastName: LastNameStruct{
//					FamilyName: "",
//					OwnName:    "asdfsdff",
//				},
//			},
//		},
//	}
//
//	if err := goliath.Validate(s); err != nil {
//		e := viewmodel.Error{}
//		if err := goliath.NewStructMapper().From(err).To(&e); err != nil {
//			fmt.Println(err)
//		}
//		prettyJson, err := json.MarshalIndent(e, "", "  ")
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%s\n", prettyJson)
//	}
//}
