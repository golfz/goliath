package goliath

import "encoding/json"

type ErrorPresenterInterface interface {
	PresentError(err Error)
}

type ErrorPresenter struct {
	View Viewer
}

func (p *ErrorPresenter) PresentError(err Error) {
	e := getGoliathError(err)
	p.View.Write(e.Status, err)
}

func getGoliathError(errInf Error) goliathError {
	b, _ := json.Marshal(errInf)
	var gErr goliathError
	_ = json.Unmarshal(b, &gErr)
	return gErr
}
