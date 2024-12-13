package cells

import (
	"testing"
)

func TestAddFloat32Cell(t *testing.T) {
	oneCell := Float32Cell{Data: 50}
	otherCell := Float32Cell{Data: 30}
	oneCell.Add(&otherCell)

	if oneCell.Data != 80 {
		t.Fatalf("oneCell.Add(otherCell) does not work.")
	}
}

func TestAddMismatchedTypeFloat32Cell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("oneCell.Add(otherCell) should be in panic.")
		}
	}()
	oneCell := Float32Cell{Data: 50}
	otherCell := Int8Cell{Data: 30}
	oneCell.Add(&otherCell)
}
