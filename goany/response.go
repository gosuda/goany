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

func NewResponseFrom(r io.Reader) (*Response, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var parsed map[string]any
	if err := jsoniter.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}
	return &Response{val: parsed}, nil
}

func (o *Response) Set(k string, v any) *Response {
	if o.val == nil {
		o.val = make(map[string]any)
	}
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

func (o *Response) Get(k string) any {
	return o.val[k]
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

func (o *Response) WriteTo(w io.Writer) (int64, error) {
	data, err := o.MarshalJSON()
	if err != nil {
		return 0, err
	}
	n, err := w.Write(data)
	return int64(n), err
}

func (o *Response) MarshalJSON() ([]byte, error) {
	if o.val == nil {
		return []byte("null"), nil
	}
	return jsoniter.Marshal(o.val)
}

func (o *Response) Value() map[string]any {
	return o.val
}
