package joey

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Cell interface {
	GetFormattedData() string
	GetRawData() interface{}
	Length() int
	Convert(to string) (Cell, error)
}

type Float64Cell struct {
	data float64
}

func (f Float64Cell) GetFormattedData() string {
	return strconv.FormatFloat(f.data, 'f', 2, 64)
}

func (f Float64Cell) Length() int {
	return len(strconv.FormatFloat(f.data, 'f', 2, 64))
}

func (f Float64Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(f.data, 'f', 2, 64)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData := int64(f.data)
		return Int64Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		return f, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (f Float64Cell) GetRawData() any {
	return f.data
}

type Int64Cell struct {
	data int64
}

func (i Int64Cell) GetFormattedData() string {
	return strconv.FormatInt(i.data, 10)
}

func (i Int64Cell) Length() int {
	return len(strconv.Itoa(int(i.data)))
}

func (i Int64Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatInt(int64(i.data), 10)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		return i, nil
	}
	if strings.EqualFold(to, "float64") {
		return Float64Cell{data: float64(i.data)}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (i Int64Cell) GetRawData() any {
	return i.data
}

type StrCell struct {
	data string
}

func (s StrCell) GetFormattedData() string {
	return s.data
}

func (s StrCell) Length() int {
	return len(s.data)
}

func (s StrCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		return s, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int64Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData, err := strconv.ParseFloat(s.data, 64)
		if err != nil {
			return nil, err
		}
		return Float64Cell{data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (s StrCell) GetRawData() any {
	return s.data
}

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
	if reflect.TypeOf(c.data[0]) == reflect.TypeOf("string") {
		return 0
	}
	for _, cellPtr := range c.data {
		cell := *cellPtr
		rawData := cell.GetRawData()
		value_int, ok_int := rawData.(int64)
		if ok_int {
			sum += float64(value_int)
		}
		value_float, ok_float := rawData.(float64)
		if ok_float {
			sum += float64(value_float)
		}
	}
	return sum
}

func (c *Column) checkType(item interface{}) bool {
	if c.validTypes[reflect.TypeOf(item)] {
		return true
	}
	panic("Invalid specified type. Must be string, floatX or intX")
}

func (c *Column) FindFirst(item interface{}) int {
	c.checkType(item)
	for i, cellPtr := range c.data {
		cell := *cellPtr
		rawData := cell.GetRawData()
		if rawData == item {
			return i
		}
	}
	return -1
}
