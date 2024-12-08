package main

import "joey/joey"

func main() {
	dataframe, err := joey.NewFromCsv("./example_sum_columns.csv")
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
	data := make([]joey.Cell, 26)
	for i := range data {
		data[i] = &joey.IntCell{Data: 50}
	}
	newColumn := joey.Column{Data: data}
	dataframe.Column("charge").Add(newColumn)
	dataframe.Show()

}
