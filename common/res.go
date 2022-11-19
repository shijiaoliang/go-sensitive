/**
 ********************************************************************************************
 * Created by go-sensitive.
 * User: shijl
 * Date: 2021/09/10
 * Time: 11:18
 ********************************************************************************************
 */

package common

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// 系统通用返回
const (
	SUCCESS = 1  //成功
	ZERO    = 0  //0值
	ERROR   = -1 //失败
)

const (
	constSuccessMsg = "success"
)

type Res struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func ResSuccess(w http.ResponseWriter, data interface{}, msg ...string) {
	var outMsg = constSuccessMsg
	if len(msg) > 0 && msg[0] != "" {
		outMsg = msg[0]
	}

	httpx.OkJson(w, &Res{
		data,
		SUCCESS,
		outMsg,
	})
}

func ResError(w http.ResponseWriter, msg string, code ...int) {
	var outCode = ERROR
	if len(code) > 0 {
		outCode = code[0]
	}

	httpx.OkJson(w, &Res{
		"",
		outCode,
		msg,
	})
}

func ResErrorData(w http.ResponseWriter, msg string, data interface{}, code ...int) {
	var outCode = ERROR
	if len(code) > 0 {
		outCode = code[0]
	}

	httpx.OkJson(w, &Res{
		data,
		outCode,
		msg,
	})
}
