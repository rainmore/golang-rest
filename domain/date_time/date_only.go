package date_time

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	apiDateOnlyLayout = "2006-01-02"
)

type DateOnly struct {
	T time.Time
}

func (timeOnly DateOnly) MarshalJSON() ([]byte, error) {
	formatted := timeOnly.T.Format(apiDateOnlyLayout)
	return json.Marshal(formatted)
}

func (timeOnly *DateOnly) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}
	t, err := time.Parse(apiDateOnlyLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse date: %w", err)
	}
	*timeOnly = DateOnly{t}
	return nil
}
