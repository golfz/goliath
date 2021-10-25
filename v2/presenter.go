package goliath

type ErrorPresenter interface {
	PresentError(err Error)
}
