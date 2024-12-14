
# Joey Framework -- Under Development

Before reading this, keep in mind that this is a library which is under development. We don't recommend you to use this, until it is available on its first production version.

Joey is a powerful framework for data manipulation and analysis developed in Go, designed to be efficient, flexible, and intuitive. It provides a robust interface for handling large volumes of data, making information processing and analysis simpler and more organized.

## Key Features

- **Intuitive Data Structure:** Easily manipulate data using `Dataframe`, which organizes information into named columns.
- **Columnar Operations:** Perform arithmetic operations like between columns.
- **CSV Compatibility:** Easily load and manipulate data from CSV files.
- **Data Visualization:** Print and view data in user-friendly and informative formats.
- **Extensibility:** Create new cell types to meet your specific needs.

## Usage Example

### Creating a Dataframe from a CSV

```go
import "joey"

func main() {
    df, err := joey.NewFromCsv("data.csv")
    if err != nil {
        panic(err)
    }

    // Display the data
    df.Show()

    // Sum the values of a column
    sum := df.Column("example_column").Sum()
    fmt.Printf("The sum is: %f\n", sum)
}
```

### Basic Operations

#### Removing a Column

```go
df, err = df.RemoveCol("column_to_remove")
if err != nil {
    fmt.Println(err)
}
```

#### Converting Data Types

```go
df, err = df.Convert("example_column", "float64")
if err != nil {
    fmt.Println(err)
}
```

#### Displaying Column Types

```go
df.ShowTypes()
```

## Installation

Make sure you have Go installed on your machine and run the following command:

```sh
go get github.com/LeonardoGrigoletto/joey
```

## Contributing

Feel free to open issues and pull requests on the GitHub repository. We are always open to suggestions and improvements.

## License

This project is licensed under the terms of the MIT license. See the LICENSE file for more details.

