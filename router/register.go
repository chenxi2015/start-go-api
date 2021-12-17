package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/voyager-go/start-go-api/api"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/docs"
	"github.com/voyager-go/start-go-api/middleware"
	"net/http"
)

func Register() *gin.Engine {
	gin.SetMode(config.Conf.Mode)
	router := gin.New()
	// 404 处理
	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		fmt.Println(path)
		method := ctx.Request.Method
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%s %s not found", method, path))
	})
	// swagger 配置
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 健康检查
	router.GET("/ping", api.Check.Ping)
	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		// 用户组
		apiGroup     = router.Group("/api", publicMiddleware...)
		apiNeedLogin = router.Group("/api", append(publicMiddleware, middleware.NeedLogin)...)
	)
	apiGroup.POST("/user/auth", api.SysUser.Login)
	apiNeedLogin.GET("/user/:id", api.SysUser.Find)
	apiNeedLogin.PUT("/user", api.SysUser.Update)
	apiNeedLogin.POST("/user", api.SysUser.Create)

	return router
}
