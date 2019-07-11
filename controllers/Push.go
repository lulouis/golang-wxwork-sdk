package controllers

import (
	"golang-wxwork-sdk/httputil"
	"golang-wxwork-sdk/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 消息推送
// @Summary 消息推送
// @Description 消息推送
// @Tags Push 消息推送
// @Accept  json
// @Produce  json
// @Param	body	body 	models.Message	true	"包体结构"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Push/PushMessage [post]
func (c *Controller) PushMessage(ctx *gin.Context) {
	v := models.Message{}
	if err := ctx.ShouldBindJSON(&v); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := models.PushMessage(v)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

/*
测试文本
{
	"agentid": 0,
	"msgtype": "text",
	"safe": 0,
	"text": {
	  "content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"
	},
	"touser": "00275"
  }
测试卡片
{
	"touser" : "00275|01953",
	"msgtype" : "textcard",
	"agentid" : 22,
	"textcard" : {
			 "title" : "领奖通知",
			 "description" : "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
			 "url" : "URL"
	},
	"safe":1
 }
测试文件
{
  "agentid": 22,
  "msgtype": "file",
  "safe": 0,
   "file" : {
		"media_id" : "3uddU0MrKnqxy1hKXA9O9oBIUJ4n7b1UgWbby-jGqfo34OuJ5Y_phEuucENi-3Pv1"
   },
  "touser": "00275"
}

*/

// 临时素材
// @Summary 临时素材
// @Description 临时素材
// @Tags Push 消息推送
// @Accept  json
// @Produce  json
// @Param 	fileUrl 	query 	string 	true 	"上传企业微信临时素材的网络文件"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Push/UploadMedia [post]
func (c *Controller) UploadMedia(ctx *gin.Context) {
	fileUrl := ctx.Request.URL.Query().Get("fileUrl")

	data, err := models.UploadMedia(fileUrl)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 消息推送(带文件)
// @Summary 消息推送(带文件)
// @Description 消息推送(带文件)
// @Tags Push 消息推送
// @Accept  json
// @Produce  json
// @Param 	fileUrl 	query 	string 	true 	"上传企业微信临时素材的网络文件"
// @Param	body	body 	models.Message	true	"包体结构"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Push/PushMessageWithFile [post]
func (c *Controller) PushMessageWithFile(ctx *gin.Context) {
	fileUrl := ctx.Request.URL.Query().Get("fileUrl")
	v := models.Message{}
	if err := ctx.ShouldBindJSON(&v); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := models.PushMessageWithFile(fileUrl, v)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
