package models

import (
	"encoding/json"
	"fmt"
	"golang-wxwork-sdk/utils"
	"net/url"
)

func GetAccessTokenAPI() (result *AccessToken, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken")
	params := url.Values{}
	params.Set("corpid", utils.Corpid)
	params.Set("corpsecret", utils.Secret)
	res, err := utils.HttpGet(reqUrl, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}
