package ntime

import (
	"encoding/json"
	"github.com/rainmore/rest-api/tests"
	"github.com/rainmore/rest-api/utils/log"
	json2 "github.com/rainmore/rest-api/utils/njson"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	tests.SetupUnitTests()
	code := m.Run()
	tests.TearDownUnitTests()
	os.Exit(code)
}

func TestLocalDate_DateLocalNew(t *testing.T) {
	date := LocalDateNew(2025, 2, 14)

	expectedDate := time.Date(2025, 2, 14, 0, 0, 0, 0, time.Local)

	assert.Equal(t, LocalDate{expectedDate}, date)
}

func TestLocalDate_DateLocalCurrent(t *testing.T) {
	date := DateLocalCurrent()

	expectedDate := time.Now()

	assert.Equal(t, expectedDate.Year(), date.Year())
	assert.Equal(t, expectedDate.Month(), date.Month())
	assert.Equal(t, expectedDate.Day(), date.Day())
}

func TestLocalDate_DateLocalFromLISOStr(t *testing.T) {
	dateStr := "2025-02-14"
	date := LocalDateFromISOStr(dateStr)
	assert.Equal(t, LocalDateNew(2025, 2, 14), date)

	emptyDateStr := ""
	assert.Panics(t, func() { LocalDateFromISOStr(emptyDateStr) })

	invalidDateStr := "invalid"
	assert.Panics(t, func() { LocalDateFromISOStr(invalidDateStr) })
}

func TestLocalDate_DateLocalFromLocalStr(t *testing.T) {
	dateStr := "14/02/2025"
	date := LocalDateFromLocalStr(dateStr)
	assert.Equal(t, LocalDateNew(2025, 2, 14), date)

	emptyDateStr := ""
	assert.Panics(t, func() { LocalDateFromLocalStr(emptyDateStr) })

	invalidDateStr := "invalid"
	assert.Panics(t, func() { LocalDateFromLocalStr(invalidDateStr) })
}

func TestLocalDate_Year(t *testing.T) {
	dateStr := "2025-02-14"
	date := LocalDateFromISOStr(dateStr)
	assert.Equal(t, 2025, date.Year())
}

func TestLocalDate_Month(t *testing.T) {
	dateStr := "2025-02-14"
	date := LocalDateFromISOStr(dateStr)
	assert.Equal(t, time.Month(2), date.Month())
}

func TestLocalDate_Day(t *testing.T) {
	dateStr := "2025-02-14"
	date := LocalDateFromISOStr(dateStr)
	assert.Equal(t, 14, date.Day())
}

func TestLocalDate_String(t *testing.T) {
	dateLocal := LocalDateNew(2025, 2, 14)

	assert.Equal(t, "2025-02-14", dateLocal.String())
}

func TestLocalDate_LocalString(t *testing.T) {
	dateLocal := LocalDateNew(2025, 2, 14)

	assert.Equal(t, "14/02/2025", dateLocal.LocalString())
}

func TestLocalDate_IsZero(t *testing.T) {
	dateLocal := LocalDate{}
	assert.True(t, dateLocal.IsZero())

	dateLocalNotZero := LocalDateNew(2025, 2, 14)
	assert.False(t, dateLocalNotZero.IsZero())
}

//func TestFormatLocalToISO(t *testing.T) {
//	value := "24/09/2025"
//	result := DateFormatLocalToISO(value)
//	expectedStr := "2025-09-24"
//
//	assert.Equal(t, expectedStr, result)
//
//	value1 := ""
//	result1 := DateFormatLocalToISO(value1)
//	expectedStr1 := ""
//
//	assert.Equal(t, expectedStr1, result1)
//
//	value2 := "Omva;odasdf asdf "
//	result2 := DateFormatLocalToISO(value2)
//
//	assert.Equal(t, expectedStr1, result2)
//}

func TestLocalDate_MarshalJSON(t *testing.T) {
	date := LocalDateNew(2025, 2, 14)
	bytes, err := date.MarshalJSON()

	// string
	assert.Nil(t, err)
	assert.Equal(t, json2.JSONQuoteStr("2025-02-14"), string(bytes[:]))
}

func TestLocalDate_UnmarshalJSON(t *testing.T) {
	blob := `"2025-02-14"`
	var dateLocal LocalDate

	err := json.Unmarshal([]byte(blob), &dateLocal)
	assert.Nil(t, err)

	assert.Equal(t, LocalDateNew(2025, 2, 14), dateLocal)
}

func TestLocalDate_UnmarshalJSON_InvalidStr(t *testing.T) {
	blob := `"invalid"`
	var dateLocal LocalDate
	err := json.Unmarshal([]byte(blob), &dateLocal)

	assert.NotNil(t, err)
}

func TestLocalDate_UnmarshalJSON_Null(t *testing.T) {
	blob := `null`
	var dateLocal LocalDate
	err := json.Unmarshal([]byte(blob), &dateLocal)

	assert.NotNil(t, err)
}

type Foo struct {
	Bar int
}

func (foo Foo) foo1() {
	log.Logger().Infof(`foo by obj value: %q, location: %p`, &foo, foo)
	foo.Bar = foo.Bar + 100
	log.Logger().Infof(`foo by obj value: %q, location: %p`, &foo, foo)
}

func (foo *Foo) foo2() {
	log.Logger().Infof(`foo by ref value: %q, location: %p`, foo, *foo)
	foo.Bar = foo.Bar - 100
	log.Logger().Infof(`foo by ref value: %q, location: %p`, foo, *foo)
}

func TestLocalDate_UnmarshalJSON_Null111(t *testing.T) {
	foo1 := Foo{Bar: 100}
	foo2 := Foo{Bar: 100}

	assert.Equal(t, foo1, foo2)
	log.Logger().Infof(`foo1: %d`, foo1.Bar)
	log.Logger().Infof(`foo2: %d`, foo2.Bar)

	foo1.foo1()
	foo2.foo2()

	log.Logger().Infof(`foo1: %d`, foo1.Bar)
	log.Logger().Infof(`foo2: %d`, foo2.Bar)

}
