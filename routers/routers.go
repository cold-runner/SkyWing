package routers

import (
	"Skywing/controller/middleware"
	"Skywing/controller/user"
	"Skywing/pkg/captcha"
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

	// v1组api路径于业务强耦合！改不动了....
	v1 := r.Group("/api/v1")
	{
		r.MaxMultipartMemory = 8 << 20 // 8 MiB
		userController := user.NewUserController(storeIns)
		v1.GET("/ApplicantCount", userController.GetCount)
		v1.GET("/captcha", captcha.Captcha)
		//v1.GET("/sendSmsCode", userController.SendSmsCode)
		v1.POST("/signUp", userController.Create)
		v1.POST("/login", userController.Login)

		v1.Use(middleware.JWTAuthMiddleware(), middleware.CasbinHandler())
		{
			v1.PUT("/update/:uuid", userController.Update)
			v1.GET("/info/:uuid", userController.GetInfo)
			v1.PUT("/logout", userController.Logout)
		}
	}

	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
