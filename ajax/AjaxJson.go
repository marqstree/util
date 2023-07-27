package ajax

import (
	"encoding/json"
	"fmt"
	"github.com/marqstree/util/bizerror"
	"github.com/marqstree/util/constant"
	"net/http"
)

type AjaxJson[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func New[T any](code int, msg string, data T) *AjaxJson[T] {
	return &AjaxJson[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessByData[T any](data T) *AjaxJson[T] {
	return &AjaxJson[T]{
		Code: constant.SUCESS_CODE,
		Msg:  "成功",
		Data: data,
	}
}

func Success() *AjaxJson[any] {
	return &AjaxJson[any]{
		Code: constant.SUCESS_CODE,
		Msg:  "成功",
	}
}

func FailByError(err error) *AjaxJson[interface{}] {
	switch e := err.(type) {
	case *bizerror.BizError:
		return &AjaxJson[interface{}]{
			Code: e.Code,
			Msg:  e.Msg,
		}
	default:
		return &AjaxJson[interface{}]{
			Code: constant.FAIL_CODE,
			Msg:  e.Error(),
		}
	}
}

func Fail(msg string) *AjaxJson[any] {
	return &AjaxJson[any]{
		Code: constant.FAIL_CODE,
		Msg:  msg,
	}
}

func ResponseAjaxJson(writer http.ResponseWriter, aj AjaxJson[any]) {
	str, _ := json.Marshal(aj)
	fmt.Fprintf(writer, "%s", str)
}
