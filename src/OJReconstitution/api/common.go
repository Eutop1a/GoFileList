package api

import (
	"OJReconstitution/serializer"
	"encoding/json"
	"fmt"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}

}
