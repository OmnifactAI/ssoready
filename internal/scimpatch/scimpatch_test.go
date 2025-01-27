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
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Work 1",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[type eq \"work\"].formatted", Value: "Remote 1"}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Remote 1",
					},
				},
			},
		},
		{
			name: "replace entire object with filter expression",
			in: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Work 1",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[type eq \"work\"]", Value: map[string]any{
				"type":      "remote",
				"formatted": "Remote 1",
			}}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "remote",
						"formatted": "Remote 1",
					},
				},
			},
		},
		{
			name: "replace with not-equal filter expression",
			in: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Work 1",
					},
					map[string]any{
						"type":      "other",
						"formatted": "Other 1",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[type ne \"home\"].formatted", Value: "Remote 1"}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Remote 1",
					},
					map[string]any{
						"type":      "other",
						"formatted": "Remote 1",
					},
				},
			},
		},
		{
			name: "replace entire object with not-equal filter expression",
			in: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "work",
						"formatted": "Work 1",
					},
					map[string]any{
						"type":      "other",
						"formatted": "Other 1",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[type ne \"home\"]", Value: map[string]any{
				"type":      "remote",
				"formatted": "Remote 1",
			}}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "Home 1",
					},
					map[string]any{
						"type":      "remote",
						"formatted": "Remote 1",
					},
					map[string]any{
						"type":      "remote",
						"formatted": "Remote 1",
					},
				},
			},
		},
		{
			name: "replace with contains filter expression",
			in: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "123 Main St",
					},
					map[string]any{
						"type":      "work",
						"formatted": "456 Main St",
					},
					map[string]any{
						"type":      "other",
						"formatted": "789 Side St",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[formatted co \"Main\"].type", Value: "primary"}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "primary",
						"formatted": "123 Main St",
					},
					map[string]any{
						"type":      "primary",
						"formatted": "456 Main St",
					},
					map[string]any{
						"type":      "other",
						"formatted": "789 Side St",
					},
				},
			},
		},
		{
			name: "replace with starts-with filter expression",
			in: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "home",
						"formatted": "123 Main St",
					},
					map[string]any{
						"type":      "work",
						"formatted": "123 Side St",
					},
					map[string]any{
						"type":      "other",
						"formatted": "456 Main St",
					},
				},
			},
			ops: []scimpatch.Operation{{Op: "Replace", Path: "addresses[formatted sw \"123\"].type", Value: "primary"}},
			out: map[string]any{
				"addresses": []any{
					map[string]any{
						"type":      "primary",
						"formatted": "123 Main St",
					},
					map[string]any{
						"type":      "primary",
						"formatted": "123 Side St",
					},
					map[string]any{
						"type":      "other",
						"formatted": "456 Main St",
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
