package routers

import (
	"golang-wxwork-sdk/routers/basic"

	"github.com/gin-gonic/gin"
)

func CreateRouter(eng *gin.Engine) (err error) {
	//逐层挂载路由
	if err = basic.CreateBasicRouter(eng); err != nil {
		return err
	}

	return err
}
