
# Joey Example - Dataframe Manipulation

This repository contains a simple example of how to use the `joey` library to work with DataFrames in Go. The example demonstrates how to load data from a CSV file, perform type conversion operations, and execute arithmetic operations to a DataFrame.


### Code Objective

1. **Load a CSV file**:
   The code starts by loading a CSV file (`example_arithmetic.csv`) from the same directory as the Go file. This file contains data that will be read and loaded into a DataFrame.

2. **Convert Data**:
   The example converts all values in the columns of the DataFrame to integers.

3. **Add Columns**:
   It then adds two columns: one by summing the values of two existing columns (`charge` and `walltime`), and another by creating a new column programmatically (with constant values).

4. **Display the Data**:
   The code displays the updated DataFrame after the data manipulation operations.

## Dependencies

Ensure you have the following dependencies installed to run the code:

- Go 1.x or higher
- The `joey` library (assuming it is an internal or imported library for DataFrame manipulation)

## Code Explanation

### `setup()` Function

The `setup()` function uses the `runtime` package to get the directory of the Go file in execution and generates the absolute path to the input CSV file (`example_sum_columns.csv`).

```go
func setup() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not get file path.")
	}
	dir := filepath.Dir(file)
	testCsvFilePath := filepath.Join(dir, "example_sum_columns.csv")
	return testCsvFilePath
}
```

### Main Function

In the main function, the code loads the data from a CSV file into a DataFrame using `joey.NewFromCsv(setup())`. Then, it performs the following operations:

1. **Convert to Integer**: It converts the data in all the columns to integers.

```go
for _, column := range *dataframe.Columns() {
	column.Convert("int")
}
```

2. **Add Columns**:
   - **Add Existing Columns**: The `charge` column is added to the `walltime` column.
   - **Add New Column**: A new column, `newColumn`, is created with 26 values of `50.0`, converted to integers, and added to the `charge` column.

```go
dataframe.Column("charge").Add(*dataframe.Column("walltime"))
```

3. **Display the DataFrame**:
   The DataFrame is displayed after each operation using the `Show()` method.

```go
dataframe.Show()
```

## Example CSV File

The `example_arithmetic.csv` file should contain structured data with at least two columns (`charge` and `walltime`) for the example to work correctly. The CSV format may look like this:

```
charge,walltime
100,5
150,10
200,15
```

## Conclusion

This example demonstrates how to load data from a CSV file, manipulate it in a DataFrame, and perform operations like type conversion and adding columns in Go using the `joey` library. The code serves as a foundation for more complex operations on DataFrames, such as aggregations, transformations, and data analysis.