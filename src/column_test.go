package joey

import (
	"joey/cells"
	"reflect"
	"testing"
)

func TestColumnAdd(t *testing.T) {
	oneColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
		},
	}

	resultColumn := oneColumn.Add(otherColumn)
	for _, cell := range resultColumn.Data {
		if cell.GetRawData() != 80 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
	for _, cell := range oneColumn.Data {
		if cell.GetRawData() != 80 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
}

func TestColumnMultiply(t *testing.T) {
	oneColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
		},
	}

	resultColumn := oneColumn.Multiply(otherColumn)
	for _, cell := range resultColumn.Data {
		if cell.GetRawData() != 1500 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
	for _, cell := range oneColumn.Data {
		if cell.GetRawData() != 1500 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
}

func TestColumnSubtract(t *testing.T) {
	oneColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
		},
	}

	resultColumn := oneColumn.Subtract(otherColumn)
	for _, cell := range resultColumn.Data {
		if cell.GetRawData() != 20 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
	for _, cell := range oneColumn.Data {
		if cell.GetRawData() != 20 {
			t.Fatalf("Column Add did not work properly.")
		}
	}
}

func TestColumnAddMismatchedLengths(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneColumn.Add(otherColumn) should be in panic.")
		}
	}()
	oneColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
			&cells.IntCell{Data: 30},
		},
	}
	oneColumn.Add(otherColumn)
}

func TestColumnConversion(t *testing.T) {
	oneColumn := Column{
		Data: []cells.Cell{
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
			&cells.IntCell{Data: 50},
		},
	}
	oneColumn.Convert("str")
	if oneColumn.GetType() != reflect.TypeOf(cells.StrCell{}) {
		t.Fatalf("Column conversion to StrCell did not work properly.")
	}
	if oneColumn.GetNativeType() != reflect.TypeOf("StrCell{}") {
		t.Fatalf("Column conversion to string did not work properly.")
	}
}
