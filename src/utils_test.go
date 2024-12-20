package joey

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func setup() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get file path.")
	}
	dir := filepath.Dir(file)
	testCsvFilePath := filepath.Join(dir, "test.csv")
	return testCsvFilePath
}

func ReadCsv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Could not open CSV test file at: " + path)
	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("Could not read CSV test file at: " + path)
	}
	return records, nil
}

func TestNewFromCsvShouldCreateANewDataframe(t *testing.T) {
	filePath := setup()
	records, err := ReadCsv(filePath)
	if err != nil {
		t.Fatal(err)
	}
	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	correctHeaders := records[0]
	correctRecords := records[1:]
	// validate headers
	for i, header := range dataframe.headers {
		if correctHeaders[i] != header {
			t.Fatalf("Dataframe Headers does not match. %s != %s", header, correctHeaders[i])
		}
		if dataframe.columns[i].name != header {
			t.Fatalf("Dataframe Headers does not match. %s != %s", header, correctHeaders[i])
		}
	}
	// validate columns
	for i, column := range dataframe.columns {
		for j, cell := range column.Data {
			if correctRecords[j][i] != cell.GetRawData() {
				t.Fatalf("Record %s from CSV != %s from Dataframe", correctRecords[j][i], cell.GetRawData())
			}
		}
	}

}
