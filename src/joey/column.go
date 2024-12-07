package joey

import (
	"errors"
	"reflect"
)

type Column struct {
	name string
	data []Cell
}

func (c *Column) New(size int, name string) Column {
	column := Column{data: make([]Cell, size), name: name}
	return column
}

func (c *Column) Sum() float64 {
	sum := float64(0)
	for _, cell := range c.data {
		cellNumberValue := cell.GetNumber()
		sum += cellNumberValue
	}
	return sum
}

func (c *Column) checkType(item interface{}) error {
	typeOfItem := reflect.TypeOf(item)
	typeOfCells := reflect.TypeOf(c.data[0].GetRawData())
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
	for i, cell := range c.data {
		rawData := cell.GetRawData()
		if rawData == item {
			return i
		}
	}
	return -1
}
