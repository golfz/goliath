package goliath

import "encoding/json"

type ErrorPresenterInterface interface {
	PresentError(err Error)
}

type ErrorPresenter struct {
	View Viewer
	Ctx  Goliath
}

func (p *ErrorPresenter) PresentError(err Error) {
	err.SetLogID(p.Ctx.LogID())
	status := getGoliathError(err).Status
	p.View.Write(status, err)
}

func getGoliathError(errInf Error) goliathError {
	b, _ := json.Marshal(errInf)
	var gErr goliathError
	_ = json.Unmarshal(b, &gErr)
	return gErr
}
