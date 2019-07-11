package controllers

import (
	"golang-wxwork-sdk/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取企业微信令牌
// @Summary 获取企业微信令牌
// @Description 获取企业微信令牌
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /Account/GetAccessToken [get]
func (c *Controller) GetAccessToken(ctx *gin.Context) {
	data := models.GetAccessToken()
	ctx.JSON(http.StatusOK, data)
}
