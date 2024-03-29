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
	"github.com/tal-tech/go-zero/core/lang"
)

func AddWordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddWordReq
		if err := httpx.Parse(r, &req); err != nil {
			common.CheckErr(w, err)
			return
		}

		l := logic.NewAddWordLogic(r.Context(), ctx)
		err := l.AddWord(req)
		if err != nil {
			common.CheckErr(w, err)
			return
		}

		common.ResSuccess(w, lang.Placeholder)
	}
}
