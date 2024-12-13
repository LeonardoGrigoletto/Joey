package cells

import (
	"reflect"
	"testing"
)

func TestAddIntCell(t *testing.T) {
	oneCell := IntCell{Data: 50}
	otherCell := IntCell{Data: 30}
	oneCell.Add(&otherCell)

	if oneCell.Data != 80 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeIntCell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := IntCell{Data: 50}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}

func TestConvertIntCell(t *testing.T) {
	oneCell := IntCell{Data: 50}
	convertedCell, err := oneCell.Convert("str")
	if err != nil {
		t.Fatalf("%s", err)
	}
	if convertedCell.GetNativeType() != reflect.TypeOf("") {
		t.Fatalf("Convertion from int to str did not work.")
	}
	if convertedCell.GetType() != reflect.TypeOf(StrCell{}) {
		t.Fatalf("Convertion from IntCell to StrCell did not work.")
	}

}
