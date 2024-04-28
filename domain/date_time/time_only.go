package date_time

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	apiTimeOnlyLayout = "15:04:05"
)

type TimeOnly struct {
	T time.Time
}

func (timeOnly TimeOnly) MarshalJSON() ([]byte, error) {
	formatted := timeOnly.T.Format(apiTimeOnlyLayout)
	return json.Marshal(formatted)
}

func (timeOnly *TimeOnly) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}
	t, err := time.Parse(apiTimeOnlyLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	*timeOnly = TimeOnly{t}
	return nil
}
