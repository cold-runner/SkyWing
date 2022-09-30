package routers

import (
	"Skywing/controller/middleware"
	"Skywing/controller/user"
	"Skywing/settings"
	"Skywing/store/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r := gin.Default()
	// 初始化存储实例
	storeIns, err := mysql.GetMySQLFactoryOr(settings.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/api/v1")
	{
		userController := user.NewUserController(storeIns)

		v1.POST("/signUp", userController.Create)
		v1.POST("/login", userController.Login)
		//v1.GET("/refresh_token", controller.RefreshTokenHandler)

		v1.Use(middleware.JWTAuthMiddleware())
		{
			v1.PUT("/update:stuNum", userController.Update)
			v1.GET("/ping", func(c *gin.Context) {
				c.String(http.StatusOK, "pong")
			})

		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
