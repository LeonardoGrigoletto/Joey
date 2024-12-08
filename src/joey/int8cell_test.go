package joey

import (
	"testing"
)

func TestAddInt8Cell(t *testing.T) {
	oneCell := Int8Cell{Data: 1}
	otherCell := Int8Cell{Data: 2}
	oneCell.Add(&otherCell)

	if oneCell.Data != 3 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeInt8Cell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := Int8Cell{Data: 50}
	otherCell := Int16Cell{Data: 30}
	oneCell.Add(&otherCell)
}
