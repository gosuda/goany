package goany

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type Request struct {
	val any
}

func NewRequest(v any) *Request {
	switch val := v.(type) {
	case []byte:
		if len(val) > 0 && (val[0] == '{' || val[0] == '[') {
			var parsed any
			if err := jsoniter.Unmarshal(val, &parsed); err == nil {
				return &Request{val: parsed}
			}
		}
		return &Request{val: string(val)}
	case string:
		return NewRequest([]byte(val)) // delegate to []byte handler
	case *Request:
		return val
	case Request:
		return NewRequest(val.val)
	default:
		return &Request{val: v}
	}
}

func NewRequestFrom(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return &Request{}, fmt.Errorf("failed to read from reader: %w", err)
	}
	return NewRequest(data), nil
}

func (a *Request) Len() int {
	switch v := a.val.(type) {
	case []byte:
		return len(v)
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

func (a *Request) Has(key string) bool {
	if m, ok := a.val.(map[string]any); ok {
		_, exists := m[key]
		return exists
	}
	return false
}

func (a *Request) Get(key string) *Request {
	if m, ok := a.val.(map[string]any); ok {
		if v, found := m[key]; found {
			return NewRequest(v)
		}
	}
	return NewRequest(nil)
}

func (a *Request) Path(p string) *Request {
	tokens := parsePath(p)
	cur := a
	for _, tok := range tokens {
		switch key := tok.(type) {
		case string:
			cur = cur.Get(key)
		case int:
			cur = cur.Index(key)
		}
	}
	return cur
}

func parsePath(path string) []any {
	var result []any
	var current strings.Builder

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case '.':
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
		case '[':
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
			j := i + 1
			for ; j < len(path) && path[j] != ']'; j++ {
				// pass
			}
			if j > i+1 {
				index := path[i+1 : j]
				var idx int
				fmt.Sscanf(index, "%d", &idx)
				result = append(result, idx)
			}
			i = j // skip ']'
		default:
			current.WriteByte(path[i])
		}
	}
	if current.Len() > 0 {
		result = append(result, current.String())
	}
	return result
}

func (a *Request) Index(i int) *Request {
	if arr, ok := a.val.([]any); ok {
		if i >= 0 && i < len(arr) {
			return NewRequest(arr[i])
		}
	}
	return NewRequest(nil)
}

func (a *Request) String() string {
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

func (a *Request) Int() int {
	if f, ok := a.val.(float64); ok {
		return int(f)
	}
	return 0
}

func (a *Request) Float() float64 {
	if f, ok := a.val.(float64); ok {
		return f
	}
	return 0.0
}

func (a *Request) Bool() bool {
	if b, ok := a.val.(bool); ok {
		return b
	}
	return false
}

func (a *Request) Slice() []*Request {
	if arr, ok := a.val.([]any); ok {
		out := make([]*Request, len(arr))
		for i, v := range arr {
			out[i] = NewRequest(v)
		}
		return out
	}
	return []*Request{}
}

func (a *Request) Map() map[string]*Request {
	if m, ok := a.val.(map[string]any); ok {
		out := make(map[string]*Request)
		for k, v := range m {
			out[k] = NewRequest(v)
		}
		return out
	}
	return map[string]*Request{}
}

func (a *Request) Value() any {
	return a.val
}

func (a *Request) IsNil() bool {
	return a.val == nil
}

func (a *Request) MarshalJSON() ([]byte, error) {
	if a.val == nil {
		return []byte("null"), nil
	}
	return jsoniter.Marshal(a.val)
}

func (a *Request) WriteTo(w io.Writer) (int64, error) {
	data, err := a.MarshalJSON()
	if err != nil {
		return 0, err
	}
	n, err := w.Write(data)
	return int64(n), err
}
