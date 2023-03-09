package middleware

import (
	"Skywing/models/response"
	"Skywing/pkg/jwt"
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strconv"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
)

func init() {
	db, err := sqlx.Connect("mysql", "zzzheng:skys!1004@tcp(rm-bp1g6r2lq2b4kb779so.mysql.rds.aliyuncs.com:3306)/skywing")
	if err != nil {
		panic(err)
	}
	// 初始化适配器
	a, err := sqlxadapter.NewAdapter(db, "casbin_rule")
	if err != nil {
		panic(err)
	}
	// 加载RBAC权限模型
	text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
		return
	}
	cachedEnforcer, err = casbin.NewCachedEnforcer(m, a)
	if err != nil {
		panic(err)
	}
}

// CasbinHandler 鉴权中间件

func CasbinHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		// 获取用户token
		token := c.Request.Header.Get("Authorization")
		// token解析
		parsedTokenStru, err := jwt.ParseToken(token)
		if err != nil {
			zap.L().Error("token解析失败", zap.Error(err))
			c.Abort()
			response.ResponseError(c, response.CodeInvalidToken)
			return
		}
		// 加载策略信息
		if err = cachedEnforcer.LoadPolicy(); err != nil {
			zap.L().Fatal("策略加载失败！", zap.Error(err))
			panic(err)
		}
		// 获取请求的sub
		sub := strconv.FormatUint(parsedTokenStru.Uuid, 10)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method

		success, _ := cachedEnforcer.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			c.Abort()
			response.ResponseError(c, response.CodePolicyFailed)
			return
		}
	}
}
