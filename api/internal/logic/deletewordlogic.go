/**
 ********************************************************************************************
 * Created by go-sensitive.
 * User: shijl
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

type DeleteWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteWordLogic {
	return DeleteWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWordLogic) DeleteWord(req types.DeleteWordReq) (error) {
	l.svcCtx.Sensitive.DelWord(req.Words...)

	return nil
}
