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
	columns []Column
}

func (d *Dataframe) Column(name string) Column {
	columnIndex := d.getColumnIndex(name)
	if columnIndex < 0 {
		return Column{}
	}
	return d.columns[columnIndex]
}

func (d *Dataframe) Convert(columnName string, to string) (Dataframe, error) {
	columnIndex := d.getColumnIndex(columnName)
	if columnIndex < 0 {
		return Dataframe{}, errors.New("Column Name %s does not exist." + columnName)
	}

	for i, cell := range d.columns[columnIndex].data {
		if cell == nil {
			return Dataframe{}, errors.New("nil cell pointer found")
		}

		convertedCell, err := cell.Convert(to)
		if err != nil {
			return Dataframe{}, err
		}

		d.columns[columnIndex].data[i] = convertedCell
	}
	return Dataframe{
		headers: d.headers,
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
		for j, column := range d.columns {
			cell := column.data[i]
			fmt.Printf("| %-*s ", colWidths[j], cell.GetFormattedData())
		}
		fmt.Println("|")
	}
}

func (d *Dataframe) calculateNumberOfRecordsToPrint(size ...int) int {
	numberOfRecordsToPrint := len(d.columns[0].data)
	if len(size) > 0 {
		numberOfRecordsToPrint = size[0]
	}
	if numberOfRecordsToPrint > len(d.columns[0].data) {
		numberOfRecordsToPrint = len(d.columns[0].data)
	}
	return numberOfRecordsToPrint
}

func (d *Dataframe) calculateColWidthsToPrint(numberOfRecordsToPrint int) []int {
	colWidths := make([]int, len(d.headers))
	for i, header := range d.headers {
		colWidths[i] = len(header)
	}

	for i, column := range d.columns {
		for k := 0; k < numberOfRecordsToPrint; k++ {
			if column.data[k].Length() > colWidths[i] {
				colWidths[i] = column.data[k].Length()
			}
		}
	}
	return colWidths
}

func (d *Dataframe) Show(size ...int) {
	numberOfRecordsToPrint := d.calculateNumberOfRecordsToPrint(size...)
	colWidths := d.calculateColWidthsToPrint(numberOfRecordsToPrint)
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
	fmt.Printf("Showing %d/%d records\n", numberOfRecordsToPrint, len(d.columns[0].data))
	fmt.Println("")
}

func (d *Dataframe) ShowTypes() {
	green := "\033[32m"
	reset := "\033[0m"
	bold := "\033[1m"
	fmt.Println("+--- Fields Types ---+")
	for i, column := range d.columns {
		cell := column.data[0]
		fmt.Printf(green+"%s"+reset+" --- "+bold+"%s\n"+reset, d.headers[i], reflect.TypeOf(cell))

	}
	fmt.Println("")
}

func (d *Dataframe) getColumnIndex(colName string) int {
	firstIndex := slices.Index(d.headers, colName)
	return firstIndex
}

func (d *Dataframe) RemoveCol(colName string) (Dataframe, error) {
	firstIndex := d.getColumnIndex(colName)
	if firstIndex < 0 {
		return Dataframe{}, errors.New("Column name %s does not exist." + colName)
	}
	d.headers = slices.Delete(d.headers, firstIndex, firstIndex+1)
	d.columns = slices.Delete(d.columns, firstIndex, firstIndex+1)
	return Dataframe{headers: d.headers, columns: d.columns}, nil
}
