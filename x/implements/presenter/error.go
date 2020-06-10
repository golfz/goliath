package presenter

import (
	"github.com/golfz/goliath/x/data/output"
	"github.com/golfz/goliath/x/data/viewmodel"
	"github.com/golfz/goliath/x/interfaces/view/restful/json"
)

type ErrorPresenter struct {
	View json.Writer
}

func (p *ErrorPresenter) PresentError(err output.GoliathError) {
	h := viewmodel.HttpHeader{
		StatusCode:    err.Errors().Status,
		Authorization: "",
		ContentType:   "",
	}
	p.View.Write(h, err)
}
