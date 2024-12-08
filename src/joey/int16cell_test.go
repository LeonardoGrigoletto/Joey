package joey

import (
	"testing"
)

func TestAddInt16Cell(t *testing.T) {
	oneCell := Int16Cell{Data: 50}
	otherCell := Int16Cell{Data: 30}
	oneCell.Add(&otherCell)

	if oneCell.Data != 80 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeInt16Cell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := Int16Cell{Data: 50}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}
