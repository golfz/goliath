package timeutils

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	now := time.Now()
	tu := New(now)
	timeNew := tu.Time

	got := timeNew.Format("2006-01-02T15:04:05.999999999Z07:00")
	want := now.Format("2006-01-02T15:04:05.999999999Z07:00")

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestChangeTimezone(t *testing.T) {
	rfc3339 := "2006-01-02T15:04:05Z07:00"
	thaiTime, _ := time.Parse(rfc3339, "2021-11-19T09:46:53+07:00")

	tu := New(thaiTime)
	newTime, _ := tu.ChangeTimezone("-10:00")

	want := "2021-11-19T09:46:53-10:00"
	got := newTime.Format(rfc3339)

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestChangeTimezone_Error(t *testing.T) {
	rfc3339 := "2006-01-02T15:04:05Z07:00"
	thaiTime, _ := time.Parse(rfc3339, "2021-11-19T09:46:53+07:00")

	tu := New(thaiTime)
	_, err := tu.ChangeTimezone("+AB:00")
	if err == nil {
		t.Errorf("want error, got nil")
	}
}
