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
	"unicode/utf8"

	"go-sensitive/api/internal/svc"
	"go-sensitive/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/gogf/gf/errors/gerror"
)

type ValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ValidateLogic {
	return ValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidateLogic) Validate(req types.ValidateReq) (*types.ValidateReply, error) {
	// check
	txtCount := utf8.RuneCountInString(req.Txt)
	if txtCount > l.svcCtx.Config.Sensitive.MaxTxtLen {
		t := l.svcCtx.I18n.T(context.TODO(), "The text to be detected is too long")
		err := gerror.New(t)
		return &types.ValidateReply{}, err
	}
	if txtCount <= 0 {
		t := l.svcCtx.I18n.T(context.TODO(), "The text is empty")
		err := gerror.New(t)
		return &types.ValidateReply{}, err
	}

	key := req.Hash

	var b bool
	var w string
	if key != "" {
		key = "Validate" + req.Hash

		type tmpRes struct {
			b bool
			w string
		}

		tr, _ := l.svcCtx.Cache.Take(key, func() (interface{}, error) {
			b, w := l.svcCtx.Sensitive.Validate(req.Txt)
			var tr = tmpRes{
				b,
				w,
			}
			return tr, nil
		})

		bw := tr.(tmpRes)
		b = bw.b
		w = bw.w
	} else {
		b, w = l.svcCtx.Sensitive.Validate(req.Txt)
	}

	return &types.ValidateReply{
		IsValidate: b,
		BadWord:    w,
	}, nil
}
