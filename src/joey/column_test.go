package joey

import (
	"reflect"
	"testing"
)

func TestColumnAdd(t *testing.T) {
	oneColumn := Column{
		Data: []Cell{
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []Cell{
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
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

func TestColumnAddMismatchedLengths(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneColumn.Add(otherColumn) should be in panic.")
		}
	}()
	oneColumn := Column{
		Data: []Cell{
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
		},
	}

	otherColumn := Column{
		Data: []Cell{
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
			&IntCell{Data: 30},
		},
	}
	oneColumn.Add(otherColumn)
}

func TestColumnConversion(t *testing.T) {
	oneColumn := Column{
		Data: []Cell{
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
			&IntCell{Data: 50},
		},
	}
	oneColumn.Convert("str")
	if oneColumn.GetType() != reflect.TypeOf(StrCell{}) {
		t.Fatalf("Column conversion to StrCell did not work properly.")
	}
	if oneColumn.GetNativeType() != reflect.TypeOf("StrCell{}") {
		t.Fatalf("Column conversion to string did not work properly.")
	}
}
