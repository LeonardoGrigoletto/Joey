package cells

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Float64Cell struct {
	Data float64
}

func (f *Float64Cell) Add(cell Cell) {
	data := cell.GetRawData()
	value, ok := data.(float64)
	if !ok {
		panic("It is not possible to sum columns of different types")
	}
	f.Data += value
}

func (f Float64Cell) GetType() interface{} {
	return reflect.TypeOf(Float64Cell{})
}

func (f Float64Cell) GetNativeType() interface{} {
	return reflect.TypeOf(f.Data)
}

func (f Float64Cell) GetFormattedData() string {
	return strconv.FormatFloat(f.Data, 'f', 2, 64)
}

func (f Float64Cell) Length() int {
	return len(strconv.FormatFloat(f.Data, 'f', 2, 64))
}

func (f Float64Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(f.Data, 'f', 2, 64)
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
	if strings.EqualFold(to, "int16") {
		convertedData := int16(f.Data)
		return &Int16Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int8") {
		convertedData := int8(f.Data)
		return &Int8Cell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		convertedData := int(f.Data)
		return &IntCell{Data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		return &f, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData := float32(f.Data)
		return &Float32Cell{Data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (f Float64Cell) GetRawData() any {
	return f.Data
}

func (f Float64Cell) GetNumber() float64 {
	return f.Data
}
