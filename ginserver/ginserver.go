package ginserver

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-wxwork-sdk/utils"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

var (
	gServer *server.Server
)

// InitServer Initialize the service
func InitServer(manager oauth2.Manager) *server.Server {
	// if err := manager.CheckInterface(); err != nil {
	// 	panic(err)'
	// }
	gServer = server.NewDefaultServer(manager)
	return gServer
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c *gin.Context) {
	err := gServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c *gin.Context) {
	//生成新的令牌前，作废此前的令牌 todo
	ti, err := utils.GenerateToken(c.Request.FormValue("client_id"), c.Request.FormValue("client_secret"))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	c.Writer.Header().Set("Cache-Control", "no-store")
	c.Writer.Header().Set("Pragma", "no-cache")
	status := http.StatusOK
	c.Writer.WriteHeader(status)

	data := gServer.GetTokenData(ti)
	json.NewEncoder(c.Writer).Encode(data)

	// err := gServer.HandleTokenRequest(c.Writer, c.Request)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	// c.Abort()
}

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {

		ti, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("AccessToken", ti)
		c.Next()
	}
}
