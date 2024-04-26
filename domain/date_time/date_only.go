package date_time

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateOnly struct {
	T time.Time
}

func (timeOnly DateOnly) MarshalJSON() ([]byte, error) {
	formatted := timeOnly.T.Format("2006-01-02")
	return json.Marshal(formatted)
}

func (timeOnly *DateOnly) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("failed to parse date: %w", err)
	}
	*timeOnly = DateOnly{t}
	return nil
}
