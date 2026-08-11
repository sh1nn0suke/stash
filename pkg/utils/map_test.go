package utils

import (
	"reflect"
	"testing"
)

func TestNestedMapGet(t *testing.T) {
	m := NestedMap{
		"foo": map[string]any{
			"bar": map[string]any{
				"baz": "qux",
			},
		},
	}

	tests := []struct {
		name  string
		key   string
		want  any
		found bool
	}{
		{
			name:  "Get a value from a nested map",
			key:   "foo.bar.baz",
			want:  "qux",
			found: true,
		},
		{
			name:  "Get a value from a nested map with a missing key",
			key:   "foo.bar.quux",
			want:  nil,
			found: false,
		},
		{
			name:  "Get a value from a nested map with a missing key",
			key:   "foo.quux.baz",
			want:  nil,
			found: false,
		},
		{
			name:  "Get a value from a nested map with a missing key",
			key:   "quux.bar.baz",
			want:  nil,
			found: false,
		},
		{
			name:  "Get a value from a nested map with a missing key",
			key:   "foo.bar",
			want:  map[string]any{"baz": "qux"},
			found: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := m.Get(tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NestedMap.Get() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("NestedMap.Get() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestNestedMapSet(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		existing NestedMap
		want     NestedMap
	}{
		{
			name:     "Set a value in a nested map",
			key:      "foo.bar.baz",
			existing: NestedMap{},
			want: NestedMap{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "qux",
					},
				},
			},
		},
		{
			name: "Overwrite existing value",
			key:  "foo.bar",
			existing: NestedMap{
				"foo": map[string]any{
					"bar": "old",
				},
			},
			want: NestedMap{
				"foo": map[string]any{
					"bar": "qux",
				},
			},
		},
		{
			name: "Set a value overwriting a primitive with a nested map",
			key:  "foo.bar",
			existing: NestedMap{
				"foo": "bar",
			},
			want: NestedMap{
				"foo": map[string]any{
					"bar": "qux",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.existing.Set(tt.key, "qux")
			if !reflect.DeepEqual(tt.existing, tt.want) {
				t.Errorf("NestedMap.Set() got = %v, want %v", tt.existing, tt.want)
			}
		})
	}
}

func TestNestedMapDelete(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		existing NestedMap
		want     NestedMap
	}{
		{
			name: "Delete non existing value",
			key:  "foo.bar.baa",
			existing: NestedMap{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "qux",
					},
				},
			},
			want: NestedMap{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "qux",
					},
				},
			},
		},
		{
			name: "Delete existing value",
			key:  "foo.bar",
			existing: NestedMap{
				"foo": map[string]any{
					"bar": "old",
				},
			},
			want: NestedMap{
				"foo": map[string]any{},
			},
		},
		{
			name: "Delete existing map",
			key:  "foo.bar",
			existing: NestedMap{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "qux",
					},
				},
			},
			want: NestedMap{
				"foo": map[string]any{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.existing.Delete(tt.key)
			if !reflect.DeepEqual(tt.existing, tt.want) {
				t.Errorf("NestedMap.Set() got = %v, want %v", tt.existing, tt.want)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name   string
		dest   map[string]any
		src    map[string]any
		result map[string]any
	}{
		{
			name: "Merge two maps",
			dest: map[string]any{
				"foo": "bar",
			},
			src: map[string]any{
				"baz": "qux",
			},
			result: map[string]any{
				"foo": "bar",
				"baz": "qux",
			},
		},
		{
			name: "Merge two maps with overlapping keys",
			dest: map[string]any{
				"foo": "bar",
				"baz": "qux",
			},
			src: map[string]any{
				"baz": "quux",
			},
			result: map[string]any{
				"foo": "bar",
				"baz": "quux",
			},
		},
		{
			name: "Merge two maps with overlapping keys and nested maps",
			dest: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			src: map[string]any{
				"foo": map[string]any{
					"qux": "quux",
				},
			},
			result: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
					"qux": "quux",
				},
			},
		},
		{
			name: "Merge two maps with overlapping keys and nested maps",
			dest: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			src: map[string]any{
				"foo": "qux",
			},
			result: map[string]any{
				"foo": "qux",
			},
		},
		{
			name: "Merge two maps with overlapping keys and nested maps",
			dest: map[string]any{
				"foo": "qux",
			},
			src: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			result: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeMaps(tt.dest, tt.src)
			if !reflect.DeepEqual(tt.dest, tt.result) {
				t.Errorf("NestedMap.Set() got = %v, want %v", tt.dest, tt.result)
			}
		})
	}
}
