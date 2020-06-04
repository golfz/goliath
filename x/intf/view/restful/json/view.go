package json

import "github.com/golfz/goliath/x/data/viewmodel"

type Writer interface {
	Write(header viewmodel.HttpHeader, body interface{})
}



