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
	"unicode/utf8"
	"strings"

	"go-sensitive/api/internal/svc"
	"go-sensitive/api/internal/types"

	"github.com/tal-tech/go-zero/core/mr"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/gogf/gf/errors/gerror"
	"github.com/importcjj/sensitive"
	"github.com/gogf/gf/frame/g"
)

type FindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type itemChan struct {
	name      string
	sensitive *sensitive.Filter
}

// 单个检测返回结构
type itemRes struct {
	name   string
	result []string
}

type findItemReply struct {
	hash       string
	isValidate bool
	badWords   map[string][]string
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindLogic {
	return FindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLogic) Find(req types.FindReq) (*types.FindReply, error) {
	fir, _ := findOne(req.Txt, req.Channel, req.Hash, l.svcCtx)

	return &types.FindReply{
		IsValidate: fir.isValidate,
		BadWords:   fir.badWords,
	}, nil
}

// find
func findOne(txt, channel, hash string, svc *svc.ServiceContext) (findItemReply, error) {
	// check
	txtCount := utf8.RuneCountInString(txt)
	if txtCount > svc.Config.Sensitive.MaxTxtLen {
		t := svc.I18n.T(context.TODO(), "The text to be detected is too long")
		err := gerror.New(t)
		return findItemReply{}, err
	}
	if txtCount <= 0 {
		t := svc.I18n.T(context.TODO(), "The text is empty")
		err := gerror.New(t)
		return findItemReply{}, err
	}

	channelSensitiveList, ok := svc.Config.Sensitive.ChannelSensitiveMap[channel]
	if !ok {
		t := svc.I18n.T(context.TODO(), "The channel is empty")
		err := gerror.New(t)
		return findItemReply{}, err
	}

	txt = strings.ToUpper(txt)

	// badWords
	badWords := make(map[string][]string)

	key := hash
	if key != "" {
		key = "Find" + key

		bws, _ := svc.Cache.Take(key, func() (interface{}, error) {
			badWords := mapReduceFindAll(channelSensitiveList, txt, svc)
			return badWords, nil
		})

		badWords = bws.(map[string][]string)
	} else {
		badWords = mapReduceFindAll(channelSensitiveList, txt, svc)
	}

	// isValidate
	isValidate := true
	if !g.IsEmpty(badWords) {
		isValidate = false
	}

	return findItemReply{
		hash,
		isValidate,
		badWords,
	}, nil
}

//mapReduce 多个词库并行检测
func mapReduceFindAll(channelSensitiveList []string, txt string, svc *svc.ServiceContext) map[string][]string {
	badWords := make(map[string][]string)

	mr.MapReduceVoid(func(source chan<- interface{}) {
		for _, item := range channelSensitiveList {
			source <- itemChan{
				item,
				svc.SensitiveMap[item],
			}
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		i := item.(itemChan)
		bw := i.sensitive.FindAll(txt)
		if len(bw) <= 0 {
			return
		}

		ir := itemRes{
			name:   i.name,
			result: bw,
		}
		writer.Write(ir)
	}, func(pipe <-chan interface{}, cancel func(error)) {
		for p := range pipe {
			ir, ok := p.(itemRes)
			if ok {
				badWords[ir.name] = ir.result
			}
		}
	})

	return badWords
}
