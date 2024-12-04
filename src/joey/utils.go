package joey

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func createRow(record []string) (Row, error) {
	newRow := Row{data: make([]Cell, len(record))}
	if reflect.TypeOf(record[0]) == reflect.TypeOf("string") {
		for i, v := range record {
			newRow.data[i] = StrCell{data: v}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(int(10)) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			intCell := int(cell)
			newRow.data[i] = IntCell{data: intCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(int8(10)) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			intCell := int8(cell)
			newRow.data[i] = Int8Cell{data: intCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(int16(10)) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			intCell := int16(cell)
			newRow.data[i] = Int16Cell{data: intCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(int32(10)) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			intCell := int32(cell)
			newRow.data[i] = Int32Cell{data: intCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(int64(10)) {
		for i, v := range record {
			cell, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			intCell := int64(cell)
			newRow.data[i] = Int64Cell{data: intCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(float32(10.0)) {
		for i, v := range record {
			cell, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			floatCell := float32(cell)
			newRow.data[i] = Float32Cell{data: floatCell}
		}
		return newRow, nil
	} else if reflect.TypeOf(record[0]) == reflect.TypeOf(float64(10.0)) {
		for i, v := range record {
			cell, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic("Error converting string to int.")
			}
			newRow.data[i] = Float64Cell{data: cell}
		}
		return newRow, nil
	} else {
		return Row{data: []Cell{}}, errors.New("the specified type is not known")
	}
}

func NewFromCsv(path string) (Dataframe, error) {
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
	columns := make([]Column, len(header))
	for i := range header {
		var col Column
		columns[i] = col.New(len(records))
	}

	for i, _ := range header {
		column := make([]string, len(records))
		for j, _ := range records {
			column[j] = records[j][i]
		}
	}
	for i, record := range records {
		row, err := createRow(record)
		if err != nil {
			return Dataframe{}, err
		}
		rows[i] = row

		for j := range record {
			columns[j].data[i] = record[j]
		}
	}

	dataframe := Dataframe{
		headers: header,
		rows:    rows,
		columns: columns,
	}
	return dataframe, nil
}
