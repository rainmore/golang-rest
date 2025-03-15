package csv

import (
	"github.com/rainmore/rest-api/tests"
	"github.com/rainmore/rest-api/utils/log"
	"github.com/rainmore/rest-api/utils/nio"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tests.SetupUnitTests()
	code := m.Run()
	tests.TearDownUnitTests()
	os.Exit(code)
}

func TestProcessCSV(t *testing.T) {
	t.Skip("Could not load resources")
	csvFilePath := nio.ResPath().Resolve("FT-529", "ae_accumulation_unit_price.csv")
	data := CSVReadByFilePath(csvFilePath, true)

	log.Logger().Infof(`result header: %q, data length: %d`, data.Header, len(data.Data))
}

func TestCSVData_Map(t *testing.T) {
	t.Skip("Could not load resources")
	csvFilePath := nio.ResPath().Resolve("FT-529", "ae_accumulation_unit_price.csv")
	data := CSVReadByFilePath(csvFilePath, true)

	mapData := data.Map()
	for idx, row := range mapData {
		log.Logger().Infof(`row idx: %d map: %q`, idx, row)
	}
}
