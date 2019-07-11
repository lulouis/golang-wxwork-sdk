package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"golang-wxwork-sdk/utils"
	"net/http"
	"net/url"
)

// 应用支持推送文本、图片、视频、文件、图文等类型。
// 请求方式：POST（HTTPS）
// 请求地址： https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN

func PushMessage(message Message) (result *PublicResult, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	jdata, err := json.Marshal(message)
	res, err := utils.HttpPostBody(reqUrl, params, string(jdata))
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}

func PushMessageWithFile(fileUrl string, message Message) (result *PublicResult, err error) {
	mediaRes, err := UploadMedia(fileUrl)
	if err != nil {
		return
	}
	if mediaRes.Errcode != 0 {
		result.Errcode = mediaRes.Errcode
		result.Errmsg = mediaRes.Errmsg
		return
	}
	message.File.Media_id = mediaRes.Media_id

	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	jdata, err := json.Marshal(message)
	res, err := utils.HttpPostBody(reqUrl, params, string(jdata))
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}

// 上传临时素材
// 素材上传得到media_id，该media_id仅三天内有效
// media_id在同一企业内应用之间可以共享

// 请求方式：POST（HTTPS）
// 请求地址：https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE

func UploadMedia(fileUrl string) (result *UploadMediaRes, err error) {
	// 上传的媒体文件限制
	// 所有文件size必须大于5个字节(不可太小)
	// 图片（image）：2MB，支持JPG,PNG格式
	// 语音（voice） ：2MB，播放长度不超过60s，仅支持AMR格式
	// 视频（video） ：10MB，支持MP4格式
	// 普通文件（file）：20MB

	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	params.Set("type", "file") //限定死本应用仅仅支持文件发送

	downloadRes, err := http.Get(fileUrl)
	if err != nil {
		return
	}
	defer downloadRes.Body.Close()
	uri, err := url.ParseRequestURI(fileUrl)
	filename := path.Base(uri.Path)
	filename = "download//" + filename

	fmt.Printf("文件名称：%s \n", filename)

	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, downloadRes.Body)
	if err != nil {
		return
	}

	res, err := utils.HttpPostFile(filename, reqUrl, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}
