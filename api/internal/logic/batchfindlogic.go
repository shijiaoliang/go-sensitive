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
	"github.com/tal-tech/go-zero/core/mr"
)

type BatchFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) BatchFindLogic {
	return BatchFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchFindLogic) BatchFind(req types.BatchFindReq) (*types.BatchFindReply, error) {
	batchFindReply := make(map[string]types.BatchItemReply)

	mr.MapReduceVoid(func(source chan<- interface{}) {
		for _, item := range req.Items {
			source <- item
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		i := item.(types.BatchItemReq)

		fir, _ := findOne(i.Txt, i.Channel, i.Hash, l.svcCtx)

		writer.Write(types.BatchItemReply{
			i.DataId,
			i.Hash,
			fir.isValidate,
			fir.badWords,
		})
	}, func(pipe <-chan interface{}, cancel func(error)) {
		for p := range pipe {
			ir, ok := p.(types.BatchItemReply)
			if ok {
				batchFindReply[ir.DataId] = ir
			}
		}
	})

	return &types.BatchFindReply{
		Items: batchFindReply,
	}, nil
}
