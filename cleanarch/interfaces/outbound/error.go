package outbound

import "github.com/golfz/goliath/cleanarch/data/output"

type ErrorPresenter interface {
	PresentError(err output.GoliathError)
}