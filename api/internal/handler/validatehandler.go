/**
 ********************************************************************************************
 * Created by go-sensitive.
 * User: shijl
 * Date: 2021/09/10
 * Time: 11:18
 ********************************************************************************************
 */

package handler

import (
	"net/http"

	"go-sensitive/api/internal/logic"
	"go-sensitive/api/internal/svc"
	"go-sensitive/api/internal/types"
	"go-sensitive/common"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func validateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewValidateLogic(r.Context(), ctx)
		resp, err := l.Validate(req)
		if err != nil {
			common.CheckErr(w, err)
			return
		}

		common.ResSuccess(w, resp)
	}
}
