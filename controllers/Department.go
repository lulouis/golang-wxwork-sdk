package controllers

import (
	"golang-wxwork-sdk/httputil"
	"golang-wxwork-sdk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取部门列表
// @Summary 获取部门列表
// @Description 获取部门列表
// @Tags Department 部门
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Department/GetDepartmentList [get]
func (c *Controller) GetDepartmentList(ctx *gin.Context) {
	data, err := models.GetDepartmentList()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 添加部门
// @Summary 添加部门
// @Description 添加部门
// @Tags Department 部门
// @Accept  json
// @Produce  json
// @Param	body	body 	models.Department2	true	"包体结构"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Department/DepartmentAdd [post]
func (c *Controller) DepartmentAdd(ctx *gin.Context) {
	v := models.Department2{}
	if err := ctx.ShouldBindJSON(&v); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := models.DepartmentAdd(v)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 更新部门
// @Summary 更新部门
// @Description 更新部门
// @Tags Department 部门
// @Accept  json
// @Produce  json
// @Param	body	body 	models.Department	true	"包体结构"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Department/DepartmentUpdate [post]
func (c *Controller) DepartmentUpdate(ctx *gin.Context) {
	v := models.Department{}
	if err := ctx.ShouldBindJSON(&v); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := models.DepartmentUpdate(v)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 删除无人员的部门
// @Summary 删除无人员的部门
// @Description 删除无人员的部门
// @Tags Department 部门
// @Accept  json
// @Produce  json
// @Param 	id 	query 	int 	true 	"删除id"
// @Success 200 {string} string "pong"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security OAuth2Application[admin]
// @Router /Department/DepartmentDelete [get]
func (c *Controller) DepartmentDelete(ctx *gin.Context) {
	_id := ctx.Request.URL.Query().Get("id")
	id, _ := strconv.Atoi(_id)

	data, err := models.DepartmentDelete(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
