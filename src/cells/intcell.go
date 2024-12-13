package cells

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type IntCell struct {
	Data int
}

func (i *IntCell) Add(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(int)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	i.Data += value
}

func (i *IntCell) Multiply(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(int)
	if !ok {
		panic("It is not possible to multiply columns of different types")
	}
	i.Data *= value
}

func (i *IntCell) Subtract(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(int)
	if !ok {
		panic("It is not possible to subtract columns of different types")
	}
	i.Data -= value
}

// GetType implements Cell.
func (i IntCell) GetType() interface{} {
	return reflect.TypeOf(IntCell{})
}

func (i IntCell) GetNativeType() interface{} {
	return reflect.TypeOf(i.Data)
}

func (i IntCell) GetFormattedData() string {
	return strconv.FormatInt(int64(i.Data), 10)
}

func (i IntCell) Length() int {
	return len(strconv.Itoa(int(i.Data)))
}

func (i IntCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatInt(int64(i.Data), 10)
		return &StrCell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		return &i, nil
	}
	if strings.EqualFold(to, "int8") {
		convertedData := int8(i.Data)
		return &Int8Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int16") {
		convertedData := int16(i.Data)
		return &Int16Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData := int32(i.Data)
		return &Int32Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData := int64(i.Data)
		return &Int64Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData := float32(i.Data)
		return &Float32Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData := float64(i.Data)
		return &Float64Cell{Data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (i IntCell) GetRawData() any {
	return i.Data
}

func (i IntCell) GetNumber() float64 {
	return float64(i.Data)
}
