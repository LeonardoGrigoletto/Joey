package cells

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Float32Cell struct {
	Data float32
}

func (f *Float32Cell) Add(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(float32)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	f.Data += value
}

func (f *Float32Cell) Multiply(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(float32)
	if !ok {
		panic("It is not possible to multiply columns of different types")
	}
	f.Data *= value
}

func (f *Float32Cell) Subtract(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(float32)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	f.Data -= value
}

func (f Float32Cell) GetType() interface{} {
	return reflect.TypeOf(Float32Cell{})
}

func (f Float32Cell) GetNativeType() interface{} {
	return reflect.TypeOf(f.Data)
}

func (f Float32Cell) GetFormattedData() string {
	return strconv.FormatFloat(float64(f.Data), 'f', 2, 64)
}

func (f Float32Cell) Length() int {
	return len(strconv.FormatFloat(float64(f.Data), 'f', 2, 64))
}

func (f Float32Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(float64(f.Data), 'f', 2, 64)
		return &StrCell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData := int32(f.Data)
		return &Int32Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData := int64(f.Data)
		return &Int64Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "float32") {
		return &f, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData := float64(f.Data)
		return &Float64Cell{Data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (f Float32Cell) GetRawData() any {
	return f.Data
}

func (f Float32Cell) GetNumber() float64 {
	return float64(f.Data)
}
