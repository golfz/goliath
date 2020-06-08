package outbound

import "github.com/golfz/goliath/x/data/output"

type ErrorPresenter interface {
	PresentError(err output.GoliathError)
}