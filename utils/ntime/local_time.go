package ntime

import (
	"encoding/json"
	"time"
)

const (
	timeFormatISO   = time.TimeOnly
	timeFormatLocal = "02/01/2006"
)

type LocalTime struct {
	Time time.Time
}

func LocalTimeNew(hour, min, sec, nsec int) LocalTime {
	return LocalTime{time.Date(0, 0, 0, hour, min, sec, nsec, time.Local)}
}

func LocalTimeCurrent(hour, min, sec, nsec int) LocalTime {
	now := time.Now()
	return LocalTimeNew(now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
}

func (t LocalTime) Hour() int {
	return t.Time.Hour()
}

func (t LocalTime) Minute() int {
	return t.Time.Minute()
}

func (t LocalTime) Second() int {
	return t.Time.Second()
}

func (t LocalTime) Nanosecond() int {
	return t.Time.Nanosecond()
}

func (t *LocalTime) IsZero() bool {
	return t.Time.IsZero()
}

func (t LocalTime) String() string {
	return t.Time.Format(timeFormatISO)
}

func (t LocalTime) LocalString() string {
	return t.Time.Format(timeFormatLocal)
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	} else {
		return json.Marshal(t.String())
	}
}

func (t *LocalTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	result, err := time.ParseInLocation(dateFormatISO, s, time.Local)
	if err == nil {
		*t = LocalTime{result}
	}

	return err
}
