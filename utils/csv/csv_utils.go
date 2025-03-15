package csv

import (
	"encoding/csv"
	"github.com/rainmore/rest-api/utils/log"
	"github.com/rainmore/rest-api/utils/nio"
	"os"
)

type CSVData struct {
	HasHeader bool
	Header    []string
	Data      [][]string
}

func (d CSVData) Map() []map[string]string {
	if !d.HasHeader {
		panic("Data has no headers")
	}

	data := make([]map[string]string, len(d.Data))

	for i, row := range d.Data {
		rowMap := make(map[string]string)
		for j, key := range d.Header {
			rowMap[key] = row[j]
		}
		data[i] = rowMap
	}

	return data
}

func CSVReadByFilePath(filePath nio.Path, hasHeader bool) CSVData {
	f, err := os.Open(filePath.String())
	defer f.Close()
	if err != nil {
		log.Logger().Panicf(`Unable to read input file %q, err: %q`, filePath, err)
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Logger().Panicf(`Unable to read input file %q, err: %q`, filePath, err)
	}

	var csvData = CSVData{}
	csvData.HasHeader = hasHeader
	if hasHeader {
		csvData.Header = records[0]
		csvData.Data = records[1:len(records)]
	} else {
		csvData.Data = records
	}

	f.Close()
	return csvData
}
