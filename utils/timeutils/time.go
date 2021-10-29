package timeutils

import (
	"fmt"
	"github.com/golfz/goliath/cleanarch/data/output"
	"net/http"
	"runtime/debug"
	"time"
)

type Time struct {
	Time time.Time
}

func TimeUtil(time time.Time) *Time {
	return &Time{Time: time}
}

// ChangeTimezone use for change timezone
// (tz format: "Z07:00")
func (t *Time) ChangeTimezone(tz string) (time.Time, output.GoliathError) {
	localTimeString := fmt.Sprintf("%s%s", t.Time.Format("2006-01-02T15:04:05"), tz)
	localTime, err := time.Parse(time.RFC3339, localTimeString)

	if err != nil {
		return time.Time{}, &output.Error{
			Status:  http.StatusInternalServerError,
			Time:    time.Now(),
			Type:    "utils",
			Code:    "utils-TimeUtil.ChangeTimezone",
			Error:   "Cannot convert time correctly",
			Message: "Cannot convert time correctly",
			ErrorDev: output.ErrorDev{
				Error:      err.Error(),
				Stacktrace: string(debug.Stack()),
			},
		}
	}

	return localTime, nil
}
