package joey

import (
	"errors"
	"strconv"
	"strings"
)

type Float32Cell struct {
	data float32
}

func (f Float32Cell) GetFormattedData() string {
	return strconv.FormatFloat(float64(f.data), 'f', 2, 64)
}

func (f Float32Cell) Length() int {
	return len(strconv.FormatFloat(float64(f.data), 'f', 2, 64))
}

func (f Float32Cell) Convert(to string) (Cell, error) {
	if strings.EqualFold(to, "str") {
		convertedData := strconv.FormatFloat(float64(f.data), 'f', 2, 64)
		return StrCell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData := int32(f.data)
		return Int32Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData := int64(f.data)
		return Int64Cell{data: convertedData}, nil
	}
	if strings.EqualFold(to, "float32") {
		return f, nil
	}
	if strings.EqualFold(to, "float64") {
		convertedData := float64(f.data)
		return Float64Cell{data: convertedData}, nil
	}
	return nil, errors.New("Cannot convert to type: " + to)
}

func (f Float32Cell) GetRawData() any {
	return f.data
}

func (f Float32Cell) GetNumber() float64 {
	return float64(f.data)
}
