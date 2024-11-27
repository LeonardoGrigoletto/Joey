package joey

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type Dataframe struct {
	headers []string
	rows    []Row
	columns []Column
}

func (d *Dataframe) Column(name string) (Column, error) {
	columnIndex, err := d.getColumnIndex(name)
	if err != nil {
		return Column{}, err
	}
	return d.columns[columnIndex], nil
}

func (d *Dataframe) Convert(columnName string, to string) (Dataframe, error) {
	columnIndex, err := d.getColumnIndex(columnName)
	if err != nil {
		return Dataframe{}, err
	}

	for i, cellPtr := range d.columns[columnIndex].data {
		if cellPtr == nil {
			return Dataframe{}, errors.New("nil cell pointer found")
		}

		cell := *cellPtr
		convertedCell, err := cell.Convert(to)
		if err != nil {
			return Dataframe{}, err
		}

		*d.columns[columnIndex].data[i] = convertedCell
	}
	return Dataframe{
		headers: d.headers,
		rows:    d.rows,
		columns: d.columns,
	}, nil
}

func (d *Dataframe) createSeparator(colWidths []int) string {
	separator := "+"
	for _, width := range colWidths {
		separator += strings.Repeat("-", width+2) + "+"
	}
	return separator
}

func (d *Dataframe) printHeader(colWidths []int) {
	for i, header := range d.headers {
		fmt.Printf("| %-*s ", colWidths[i], header)
	}
}

func (d *Dataframe) printRecords(colWidths []int, numberOfRecords int) {
	for i := 0; i < numberOfRecords; i++ {
		for j, cell := range d.rows[i].data {
			fmt.Printf("| %-*s ", colWidths[j], cell.GetFormattedData())
		}
		fmt.Println("|")
	}
}

func (d *Dataframe) calculateNumberOfRecordsToPrint(size ...int) int {
	numberOfRecordsToPrint := len(d.rows)
	if len(size) > 0 {
		numberOfRecordsToPrint = size[0]
	}
	if numberOfRecordsToPrint > len(d.rows) {
		numberOfRecordsToPrint = len(d.rows)
	}
	return numberOfRecordsToPrint
}

func (d *Dataframe) calculateColWidthsToPrint(size ...int) []int {
	colWidths := make([]int, len(d.headers))
	for i, header := range d.headers {
		colWidths[i] = len(header)
	}
	for k := 0; k < size[0]; k++ {
		for i, cell := range d.rows[k].data {
			if cell.Length() > colWidths[i] {
				colWidths[i] = cell.Length()
			}
		}
	}
	return colWidths
}

func (d *Dataframe) Show(size ...int) {
	numberOfRecordsToPrint := d.calculateNumberOfRecordsToPrint(size...)
	colWidths := d.calculateColWidthsToPrint(size...)
	separator := d.createSeparator(colWidths)

	// HEADER
	fmt.Println(separator)
	d.printHeader(colWidths)
	fmt.Println("|")
	fmt.Println(separator)

	// DATA
	d.printRecords(colWidths, numberOfRecordsToPrint)
	fmt.Println(separator)

	// FOOTER
	fmt.Printf("Showing %d/%d records\n", numberOfRecordsToPrint, len(d.rows))
	fmt.Println("")
}

func (d *Dataframe) ShowTypes() {
	green := "\033[32m"
	reset := "\033[0m"
	bold := "\033[1m"
	fmt.Println("+--- Fields Types ---+")
	for i, cell := range d.rows[0].data {
		fmt.Printf(green+"%s"+reset+" --- "+bold+"%s\n"+reset, d.headers[i], reflect.TypeOf(cell))
	}
	fmt.Println("")
}

func (d *Dataframe) getColumnIndex(colName string) (int, error) {
	firstIndex := slices.Index(d.headers, colName)
	if firstIndex < 0 {
		return firstIndex, errors.New("The specified column name does not exist in the dataframe. Column: " + colName)
	}
	return firstIndex, nil
}

func (d *Dataframe) RemoveCol(colName string) (Dataframe, error) {
	firstIndex, err := d.getColumnIndex(colName)
	if err != nil {
		return Dataframe{}, err
	}
	d.headers = slices.Delete(d.headers, firstIndex, firstIndex+1)
	for index := range d.rows {
		d.rows[index].data = slices.Delete(d.rows[index].data, firstIndex, firstIndex+1)
	}
	return Dataframe{headers: d.headers, rows: d.rows, columns: d.columns}, nil
}

// func (d *Dataframe) Sum(columns ...string) Dataframe {
// 	if len(columns) <= 1 {
// 		panic("You must specify at least two columns.")
// 	}

// 	var columnIndexes []int = make([]int, len(columns))
// 	for i, column := range columns {
// 		columnIndex := slices.Index(d.headers, column)
// 		if columnIndex < 0 {
// 			panic("The given column: " + column + ", does not exist on the dataframe.")
// 		}
// 		columnIndexes[i] = columnIndex
// 	}

// 	d.headers = append(d.headers, "sum")

// 	// for i, row := range d.rows {
// 	// 	sum_result := 0
// 	// 	for _, index := range columnIndexes {
// 	// 		sum_result += d.rows[i].data

// 	// 	}
// 	// 	row.data = append(row.data, )
// 	// }
// 	return Dataframe{}

// }
