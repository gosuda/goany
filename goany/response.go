package goany

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

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

func (o *Response) WriteTo(writer io.Writer) (int64, error) {
	data, err := o.MarshalJSON()
	if err != nil {
		return 0, err
	}
	n, err := writer.Write(data)
	return int64(n), err
}
