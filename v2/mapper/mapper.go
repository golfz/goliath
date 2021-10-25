package mapper

import (
	"github.com/golfz/goliath/v2"
	"github.com/jinzhu/copier"
	"net/http"
)

type structMapper struct {
	from interface{}
}

func From(v interface{}) *structMapper {
	return &structMapper{from: v}
}

func (s *structMapper) To(v interface{}) goliath.Error {
	if err := copier.Copy(v, s.from); err != nil {
		errStatus := http.StatusInternalServerError
		errCode := "goliath.mapper.To.MapperError"
		errMessage := "goliath mapper error"
		return goliath.NewError(errStatus, errCode, nil, err, "", errMessage)
	}

	return nil
}
