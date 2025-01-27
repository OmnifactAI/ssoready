package scimpatch_test

import (
	"testing"

	"github.com/ssoready/ssoready/internal/scimpatch"
	"github.com/stretchr/testify/assert"
)

func TestPatch(t *testing.T) {
	testCases := []struct {
		name string
		in   map[string]any
		ops  []scimpatch.Operation
		out  map[string]any
	}{
		{
			name: "replace entire value",
			in:   map[string]any{"foo": "xxx"},
			ops:  []scimpatch.Operation{{Op: "replace", Path: "", Value: map[string]any{"bar": "yyy"}}},
			out:  map[string]any{"bar": "yyy"},
		},
		{
			name: "replace top-level prop",
			in:   map[string]any{"foo": "xxx"},
			ops:  []scimpatch.Operation{{Op: "replace", Path: "foo", Value: "yyy"}},
			out:  map[string]any{"foo": "yyy"},
		},
		{
			name: "replace nested prop",
			in:   map[string]any{"foo": map[string]any{"bar": "xxx"}},
			ops:  []scimpatch.Operation{{Op: "replace", Path: "foo.bar", Value: "yyy"}},
			out:  map[string]any{"foo": map[string]any{"bar": "yyy"}},
		},
		{
			name: "replace map prop",
			in:   map[string]any{"foo": map[string]any{"bar": "xxx"}},
			ops:  []scimpatch.Operation{{Op: "replace", Path: "foo", Value: map[string]any{"bar": "yyy"}}},
			out:  map[string]any{"foo": map[string]any{"bar": "yyy"}},
		},
		{
			name: "replace scalar with map",
			in:   map[string]any{"foo": map[string]any{"bar": "xxx"}},
			ops:  []scimpatch.Operation{{Op: "replace", Path: "foo.bar", Value: map[string]any{"baz": "yyy"}}},
			out:  map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "yyy"}}},
		},
		{
			name: "add to slice",
			in:   map[string]any{"foo": []any{"xxx"}},
			ops:  []scimpatch.Operation{{Op: "add", Path: "foo", Value: []any{"yyy"}}},
			out:  map[string]any{"foo": []any{"xxx", "yyy"}},
		},
		{
			name: "add multiple to slice", // this is inferred from spec; unclear if used in the wild
			in:   map[string]any{"foo": []any{"xxx"}},
			ops:  []scimpatch.Operation{{Op: "add", Path: "foo", Value: []any{"yyy", "zzz"}}},
			out:  map[string]any{"foo": []any{"xxx", "yyy", "zzz"}},
		},
		{
			name: "add to empty property",
			in:   map[string]any{},
			ops:  []scimpatch.Operation{{Op: "add", Path: "foo", Value: "yyy"}},
			out:  map[string]any{"foo": "yyy"},
		},
		{
			name: "add to sub-object",
			in:   map[string]any{"foo": map[string]any{"bar": "xxx"}},
			ops:  []scimpatch.Operation{{Op: "add", Path: "foo", Value: map[string]any{"baz": "yyy"}}},
			out:  map[string]any{"foo": map[string]any{"bar": "xxx", "baz": "yyy"}},
		},

		{
			name: "uppercase Replace op",
			in:   map[string]any{"foo": "xxx"},
			ops:  []scimpatch.Operation{{Op: "Replace", Path: "", Value: map[string]any{"bar": "yyy"}}},
			out:  map[string]any{"bar": "yyy"},
		},
		{
			name: "uppercase Add op",
			in:   map[string]any{"foo": []any{"xxx"}},
			ops:  []scimpatch.Operation{{Op: "Add", Path: "foo", Value: []any{"yyy"}}},
			out:  map[string]any{"foo": []any{"xxx", "yyy"}},
		},

		{
			name: "special-case for entra patches on enterprise user",
			in: map[string]any{
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": map[string]any{
					"foo": "xxx",
				},
			},
			ops: []scimpatch.Operation{
				{
					Op:    "Add",
					Path:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:bar",
					Value: "yyy",
				},
			},
			out: map[string]any{
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": map[string]any{
					"foo": "xxx",
					"bar": "yyy",
				},
			},
		},
		{
			// inferred behavior; not seen in wild -- case where there's no sub-":" in the path
			name: "special-case for entra patches on enterprise user",
			in:   map[string]any{},
			ops: []scimpatch.Operation{
				{
					Op:    "Add",
					Path:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
					Value: map[string]any{"foo": "xxx"},
				},
			},
			out: map[string]any{
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": map[string]any{
					"foo": "xxx",
				},
			},
		},
		{
			name: "replace with filter expression in path",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "yyy",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[type eq \"bar\"].str", Value: "zzz"}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "zzz",
					},
				},
			},
		},
		{
			name: "replace entire object with filter expression",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "yyy",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[type eq \"bar\"]", Value: map[string]any{
				"type": "baz",
				"str":  "zzz",
			}}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "baz",
						"str":  "zzz",
					},
				},
			},
		},
		{
			name: "replace with not-equal filter expression",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "yyy",
					},
					map[string]any{
						"type": "baz",
						"str":  "zzz",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[type ne \"foo\"].str", Value: "aaa"}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "aaa",
					},
					map[string]any{
						"type": "baz",
						"str":  "aaa",
					},
				},
			},
		},
		{
			name: "replace entire object with not-equal filter expression",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "bar",
						"str":  "yyy",
					},
					map[string]any{
						"type": "baz",
						"str":  "zzz",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[type ne \"foo\"]", Value: map[string]any{
				"type": "aaa",
				"str":  "bbb",
			}}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx",
					},
					map[string]any{
						"type": "aaa",
						"str":  "bbb",
					},
					map[string]any{
						"type": "aaa",
						"str":  "bbb",
					},
				},
			},
		},
		{
			name: "replace with contains filter expression",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx_abc",
					},
					map[string]any{
						"type": "bar",
						"str":  "yyy_abc",
					},
					map[string]any{
						"type": "baz",
						"str":  "zzz",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[str co \"abc\"].type", Value: "aaa"}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "aaa",
						"str":  "xxx_abc",
					},
					map[string]any{
						"type": "aaa",
						"str":  "yyy_abc",
					},
					map[string]any{
						"type": "baz",
						"str":  "zzz",
					},
				},
			},
		},
		{
			name: "replace with starts-with filter expression",
			in: map[string]any{
				"items": []any{
					map[string]any{
						"type": "foo",
						"str":  "xxx_abc",
					},
					map[string]any{
						"type": "bar",
						"str":  "xxx_def",
					},
					map[string]any{
						"type": "baz",
						"str":  "yyy_abc",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "items[str sw \"xxx\"].type", Value: "aaa"}},
			out: map[string]any{
				"items": []any{
					map[string]any{
						"type": "aaa",
						"str":  "xxx_abc",
					},
					map[string]any{
						"type": "aaa",
						"str":  "xxx_def",
					},
					map[string]any{
						"type": "baz",
						"str":  "yyy_abc",
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := scimpatch.Patch(tt.ops, &tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}
