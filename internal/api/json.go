package api

import (
	"encoding/json"
	"strings"
)

// jsonNumberToNumber converts a JSON number to either a float64 or int64.
func jsonNumberToNumber(n json.Number) any {
	if strings.Contains(string(n), ".") {
		f, _ := n.Float64()
		return f
	}
	ret, _ := n.Int64()
	return ret
}

// ConvertMapJSONNumbers converts all JSON numbers in a map to either float64 or int64.
func convertMapJSONNumbers(m map[string]any) (ret map[string]any) {
	if m == nil {
		return nil
	}

	ret = make(map[string]any)
	for k, v := range m {
		if n, ok := v.(json.Number); ok {
			ret[k] = jsonNumberToNumber(n)
		} else if mm, ok := v.(map[string]any); ok {
			ret[k] = convertMapJSONNumbers(mm)
		} else {
			ret[k] = v
		}
	}

	return ret
}
