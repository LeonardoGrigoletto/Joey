package main

import (
	"joey"
	"path/filepath"
	"runtime"
)

func setup() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get file path.")
	}
	dir := filepath.Dir(file)
	testCsvFilePath := filepath.Join(dir, "example_arithmetic.csv")
	return testCsvFilePath
}

func main() {
	dataframe, err := joey.NewFromCsv(setup())
	if err != nil {
		panic(err)
	}
	// Convert to int
	for _, column := range *dataframe.Columns() {
		column.Convert("int")
	}

	// Adding one column to another in the same Dataframe
	dataframe.Column("charge").Add(*dataframe.Column("walltime"))
	dataframe.Show()

	// Adding a new column to another
	newColumn := joey.Repeat(26, float64(50.0), "newColumn").Convert("int")
	dataframe.Column("charge").Add(newColumn)
	dataframe.Show()

	// Subtracting a new column from another
	newColumn = joey.Repeat(26, float64(10000), "newColumn").Convert("int")
	dataframe.Column("charge").Subtract(newColumn)
	dataframe.Show()

	// Multiplying a new column with another
	newColumn = joey.Repeat(26, float64(10000), "newColumn").Convert("int")
	dataframe.Column("charge").Multiply(newColumn)
	dataframe.Show()

}
