/**
 ********************************************************************************************
 * Created by go-sensitive.
 * User: shijl
 * Date: 2021/09/10
 * Time: 11:18
 ********************************************************************************************
 */

package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Sensitive struct {
		MaxTxtLen           int
		WordsFilePath       string
		ChannelSensitiveMap map[string][]string
		Cache struct {
			Duration int64
			Name     string
			Limit    int
		}
	}
}
