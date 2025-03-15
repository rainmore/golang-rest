package ntime

import (
	"encoding/json"
	"time"
)

const (
	dateFormatISO   = time.DateOnly
	dateFormatLocal = "02/01/2006"
)

type LocalDate struct {
	Time time.Time
}

func LocalDateNew(year int, month time.Month, day int) LocalDate {
	return LocalDate{time.Date(year, month, day, 0, 0, 0, 0, time.Local)}
}

func DateLocalCurrent() LocalDate {
	now := time.Now()
	return LocalDateNew(now.Year(), now.Month(), now.Day())
}

func LocalDateFromISOStr(str string) LocalDate {
	date, err := time.ParseInLocation(dateFormatISO, str, time.Local)
	if err != nil {
		panic(err)
	}
	return LocalDate{date}
}

func LocalDateFromLocalStr(str string) LocalDate {
	date, err := time.ParseInLocation(dateFormatLocal, str, time.Local)
	if err != nil {
		panic(err)
	}
	return LocalDate{date}
}

func (d LocalDate) Year() int {
	return d.Time.Year()
}

func (d LocalDate) Month() time.Month {
	return d.Time.Month()
}

func (d LocalDate) Day() int {
	return d.Time.Day()
}

func (d LocalDate) String() string {
	return d.Time.Format(dateFormatISO)
}

func (d LocalDate) LocalString() string {
	return d.Time.Format(dateFormatLocal)
}

func (d *LocalDate) IsZero() bool {
	return d.Time.IsZero()
}

func (d LocalDate) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	} else {
		return json.Marshal(d.String())
	}
}

func (d *LocalDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	result, err := time.ParseInLocation(dateFormatISO, s, time.Local)
	if err == nil {
		*d = LocalDate{result}
	}

	return err
}
