package goany

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	res := NewResponse().Set("message", "hello").Set("count", 3)
	data, _ := res.MarshalJSON()
	expected := `{"message":"hello","count":3}`
	require.Equal(t, string(data), expected)

	res2 := NewResponse().Sets("a", 1, "b", true, "c", "test")
	data2, _ := res2.MarshalJSON()
	expected2 := `{"a":1,"b":true,"c":"test"}`
	require.Equal(t, string(data2), expected2)

	res3 := NewResponse().Set("user", map[string]any{"name": "alice", "age": 30})
	data3, _ := res3.MarshalJSON()
	expected3 := `{"user":{"name":"alice","age":30}}`
	require.Equal(t, string(data3), expected3)

	require.Equal(t, http.StatusOK, NewResponse().HTTPStatus(nil))
	require.Equal(t, http.StatusOK, NewResponse().SetHTTPStatus(http.StatusOK).HTTPStatus(nil))
	require.Equal(t, 418, NewResponse().SetHTTPStatus(418).HTTPStatus(nil))
}
