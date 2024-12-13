package joey

import (
	"joey/cells"
	"reflect"
	"testing"
)

func TestColumnShouldGetColumn(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"
	column := dataframe.Column(columnNameToTest)
	if column.name != columnNameToTest {
		t.Fatalf("Column names does not match. %s != %s", column.name, columnNameToTest)
	}
}

func TestColumnNotExist(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals2"
	column := dataframe.Column(columnNameToTest)
	if len(column.Data) != 0 {
		t.Fatalf("Column should be empty")
	}
}

func TestConvertShouldConvertColumnToString(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to String Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "str")
	cellType := reflect.TypeOf(&cells.StrCell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToInt8(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Int8 Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "int8")
	cellType := reflect.TypeOf(&cells.Int8Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToInt16(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Int16Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "int16")
	cellType := reflect.TypeOf(&cells.Int16Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToInt32(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Int32Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "int32")
	cellType := reflect.TypeOf(&cells.Int32Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToInt64(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Int64Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "int64")
	cellType := reflect.TypeOf(&cells.Int64Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToInt(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to IntCell
	dataframe, err = dataframe.Convert(columnNameToTest, "int")
	cellType := reflect.TypeOf(&cells.IntCell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToFloat32(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Float32Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "float32")
	cellType := reflect.TypeOf(&cells.Float32Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestConvertShouldConvertColumnToFloat64(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}
	columnNameToTest := "number_of_visuals"

	// Testing conversion to Float64Cell
	dataframe, err = dataframe.Convert(columnNameToTest, "float64")
	cellType := reflect.TypeOf(&cells.Float64Cell{})
	convertedColumnType := reflect.TypeOf(dataframe.Column(columnNameToTest).Data[0])
	isSameType := convertedColumnType == cellType
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !isSameType {
		t.Fatalf("Types does not match. %s != %s", convertedColumnType, cellType)
	}
}

func TestCreateSeparator(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}

	separator := dataframe.createSeparator([]int{5, 10, 3, 4, 15})
	correctSeparator := "+-------+------------+-----+------+-----------------+"
	if separator != correctSeparator {
		t.Fatalf("Separators does not match.\ncorrectSeparator = %s\n       Separator = %s", correctSeparator, separator)
	}
}

func TestRemoveCol(t *testing.T) {
	filePath := setup()

	dataframe, err := NewFromCsv(filePath)
	if err != nil {
		t.Fatalf("%s", err)
	}

	colToRemove := "charge"
	dataframe, err = dataframe.RemoveCol(colToRemove)
	if err != nil {
		t.Fatalf("%s", err)
	}

	column := dataframe.Column("charge")
	if len(column.Data) != 0 {
		t.Fatalf("Column was not removed.")
	}

}
