package goany

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Response struct {
	val        map[string]any
	httpStatus int
}

func NewResponse() *Response {
	return &Response{val: make(map[string]any)}
}

func (o *Response) Set(k string, v any) *Response {
	o.val[k] = v
	return o
}

func (o *Response) Sets(kv ...any) *Response {
	for i := 0; i+1 < len(kv); i += 2 {
		key, ok := kv[i].(string)
		if !ok {
			continue
		}
		o.val[key] = kv[i+1]
	}
	return o
}

func (o *Response) SetHTTPStatus(code int) *Response {
	o.httpStatus = code
	return o
}

func (o *Response) HTTPStatus(err error) int {
	if o.httpStatus != 0 {
		return o.httpStatus
	}
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (o *Response) ToRequest() *Request {
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
