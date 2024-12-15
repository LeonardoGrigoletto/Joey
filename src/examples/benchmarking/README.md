
W# Joey Library - Examples and Benchmarking

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

### Go Code Example: `benchmarking_multiproc.go`
```go
package main

import (
    "fmt"
    "joey"
    "time"
)

func main() {
    dataframe := joey.ReadCSV("benchmarking_data.csv")

    // Benchmark adding columns
    joey.N_PROC = 1 // Single-threaded
    start := time.Now()
    dataframe.Column("charge").Add(*dataframe.Column("walltime"))
    fmt.Println("Single Proc Add Duration:", time.Since(start).Seconds(), "seconds.")

    joey.N_PROC = 8 // Multi-threaded
    start = time.Now()
    dataframe.Column("charge").Add(*dataframe.Column("walltime"))
    fmt.Println("Multi Proc Add Duration:", time.Since(start).Seconds(), "seconds.")

    // Benchmark subtracting columns
    joey.N_PROC = 1 // Single-threaded
    start = time.Now()
    dataframe.Column("charge").Subtract(*dataframe.Column("walltime"))
    fmt.Println("Single Proc Subtract Duration:", time.Since(start).Seconds(), "seconds.")

    joey.N_PROC = 8 // Multi-threaded
    start = time.Now()
    dataframe.Column("charge").Subtract(*dataframe.Column("walltime"))
    fmt.Println("Multi Proc Subtract Duration:", time.Since(start).Seconds(), "seconds.")

    // Benchmark multiplying columns
    joey.N_PROC = 1 // Single-threaded
    start = time.Now()
    dataframe.Column("charge").Multiply(*dataframe.Column("walltime"))
    fmt.Println("Single Proc Multiply Duration:", time.Since(start).Seconds(), "seconds.")

    joey.N_PROC = 8 // Multi-threaded
    start = time.Now()
    dataframe.Column("charge").Multiply(*dataframe.Column("walltime"))
    fmt.Println("Multi Proc Multiply Duration:", time.Since(start).Seconds(), "seconds.")
}
```

---

## Additional Python Code Details

### File: `benchmarking_columns.py`

This Python script provides a simple benchmarking example for performing arithmetic operations on dataframe columns.

#### How It Works
1. **CSV Loading**: A dataframe is loaded from a file named `benchmarking_data.csv`, located in the same directory as the script.
2. **Adding Columns**: Benchmarks the time to add the `charge` and `walltime` columns.
3. **Subtracting Columns**: Benchmarks the time to subtract the `walltime` column from the `charge` column.
4. **Multiplying Columns**: Benchmarks the time to multiply the `charge` and `walltime` columns.
5. **Performance Measurement**: The execution time for each operation is measured using the `time` module and printed in milliseconds.

#### Python Code Example: `benchmarking_columns.py`
```python
import pandas as pd
import time, os

if __name__ == "__main__":
    dir = os.path.dirname(os.path.abspath(__file__))
    dataframe = pd.read_csv(os.path.join(dir, "benchmarking_data.csv"))

    print("----- Benchmarking Adding Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] + dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")

    print("----- Benchmarking Subtracting Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] - dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")

    print("----- Benchmarking Multiplying Columns -----")
    start_time = time.time()
    dataframe["charge"] = dataframe["charge"] * dataframe["walltime"]
    print(f"{(time.time() - start_time) * 1000} ms")
```

---

## Requirements

- **Go**: Ensure you have Go installed on your system.
- **Joey Library**: This example depends on the Joey library, which must be installed and available in your Go workspace.
- **Python**: Ensure you have Python installed on your system.
- **Pandas Library**: Install the Pandas library using `pip install pandas`.
- **CSV File**: Place the `benchmarking_data.csv` file in the same directory as both scripts.

---

## Running the Examples

### Go Example
1. Place your CSV file in the same directory as the Go program.
2. Run the Go example using:

   ```sh
   go run examples/benchmarking_multiproc.go
   ```

### Python Example
1. Place your CSV file in the same directory as the Python script.
2. Run the Python script using:

   ```sh
   python benchmarking_columns.py
   ```

---

## Expected Output

For the Go example:

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

For the Python example:

```
----- Benchmarking Adding Columns -----
12.345 ms
----- Benchmarking Subtracting Columns -----
23.456 ms
----- Benchmarking Multiplying Columns -----
34.567 ms
```

---

## Notes
- Multi-threaded performance in the Go example may vary depending on your hardware and the size of the dataset.
- The Python script provides a single-threaded implementation for simplicity, focusing on clarity and ease of understanding.

Enjoy benchmarking with Go, Python, and the Joey library!