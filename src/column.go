package joey

import (
	"errors"
	"joey/cells"
	"reflect"
	"sync"
)

type Column struct {
	name string
	Data []cells.Cell
}

func Repeat(size int, value float64, name string) *Column {
	column := &Column{Data: make([]cells.Cell, size), name: name}
	for i := range column.Data {
		column.Data[i] = &cells.Float64Cell{Data: value}
	}
	return column
}

func (c *Column) New(size int, name string) Column {
	column := Column{Data: make([]cells.Cell, size), name: name}
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

func (c *Column) validateToPerformArithmeticOperation(otherColumn Column) {
	if !c.hasEqualTypes(otherColumn) {
		panic("It is not possible to sum columns of different types")
	}
	if !c.hasEqualLengths(otherColumn) {
		panic("Column Lengths mismatched.")
	}
}

func (c *Column) calculateChunckSize() int {
	return (len(c.Data) + N_PROC - 1) / N_PROC
}

func (c *Column) applyMultiProc(otherColumn Column, operation func(c1, c2 cells.Cell)) Column {
	var wg sync.WaitGroup
	chunkSize := c.calculateChunckSize()

	// Creating processes
	for i := 0; i < N_PROC; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(c.Data) {
			end = len(c.Data)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				operation(c.Data[j], otherColumn.Data[j])
			}
		}(start, end)
	}

	// Wait for the workers have finished
	wg.Wait()

	return *c
}

func (c *Column) Add(otherColumn Column) Column {
	c.validateToPerformArithmeticOperation(otherColumn)
	if N_PROC == 1 {
		for i, cell := range c.Data {
			cell.Add(otherColumn.Data[i])
		}
		return *c
	}
	return c.applyMultiProc(otherColumn, func(c1, c2 cells.Cell) {
		c1.Add(c2)
	})
}

func (c *Column) Subtract(otherColumn Column) Column {
	c.validateToPerformArithmeticOperation(otherColumn)
	if N_PROC == 1 {
		for i, cell := range c.Data {
			cell.Subtract(otherColumn.Data[i])
		}
		return *c
	}
	return c.applyMultiProc(otherColumn, func(c1, c2 cells.Cell) {
		c1.Subtract(c2)
	})
}

func (c *Column) Multiply(otherColumn Column) Column {
	c.validateToPerformArithmeticOperation(otherColumn)
	if N_PROC == 1 {
		for i, cell := range c.Data {
			cell.Multiply(otherColumn.Data[i])
		}
		return *c
	}
	return c.applyMultiProc(otherColumn, func(c1, c2 cells.Cell) {
		c1.Multiply(c2)
	})
}

func (c *Column) Convert(to string) Column {
	for i, cell := range c.Data {
		convertedCell, err := cell.Convert(to)
		if err != nil {
			panic(err)
		}
		c.Data[i] = convertedCell
	}
	return *c
}
