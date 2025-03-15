package njson

import (
	"encoding/json"
)

func JSONQuoteStr(str string) string {
	result, err := json.Marshal(str)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func JSONAddress[T any](t T) *T { return &t }

func ToBytes[T any](t T) []byte {
	result, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}
	return result
}
