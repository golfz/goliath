package goliath

import (
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

func (s *structMapper) To(v interface{}) error {
	if err := copier.Copy(v, s.from); err != nil {
		return err
	}
	return nil
}
