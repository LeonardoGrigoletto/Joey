package joey

import (
	"reflect"
	"testing"
)

func TestAddStrCell(t *testing.T) {
	oneCell := StrCell{Data: "50"}
	otherCell := StrCell{Data: "30"}
	oneCell.Add(&otherCell)

	if oneCell.Data != "5030" {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeStrCell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := StrCell{Data: "50"}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}

func TestConvertStrCell(t *testing.T) {
	oneCell := StrCell{Data: "50"}
	convertedCell, err := oneCell.Convert("int")
	if err != nil {
		t.Fatalf("%s", err)
	}
	if convertedCell.GetNativeType() != reflect.TypeOf(int(50)) {
		t.Fatalf("Convertion from int to str did not work.")
	}
	if convertedCell.GetType() != reflect.TypeOf(IntCell{}) {
		t.Fatalf("Convertion from StrCell to IntCell did not work.")
	}
}
