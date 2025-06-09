package goany

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Any struct {
	val any
}

func Wrap(v any) Any {
	if s, ok := v.(string); ok && len(s) > 0 && (s[0] == '{' || s[0] == '[') {
		var parsed any
		if err := json.Unmarshal([]byte(s), &parsed); err == nil {
			return Any{val: parsed}
		}
	}
	return Any{val: v}
}

func (a Any) Len() int {
	switch v := a.val.(type) {
	case string:
		return len(v)
	case []any:
		return len(v)
	case map[string]any:
		return len(v)
	default:
		return 0
	}
}

func (a Any) Has(key string) bool {
	if m, ok := a.val.(map[string]any); ok {
		_, exists := m[key]
		return exists
	}
	return false
}

func (a Any) Get(key string) Any {
	if m, ok := a.val.(map[string]any); ok {
		if v, found := m[key]; found {
			return Wrap(v)
		}
	}
	return Wrap(nil)
}

func (a Any) Path(p string) Any {
	parts := strings.Split(p, ".")
	cur := a
	for _, part := range parts {
		cur = cur.Get(part)
	}
	return cur
}

func (a Any) Index(i int) Any {
	if arr, ok := a.val.([]any); ok {
		if i >= 0 && i < len(arr) {
			return Wrap(arr[i])
		}
	}
	return Wrap(nil)
}

func (a Any) String() string {
	switch v := a.val.(type) {
	case string:
		return v
	case float64:
		return fmt.Sprintf("%g", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case nil:
		return "undefined"
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func (a Any) Int() int {
	if f, ok := a.val.(float64); ok {
		return int(f)
	}
	return 0
}

func (a Any) Float() float64 {
	if f, ok := a.val.(float64); ok {
		return f
	}
	return 0.0
}

func (a Any) Bool() bool {
	if b, ok := a.val.(bool); ok {
		return b
	}
	return false
}

func (a Any) Slice() []Any {
	if arr, ok := a.val.([]any); ok {
		out := make([]Any, len(arr))
		for i, v := range arr {
			out[i] = Wrap(v)
		}
		return out
	}
	return nil
}

func (a Any) Map() map[string]Any {
	if m, ok := a.val.(map[string]any); ok {
		out := make(map[string]Any)
		for k, v := range m {
			out[k] = Wrap(v)
		}
		return out
	}
	return nil
}

func (a Any) Value() any {
	return a.val
}

func (a Any) IsNil() bool {
	return a.val == nil
}
