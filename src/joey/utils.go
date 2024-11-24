package joey

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func createAxis(record []string) (Row, Column, error) {
	newRow := Row{data: make([]Cell, len(record))}
	var newColumn Column

	if reflect.TypeOf(record[0]) == reflect.TypeOf("string") {
		for i, v := range record {
			newRow.data[i] = StrCell{data: v}
		}
		*newColumn.data = append(*newColumn.data, newRow.data[0])
		return newRow, newColumn, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(10) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			newRow.data[i] = IntCell{data: cell}
		}
		return newRow, newColumn, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(10.0) {
		for i, v := range record {
			cell, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			newRow.data[i] = FloatCell{data: cell}
		}
		return newRow, newColumn, nil
	} else {
		return nil, nil, errors.New("Unknwon types.")
	}
}

func NewFromCsv(path string) Dataframe {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to open the file: %s %s", path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to open the file: %s %s", path, err)
	}
	header := records[0]
	records = records[1:]

	rows := make([]Row, len(records))
	for i, record := range records {
		rows[i] = createAxis(record)
	}

	dataframe := Dataframe{
		headers: header,
		rows:    rows,
	}
	return dataframe
}
