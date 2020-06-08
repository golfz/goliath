package outbound

import "github.com/golfz/goliath/x/data/output"

type Error interface {
	PresentError(err output.GoliathError)
}