/**
 *******************************************51cto********************************************
 * Created by go-sensitive.
 * User: 605724193@qq.com
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

func batchFindHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchFindReq
		if err := httpx.Parse(r, &req); err != nil {
			common.CheckErr(w, err)
			return
		}

		l := logic.NewBatchFindLogic(r.Context(), ctx)
		resp, err := l.BatchFind(req)

		if err != nil {
			common.CheckErr(w, err)
			return
		}

		common.ResSuccess(w, resp)
	}
}
