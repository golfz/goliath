package json

import "github.com/golfz/goliath/cleanarch/data/viewmodel"

type Writer interface {
	Write(header viewmodel.HttpHeader, body interface{})
}
