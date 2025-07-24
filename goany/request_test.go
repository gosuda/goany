package goany

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name     string
		input    *Request
		key      string
		expected string
	}{
		{"Get existing string", NewRequest(`{"key": "value"}`), "key", "value"},
		{"Get non-existing key", NewRequest(`{"key": "value"}`), "nonexistent", "undefined"},
		{"Get from empty map", NewRequest(`{}`), "key", "undefined"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Get(tt.key).String()
			require.Equal(t, tt.expected, result, tt.name)

		})
	}
}

func TestExample(t *testing.T) {
	jsonStr := `{
		"user": {
			"name": "KIM",
			"age": 30,
			"active": true,
			"address": {
				"city": "Seoul"
			},
			"friends": [
				{"name": "Alice"},
				{"name": "Bob"},
				{"name": "Charlie"}
			],
			"tags": ["admin", "editor"]
		},
		"metadata": "{\"verified\":true}"
	}`

	d := NewRequest(jsonStr)

	// get
	require.True(t, d.Path("metadata").Get("verified").Bool(), "Expected verified to be true")
	require.Equal(t, "true", d.Path("metadata").Get("verified").String(), "Expected verified to be true")
	require.Equal(t, 0, d.Path("metadata").Get("verified").Int(), "Expected verified to be true")

	// chaining path
	require.Equal(t, "KIM", d.Path("user.name").String())
	require.Equal(t, 30, d.Path("user.age").Int())
	require.Equal(t, "Seoul", d.Path("user.address.city").String())
	require.Equal(t, "Alice", d.Path("user.friends[0].name").String())
	require.Equal(t, "Bob", d.Path("user.friends[1].name").String())

	// index
	require.Equal(t, "editor", d.Path("user.tags").Index(1).String())

	// has, len, truthy
	require.True(t, d.Path("user").Has("active"), "Expected user to have 'active' key")
	require.Equal(t, 2, d.Path("user.tags").Len(), "Expected user.tags to have length 2")
}
