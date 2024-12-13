package cells

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Int8Cell struct {
	Data int8
}

func (i *Int8Cell) Add(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(int8)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	i.Data += value
}

func (i Int8Cell) GetType() interface{} {
	return reflect.TypeOf(Int8Cell{})
}

func (i Int8Cell) GetNativeType() interface{} {
	return reflect.TypeOf(i.Data)
}

func (i Int8Cell) GetFormattedData() string {
	return strconv.FormatInt(int64(i.Data), 10)
}

func (i Int8Cell) Length() int {
	return len(strconv.Itoa(int(i.Data)))
}

func (i Int8Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatInt(int64(i.Data), 10)
		return &StrCell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int8") {
		return &i, nil
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
		return &Float64Cell{Data: float64(i.Data)}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (i Int8Cell) GetRawData() any {
	return i.Data
}

func (i Int8Cell) GetNumber() float64 {
	return float64(i.Data)
}
