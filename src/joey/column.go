package joey

import (
	"errors"
	"reflect"
)

type Column struct {
	name string
	Data []Cell
}

func (c *Column) New(size int, name string) Column {
	column := Column{Data: make([]Cell, size), name: name}
	return column
}

func (c *Column) Sum() float64 {
	sum := float64(0)
	for _, cell := range c.Data {
		cellNumberValue := cell.GetNumber()
		sum += cellNumberValue
	}
	return sum
}

func (c *Column) GetNativeType() interface{} {
	return c.Data[0].GetNativeType()
}

func (c *Column) GetType() interface{} {
	return c.Data[0].GetType()
}

func (c *Column) checkType(item interface{}) error {
	typeOfItem := reflect.TypeOf(item)
	typeOfCells := c.GetNativeType()
	if typeOfItem != typeOfCells {
		return errors.New("invalid specified type. Must be string, floatX or intX")
	}
	return nil
}

func (c *Column) FindFirst(item interface{}) int {
	err := c.checkType(item)
	if err != nil {
		return -1
	}
	for i, cell := range c.Data {
		rawData := cell.GetRawData()
		if rawData == item {
			return i
		}
	}
	return -1
}

func (c *Column) hasEqualTypes(otherColumn Column) bool {
	thisType := c.GetType()
	otherColumnType := otherColumn.GetType()
	return thisType == otherColumnType
}

func (c *Column) hasEqualLengths(otherColumn Column) bool {
	return len(c.Data) == len(otherColumn.Data)
}

func (c *Column) Add(otherColumn Column) Column {
	if !c.hasEqualTypes(otherColumn) {
		panic("It is not possible to sum columns of different types")
	}
	if !c.hasEqualLengths(otherColumn) {
		panic("Column Lengths mismatched.")
	}
	for i, cell := range c.Data {
		cell.Add(otherColumn.Data[i])
	}
	return *c
}

func (c *Column) Convert(to string) error {
	for i, cell := range c.Data {
		convertedCell, err := cell.Convert(to)
		if err != nil {
			return err
		}
		c.Data[i] = convertedCell
	}
	return nil
}
