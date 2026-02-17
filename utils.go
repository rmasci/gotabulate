package gotabulate

import (
	"fmt"
	"strconv"
)

// createFromString creates normalized rows from a string slice.
func createFromString(data [][]string) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))

	for i, row := range data {
		rows[i] = &TabulateRow{Elements: row}
	}
	return rows
}

// createFromMixed creates normalized rows from mixed interface{} data.
func createFromMixed(data [][]interface{}, format byte) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, element := range data {
		normalized := make([]string, len(element))
		for j, el := range element {
			normalized[j] = formatValue(el, format)
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// formatValue converts a value to a string based on its type.
func formatValue(el interface{}, format byte) string {
	switch v := el.(type) {
	case int32:
		quoted := strconv.QuoteRuneToASCII(v)
		return quoted[1 : len(quoted)-1]
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, format, -1, 64)
	case uint64:
		return strconv.FormatUint(v, 10)
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// createFromInt creates normalized rows from int data.
func createFromInt(data [][]int) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, row := range data {
		normalized := make([]string, len(row))
		for j, val := range row {
			normalized[j] = strconv.Itoa(val)
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// createFromFloat64 creates normalized rows from float64 data.
func createFromFloat64(data [][]float64, format byte) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, row := range data {
		normalized := make([]string, len(row))
		for j, val := range row {
			normalized[j] = strconv.FormatFloat(val, format, -1, 64)
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// createFromInt32 creates normalized rows from int32 data.
func createFromInt32(data [][]int32) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, row := range data {
		normalized := make([]string, len(row))
		for j, val := range row {
			quoted := strconv.QuoteRuneToASCII(val)
			normalized[j] = quoted[1 : len(quoted)-1]
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// createFromInt64 creates normalized rows from int64 data.
func createFromInt64(data [][]int64) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, row := range data {
		normalized := make([]string, len(row))
		for j, val := range row {
			normalized[j] = strconv.FormatInt(val, 10)
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// createFromBool creates normalized rows from bool data.
func createFromBool(data [][]bool) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for i, row := range data {
		normalized := make([]string, len(row))
		for j, val := range row {
			normalized[j] = strconv.FormatBool(val)
		}
		rows[i] = &TabulateRow{Elements: normalized}
	}
	return rows
}

// createFromMapMixed creates normalized rows from a map of mixed elements.
// Map keys will be used as headers.
func createFromMapMixed(data map[string][]interface{}, format byte) ([]string, []*TabulateRow) {
	var headers []string
	var dataSlice [][]interface{}
	for key, value := range data {
		headers = append(headers, key)
		dataSlice = append(dataSlice, value)
	}
	return headers, createFromMixed(dataSlice, format)
}

// createFromMapString creates normalized rows from a map of strings.
// Map keys will be used as headers.
func createFromMapString(data map[string][]string) ([]string, []*TabulateRow) {
	var headers []string
	var dataSlice [][]string
	for key, value := range data {
		headers = append(headers, key)
		dataSlice = append(dataSlice, value)
	}
	return headers, createFromString(dataSlice)
}

// inSlice checks if an element exists in a string slice.
func inSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
