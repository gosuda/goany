package goany

import jsoniter "github.com/json-iterator/go"

type Response struct {
	val map[string]any
}

func NewResponse() Response {
	return Response{val: make(map[string]any)}
}

func (o *Response) Set(k string, v any) *Response {
	o.val[k] = v
	return o
}

func (o *Response) Wrap() Request {
	return NewRequest(o.val)
}

func (o Response) MarshalJSON() ([]byte, error) {
	if o.val == nil {
		return []byte("null"), nil
	}
	return jsoniter.Marshal(o.val)
}
