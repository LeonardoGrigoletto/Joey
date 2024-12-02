package joey

import (
	"errors"
	"strconv"
	"strings"
)

type IntCell struct {
	data int
}

func (i IntCell) GetFormattedData() string {
	return strconv.FormatInt(int64(i.data), 10)
}

func (i IntCell) Length() int {
	return len(strconv.Itoa(int(i.data)))
}

func (i IntCell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatInt(int64(i.data), 10)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int") {
		return i, nil
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
		convertedData := int64(i.data)
		return Int64Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData := float32(i.data)
		return Float32Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData := float64(i.data)
		return Float64Cell{data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (i IntCell) GetRawData() any {
	return i.data
}

func (i IntCell) GetNumber() float64 {
	return float64(i.data)
}
