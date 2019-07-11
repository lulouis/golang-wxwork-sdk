package basic

import (
	"golang-wxwork-sdk/controllers"
	"golang-wxwork-sdk/ginserver"

	"github.com/gin-gonic/gin"
)

func CreateBasicRouter(eng *gin.Engine) (err error) {
	c := controllers.NewController()
	v1 := eng.Group("/v1")
	{

		Account := v1.Group("/Account")
		{
			Account.GET("Login", c.Login)
			Account.GET("GetAccessToken", c.GetAccessToken)
		}
		Department := v1.Group("/Department")
		{
			Department.Use(ginserver.HandleTokenVerify())
			Department.GET("GetDepartmentList", c.GetDepartmentList)
			Department.POST("DepartmentAdd", c.DepartmentAdd)
			Department.POST("DepartmentUpdate", c.DepartmentUpdate)
			Department.GET("DepartmentDelete", c.DepartmentDelete)

		}
		Push := v1.Group("/Push")
		{
			Push.Use(ginserver.HandleTokenVerify())
			Push.POST("PushMessage", c.PushMessage)
			Push.POST("PushMessageWithFile", c.PushMessageWithFile)
			Push.POST("UploadMedia", c.UploadMedia)

		}

	}
	return nil
}
