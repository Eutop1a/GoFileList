package api

import (
	"OnlineJudge/serializer"
	"fmt"
)

func ErrorResponse(err error) serializer.Response {
	return serializer.Response{
		Status: 40001,

		Error: fmt.Sprint(err),
	}
}
