package goany

import (
	"fmt"
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
			"tags": ["admin", "editor"]
		},
		"metadata": "{\"verified\":true}"
	}`

	d := NewRequest(jsonStr)

	// get
	verified := d.Path("metadata").Get("verified").Bool()
	fmt.Println("Verified:", verified) // true

	// chaining path
	fmt.Println(d.Path("user.name").String())
	fmt.Println(d.Path("user.age").Int())
	fmt.Println(d.Path("user.address.city").String())

	// index
	fmt.Println(d.Path("user.tags").Index(1).String())

	// has, len, truthy
	fmt.Println(d.Path("user").Has("active"))
	fmt.Println(d.Path("user.tags").Len())
}
