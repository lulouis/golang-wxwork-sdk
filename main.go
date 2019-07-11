package main

import (
	"fmt"
	_ "golang-wxwork-sdk/docs"
	"net/http"
	"os"
	"runtime"
	"time"

	"golang-wxwork-sdk/ginserver"
	"golang-wxwork-sdk/routers"
	utils "golang-wxwork-sdk/utils"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/lulouis/gin-swagger"
	"github.com/lulouis/gin-swagger/swaggerFiles"

	redis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/manage"
	aserver "gopkg.in/oauth2.v3/server"
)

// @title 企业微信-GO语言服务端Restful API
// @version 2020.01.20
// @description 企业微信企业网关，带有负载均衡能力. (带安全验证，令牌已设置4小时有效期)
// @termsOfService http://swagger.io/terms/

// hide@contact.name API Support
// hide@contact.url http://www.swagger.io/support
// hide@contact.email sunjoin@qq.com

// hide@license.name Apache 2.0
// hide@license.url http://www.apache.org/licenses/LICENSE-2.0.html

// // @host localhost:8766
// @BasePath /v1

// // @securityDefinitions.basic BasicAuth

// // @securityDefinitions.apikey ApiKeyAuth
// // @in header
// // @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl /oauth2/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	param1 := os.Args
	if len(param1) < 2 {
		fmt.Printf("必须带参数执行 DEV or PRD or TEST \n")
		return
	}

	if param1[1] != "DEV" && param1[1] != "PRD" && param1[1] != "TEST" {
		fmt.Printf("不支持该运行模式:%s,请检查app.conf文件 \n", param1[1])
		return
	}
	utils.CurrentMode = param1[1]
	utils.RunMode(utils.CurrentMode)

	// 负载均衡版本的令牌保存
	utils.TokenManager = manage.NewDefaultManager()
	// 令牌的有效期
	utils.TokenManager.SetClientTokenCfg(&manage.Config{AccessTokenExp: time.Hour * 4})
	// 这里需要根据实际的令牌持久化存储服务器进行相关设置
	utils.TokenManager.MapTokenStorage(redis.NewRedisStore(&redis.Options{
		Addr: "192.168.11.234:6379",
		DB:   1,
	}))

	// Initialize the oauth2 service
	ginserver.InitServer(utils.TokenManager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(aserver.ClientFormHandler)
	// ginserver.SetUserAuthorizationHandler(aserver.UserAuthorizationHandler)
	r := gin.Default()
	gin.SetMode("release")
	r.StaticFS("/UploadFile", http.Dir("UploadFile"))
	r.Use(Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	oauth2 := r.Group("/oauth2")
	{
		oauth2.GET("/token", ginserver.HandleTokenRequest)
		oauth2.POST("/token", ginserver.HandleTokenRequest)
		oauth2.GET("/authorize", ginserver.HandleAuthorizeRequest)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	routers.CreateRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("运行模式:%s \n", utils.CurrentMode)
	fmt.Printf("运行连接:%s \n", utils.MONGODB)
	fmt.Printf("运行端口:%d \n", utils.PORT)
	fmt.Printf("快速访问:http://localhost:%d \n", utils.PORT)
	r.Run(fmt.Sprintf(":%d", utils.PORT))
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
