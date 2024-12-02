package joey

import (
	"errors"
	"strconv"
	"strings"
)

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
		convertedData := int(i.data)
		return IntCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int8") {
		convertedData := int8(i.data)
		return Int8Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int16") {
		convertedData := int16(i.data)
		return Int16Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData := int32(i.data)
		return Int32Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int64") {
		return i, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData := float32(i.data)
		return Float32Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		return Float64Cell{data: float64(i.data)}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (i Int64Cell) GetRawData() any {
	return i.data
}

func (i Int64Cell) GetNumber() float64 {
	return float64(i.data)
}
