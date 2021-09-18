/**
 *******************************************51cto********************************************
 * Created by go-sensitive.
 * User: 605724193@qq.com
 * Date: 2021/09/10
 * Time: 11:18
 ********************************************************************************************
 */

package logic

import (
	"context"

	"go-sensitive/api/internal/svc"
	"go-sensitive/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddWordLogic {
	return AddWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddWordLogic) AddWord(req types.AddWordReq) (error) {
	l.svcCtx.Sensitive.AddWord(req.Words...)

	return nil
}
