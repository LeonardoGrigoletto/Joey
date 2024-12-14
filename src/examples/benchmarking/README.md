
# Joey Library - Examples and Benchmarking

This repository contains an example showcasing how to use the **Joey** library for column-based operations on dataframes. The included example demonstrates the performance difference between single-threaded and multi-threaded operations for adding, subtracting, and multiplying columns.

## Example Overview

### File: `examples/benchmarking_multiproc.go`

This Go program benchmarks various arithmetic operations on columns of a dataframe loaded from a CSV file. The operations are tested using both single-threaded (1 process) and multi-threaded (8 processes) configurations.

### Key Features
- **Single-threaded operations**: Perform column operations using a single worker process.
- **Multi-threaded operations**: Perform column operations in parallel using multiple worker processes.
- **Flexible benchmarking**: Measure the execution time for `Add`, `Subtract`, and `Multiply` operations.

### How It Works
1. The CSV file is loaded into a dataframe using the Joey library.
2. Columns in the dataframe are converted to integer type for arithmetic operations.
3. Each operation is benchmarked in both single-threaded and multi-threaded configurations.
4. Results are printed to the console.

---

## Code Details

### Main Code: `benchmarking_multiproc.go`
```go
// Benchmark adding columns
joey.N_PROC = 1 // Single-threaded
dataframe.Column("charge").Add(*dataframe.Column("walltime"))

// Benchmark subtracting columns
joey.N_PROC = 8 // Multi-threaded
dataframe.Column("charge").Subtract(*dataframe.Column("walltime"))

// Benchmark multiplying columns
dataframe.Column("charge").Multiply(*dataframe.Column("walltime"))
```

### Helper Functions
- **`setup()`**: Configures the path to the CSV file dynamically based on the current file's location.

---

## Requirements

- **Go**: Ensure you have Go installed on your system.
- **Joey Library**: This example depends on the Joey library, which must be installed and available in your Go workspace.

---

## Running the Example

1. Place your CSV file in the same directory as the example.
2. Update the `benchmarking_data.csv` file path in the `setup()` function, if necessary.
3. Run the example using:

   ```sh
   go run examples/benchmarking_multiproc.go
   ```

---

## Expected Output

You will see benchmarking results for each operation printed to the console, such as:

```
----- Benchmarking Adding Columns -----
Single Proc Add Duration: 0.123456 seconds.
Multi Proc Add Duration: 0.012345 seconds.

----- Benchmarking Subtracting Columns -----
Single Proc Subtract Duration: 0.234567 seconds.
Multi Proc Subtract Duration: 0.023456 seconds.

----- Benchmarking Multiplying Columns -----
Single Proc Multiply Duration: 0.345678 seconds.
Multi Proc Multiply Duration: 0.034567 seconds.
```

---

## Notes
- Multi-threaded performance may vary depending on your hardware and the size of the dataset.
- The number of processes used for multi-threading is configurable via `joey.N_PROC`.

Enjoy using the Joey library!