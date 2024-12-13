package cells

import (
	"testing"
)

func TestAddFloat64Cell(t *testing.T) {
	oneCell := Float64Cell{Data: 50}
	otherCell := Float64Cell{Data: 30}
	oneCell.Add(&otherCell)

	if oneCell.Data != 80 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeFloat64Cell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := Float64Cell{Data: 50}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}
