package date_time

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeOnly struct {
	T time.Time
}

func (timeOnly TimeOnly) MarshalJSON() ([]byte, error) {
	formatted := timeOnly.T.Format("15:04:05")
	return json.Marshal(formatted)
}

func (timeOnly *TimeOnly) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}
	t, err := time.Parse("15:04:05", s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	*timeOnly = TimeOnly{t}
	return nil
}
