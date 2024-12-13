package joey

import (
	"encoding/csv"
	"fmt"
	"joey/cells"
	"os"
)

func createDefaultCell(cell string) cells.Cell {
	strCell := cells.StrCell{Data: cell}
	return &strCell
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

	columns := make([]Column, len(header))
	for i := range header {
		var col Column
		columns[i] = col.New(len(records), header[i])
		for j := range records {
			columns[i].Data[j] = createDefaultCell(records[j][i])
		}
	}
	dataframe := Dataframe{
		headers: header,
		columns: columns,
	}
	return dataframe, nil
}
