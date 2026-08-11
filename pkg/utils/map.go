package utils

import (
	"strings"
)

// NestedMap is a map that supports nested keys.
// It is expected that the nested maps are of type map[string]interface{}
type NestedMap map[string]any

func (m NestedMap) Get(key string) (any, bool) {
	fields := strings.Split(key, ".")

	current := m

	for _, f := range fields[:len(fields)-1] {
		v, found := current[f]
		if !found {
			return nil, false
		}

		current, _ = v.(map[string]any)
		if current == nil {
			return nil, false
		}
	}

	ret, found := current[fields[len(fields)-1]]
	return ret, found
}

func (m NestedMap) Set(key string, value any) {
	fields := strings.Split(key, ".")

	current := m

	for _, f := range fields[:len(fields)-1] {
		v, ok := current[f].(map[string]any)
		if !ok {
			v = make(map[string]any)
			current[f] = v
		}

		current = v
	}

	current[fields[len(fields)-1]] = value
}

func (m NestedMap) Delete(key string) {
	fields := strings.Split(key, ".")

	current := m

	for _, f := range fields[:len(fields)-1] {
		v, ok := current[f].(map[string]any)
		if !ok {
			return
		}

		current = v
	}

	delete(current, fields[len(fields)-1])
}

// MergeMaps merges src into dest. If a key exists in both maps, the value from src is used.
func MergeMaps(dest map[string]any, src map[string]any) {
	for k, v := range src {
		if _, ok := dest[k]; ok {
			if srcMap, ok := v.(map[string]any); ok {
				if destMap, ok := dest[k].(map[string]any); ok {
					MergeMaps(destMap, srcMap)
					continue
				}
			}
		}

		dest[k] = v
	}
}
