package structmapper

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/golfz/goliath/cleanarch/data/output"
	"github.com/jinzhu/copier"
)

type structMapper struct {
	from interface{}
}

func NewStructMapper() *structMapper {
	return &structMapper{}
}

func (s *structMapper) From(v interface{}) *structMapper {
	s.from = v
	return s
}

func (s *structMapper) To(v interface{}) output.GoliathError {
	if err := copier.Copy(v, s.from); err != nil {
		stack := string(debug.Stack())
		errorStr := fmt.Sprint(err)
		return &output.Error{
			Status:  http.StatusInternalServerError,
			Time:    time.Now(),
			Type:    "internal",
			Code:    "internal-mapstruct",
			Error:   "Something wrong",
			Message: "Something wrong (cannot mapping data).",
			ErrorDev: output.ErrorDev{
				Error:      errorStr,
				Stacktrace: stack,
			},
		}
	}
	return nil
}
