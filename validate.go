package goliath

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golfz/goliath/x/data/output"
	"reflect"
	"runtime/debug"
	"time"
)

func Validate(v interface{}) output.GoliathError {
	if err := validator.New().Struct(v); err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		for _, e := range err.(validator.ValidationErrors) {
			s := string(debug.Stack())
			fmt.Println(s)
			fmt.Println("-----------------")
			eStr := fmt.Sprint(e)
			fmt.Println("eStr: ", eStr)
			fmt.Println(e.Namespace())
			fmt.Println(e.ActualTag())
			fmt.Println(e.Field())
			fmt.Println(e.Tag())
			fmt.Println(e.Kind())
			fmt.Println(e.Param())
			fmt.Println(e.StructField())
			fmt.Println(e.StructNamespace())
			fmt.Println(e.Type())
			fmt.Println(e.Value())
			fmt.Println("+++++++++++++++++++")
			t := reflect.TypeOf(v)
			n, _ := t.FieldByName(e.Field())
			fmt.Println(t)
			fmt.Println(n.Tag.Get("json"))
			return &output.Error{
				Status:   0,
				Time:     time.Now(),
				Type:     "",
				Code:     "",
				Error:    "",
				Message:  "",
				ErrorDev: output.ErrorDev{},
			}
		}
	}

	return nil
}
