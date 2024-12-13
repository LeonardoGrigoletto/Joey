package cells

import (
	"testing"
)

func TestAddInt32Cell(t *testing.T) {
	oneCell := Int32Cell{Data: 50}
	otherCell := Int32Cell{Data: 30}
	oneCell.Add(&otherCell)

	if oneCell.Data != 80 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeInt32Cell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := Int32Cell{Data: 50}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}
