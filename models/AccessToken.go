package models

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// 参数说明：

// 参数	说明
// errcode	出错返回码，为0表示成功，非0表示调用失败
// errmsg	返回码提示语
// access_token	获取到的凭证，最长为512字节
// expires_in	凭证的有效时间（秒）

type AccessToken struct {
	PublicResult
	Access_token string
	Expires_in   int
}

var CacheToken = cache.New(2*time.Hour, 2*time.Hour)

func GetAccessToken() (token *AccessToken) {
	cache_token, period, found := CacheToken.GetWithExpiration("AccessToken")
	if found {
		fmt.Println(cache_token)
		token = cache_token.(*AccessToken)
		token.Expires_in = int(period.Unix() - time.Now().Unix())
	} else {
		token, _ = GetAccessTokenAPI()
		CacheToken.Set("AccessToken", token, cache.DefaultExpiration)
	}

	return
}
