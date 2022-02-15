package timeutils

import (
	"fmt"
	"github.com/golfz/goliath/v2"
	"net/http"
	"time"
)

type TimeUtils struct {
	Time time.Time
}

// New return *TimeUtils data
func New(time time.Time) *TimeUtils {
	return &TimeUtils{Time: time}
}

// ChangeTimezone is used to change timezone information of a time without changing other information for that time.
// (time format: "Z07:00"),
// e.g. January 1, 2021 12:15:30 +07:00 to 2021 12:15:30 -05:00
func (t *TimeUtils) ChangeTimezone(tz string) (time.Time, goliath.Error) {
	localTimeString := fmt.Sprintf("%s%s", t.Time.Format("2006-01-02T15:04:05"), tz)
	localTime, err := time.Parse(time.RFC3339, localTimeString)

	if err != nil {
		errStatus := http.StatusInternalServerError
		errCode := "goliath.timeutils.ChangeTimezone.ParsingError"
		logID := ""
		errMessage := "parsing time to new timezone error"
		return time.Time{}, goliath.NewError(errStatus, errCode, nil, err, logID, errMessage)
	}

	return localTime, nil
}
