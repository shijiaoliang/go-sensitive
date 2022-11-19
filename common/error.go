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
	"github.com/gogf/gf/errors/gerror"
)

func init()  {
	gerror.New("error")
}

func CheckErr(w http.ResponseWriter, err error) {
	if err != nil {
		ResError(w, err.Error())
	}
}
