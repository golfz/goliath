package mapper

import (
	"testing"
)

type structA struct {
	Name string
	Age  int
}

type structB struct {
	Name string
	Age  int
}

func (b structB) Firstname() string {
	return b.Name
}

type structC struct {
	Firstname string
	Age       int
}

func TestFrom_MapSameField(t *testing.T) {
	a := structA{
		Name: "Tom",
		Age:  35,
	}
	b := structB{}

	if err := From(&a).To(&b); err != nil {
		t.Errorf("want = %v, got = %v", nil, err)
	}
}

func TestFrom_MapMethodToField(t *testing.T) {
	b := structB{
		Name: "Tom",
		Age:  35,
	}
	c := structC{}

	if err := From(&b).To(&c); err != nil {
		t.Errorf("want = %v, got = %v", nil, err)
	}

	if c.Firstname != b.Name {
		t.Errorf("want = %v, got = %v", b.Name, c.Firstname)
	}
}

func TestFrom_NotSameFieldName(t *testing.T) {
	a := structA{
		Name: "Tom",
		Age:  35,
	}
	c := structC{}

	if err := From(&a).To(&c); err != nil {
		t.Errorf("want = %v, got = %v", nil, err)
	}

	if c.Firstname != "" {
		t.Errorf("want = %v, got = %v", "", c.Firstname)
	}
}

func TestFrom_Error(t *testing.T) {
	a := structA{
		Name: "Tom",
		Age:  35,
	}
	b := structB{}

	if err := From(a).To(b); err == nil {
		t.Errorf("want error, got = %v", err)
	}
}
