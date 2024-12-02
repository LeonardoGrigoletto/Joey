package joey

import (
	"errors"
	"reflect"
)

type Row struct {
	data []Cell
}

type Column struct {
	data       []*Cell
	validTypes map[reflect.Type]bool
}

func (c *Column) New(size int) Column {
	validTypes := map[reflect.Type]bool{
		reflect.TypeOf("string"):   true, // string
		reflect.TypeOf(float64(2)): true, // float64
		reflect.TypeOf(int64(2)):   true, // int64
	}
	column := Column{data: make([]*Cell, size), validTypes: validTypes}
	return column
}

func (c *Column) Sum() float64 {
	sum := float64(0)
	for _, cellPtr := range c.data {
		cell := *cellPtr
		cellNumberValue := cell.GetNumber()
		sum += cellNumberValue
	}
	return sum
}

func (c *Column) checkType(item interface{}) error {
	cell := *c.data[0]
	typeOfItem := reflect.TypeOf(item)
	typeOfCells := reflect.TypeOf(cell.GetRawData())
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
	for i, cellPtr := range c.data {
		cell := *cellPtr
		rawData := cell.GetRawData()
		if rawData == item {
			return i
		}
	}
	return -1
}
