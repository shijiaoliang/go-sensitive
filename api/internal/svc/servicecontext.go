/**
 *******************************************51cto********************************************
 * Created by go-sensitive.
 * User: 605724193@qq.com
 * Date: 2021/09/10
 * Time: 11:18
 ********************************************************************************************
 */

package svc

import (
	"time"

	"go-sensitive/api/internal/config"

	"github.com/tal-tech/go-zero/core/collection"
	"github.com/importcjj/sensitive"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/i18n/gi18n"
)

type ServiceContext struct {
	Config       config.Config
	Cache        *collection.Cache
	I18n         *gi18n.Manager
	Sensitive    *sensitive.Filter            //总的敏感词库
	SensitiveMap map[string]*sensitive.Filter //分类敏感词库
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Sensitive & SensitiveMap
	s := sensitive.New()
	sm := make(map[string]*sensitive.Filter)
	wordsFiles, err := gfile.ScanDirFile(c.Sensitive.WordsFilePath, "*")
	if err != nil {
		panic(err)
	}
	if len(wordsFiles) == 0 {
		panic("sensitive file is empty")
	}
	for _, f := range wordsFiles {
		errLoad := s.LoadWordDict(f)
		if errLoad != nil {
			panic("[" + f + "] is not found")
		}

		tmpS := sensitive.New()
		tmpS.LoadWordDict(f)
		sm[gfile.Name(f)] = tmpS
	}

	// Cache
	cache, _ := collection.NewCache(time.Second*time.Duration(c.Sensitive.Cache.Duration),
		collection.WithName(c.Sensitive.Cache.Name),
		collection.WithLimit(c.Sensitive.Cache.Limit),
	)

	// I18n
	i18n := gi18n.New()
	i18n.SetPath("etc/i18n")
	i18n.SetLanguage("zh-CN")

	// res
	return &ServiceContext{
		Config:       c,
		Cache:        cache,
		I18n:         i18n,
		Sensitive:    s,
		SensitiveMap: sm,
	}
}
