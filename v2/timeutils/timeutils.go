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

func New(time time.Time) *TimeUtils {
	return &TimeUtils{Time: time}
}

// ChangeTimezone use for change timezone only without change time.
// (tz format: "Z07:00"),
// e.g. 10:00+07:00 to 10:00-05:00
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
