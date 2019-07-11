package controllers

import (
	"golang-wxwork-sdk/httputil"
	"golang-wxwork-sdk/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登陆HR
// @Summary 登陆HR
// @Description 登陆HR
// @Tags Account
// @Accept  json
// @Produce  json
// @Param	employeeNo	query	string	true	"用户名"
// @Param	password	query	string	true	"密码"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /Account/Login [get]
func (c *Controller) Login(ctx *gin.Context) {
	employeeNo := ctx.Request.URL.Query().Get("employeeNo")
	password := ctx.Request.URL.Query().Get("password")
	usr, err := models.Login(employeeNo, password)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, usr)
}
