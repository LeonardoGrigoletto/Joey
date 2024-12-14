package main

import (
	"fmt"
	"joey"
	"path/filepath"
	"runtime"
	"time"
)

func setup() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get file path.")
	}
	dir := filepath.Dir(file)
	testCsvFilePath := filepath.Join(dir, "benchmarking_data.csv")
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

	fmt.Println("----- Benchmarking Adding Columns -----")
	// Single Proc
	start_time := time.Now()
	joey.N_PROC = 1
	dataframe.Column("charge").Add(*dataframe.Column("walltime"))
	duration := time.Since(start_time)
	fmt.Printf("Single Proc Add Duration: %v seconds.\n", duration)

	// Multi Proc
	start_time = time.Now()
	joey.N_PROC = 8
	dataframe.Column("charge").Add(*dataframe.Column("walltime"))
	duration = time.Since(start_time)
	fmt.Printf("Multi Proc Add Duration: %v seconds.\n\n", duration)

	fmt.Println("----- Benchmarking Subtracting Columns -----")
	// Subtracting a new column from another
	// Single Proc
	start_time = time.Now()
	joey.N_PROC = 1
	dataframe.Column("charge").Subtract(*dataframe.Column("walltime"))
	duration = time.Since(start_time)
	fmt.Printf("Single Proc Subtract Duration: %v seconds.\n", duration)

	// Multi Proc
	start_time = time.Now()
	joey.N_PROC = 8
	dataframe.Column("charge").Subtract(*dataframe.Column("walltime"))
	duration = time.Since(start_time)
	fmt.Printf("Multi Proc Subtract Duration: %v seconds.\n\n", duration)

	fmt.Println("----- Benchmarking Multiplying Columns -----")
	// Subtracting a new column from another
	// Single Proc
	start_time = time.Now()
	joey.N_PROC = 1
	dataframe.Column("charge").Multiply(*dataframe.Column("walltime"))
	duration = time.Since(start_time)
	fmt.Printf("Single Proc Multiply Duration: %v seconds.\n", duration)

	// Multi Proc
	start_time = time.Now()
	joey.N_PROC = 8
	dataframe.Column("charge").Multiply(*dataframe.Column("walltime"))
	duration = time.Since(start_time)
	fmt.Printf("Multi Proc Multiply Duration: %v seconds.\n\n", duration)

}
