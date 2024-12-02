package joey

import (
	"errors"
	"strconv"
	"strings"
)

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
		return IntCell{data: int(convertedData)}, nil
	}
	if strings.EqualFold(to, "int8") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int8Cell{data: int8(convertedData)}, nil
	}
	if strings.EqualFold(to, "int16") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int16Cell{data: int16(convertedData)}, nil
	}
	if strings.EqualFold(to, "int32") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int32Cell{data: int32(convertedData)}, nil
	}
	if strings.EqualFold(to, "int64") {
		convertedData, err := strconv.ParseInt(s.data, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int64Cell{data: int64(convertedData)}, nil
	}
	if strings.EqualFold(to, "float32") {
		convertedData, err := strconv.ParseFloat(s.data, 64)
		if err != nil {
			return nil, err
		}
		return Float32Cell{data: float32(convertedData)}, nil
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

func (s StrCell) GetNumber() float64 {
	return float64(0)
}
