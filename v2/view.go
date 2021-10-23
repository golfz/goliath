package goliath

type Viewer interface {
	Write(status int, data interface{})
}
