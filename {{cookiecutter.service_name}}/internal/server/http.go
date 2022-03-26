package server

import (
	v1 "nekowindow-backend/api/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/conf"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.AccountService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	router := gin.Default()
	handler := router.Group("/")

	handler.Handle("OPTIONS", "/", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	account := router.Group("/x/web-interface/account/")

	account.Handle("POST", "/register", func(ctx *gin.Context) {
		resp, err := service.CreateUser(ctx, &v1.CreateUserRequest{
			Email:    ctx.PostForm("email"),
			Password: ctx.PostForm("password"),
			Code:     ctx.PostForm("code"),
		})

		if err != nil {
			ctx.AbortWithStatusJSON(200, err)
			return
		}

		expire := 31536000000 // 365*24*60*60*1000
		ctx.SetCookie("uid", strconv.FormatInt(resp.Uid, 10), expire, "/", ".inkneko.com", false, true)
		ctx.SetCookie("SESSIONKEY", resp.SessionKey, expire, "/", ".inkneko.com", false, true)

		ctx.JSON(200, gin.H{"code": 0, "message": "注册成功"})
	})

	account.Handle("POST", "/login", func(ctx *gin.Context) {
		resp, err := service.EmailLogin(ctx, &v1.EmailLoginRequest{
			Email:    ctx.PostForm("email"),
			Password: ctx.PostForm("password"),
		})

		if err != nil {
			ctx.AbortWithStatusJSON(200, err)
			return
		}

		expire := 31536000000 // 365*24*60*60*1000
		ctx.SetCookie("uid", strconv.FormatInt(resp.Uid, 10), expire, "/", ".inkneko.com", false, true)
		ctx.SetCookie("SESSIONKEY", resp.SessionKey, expire, "/", ".inkneko.com", false, true)
		ctx.Redirect(301, "https://window.inkneko.com")
	})

	account.Handle("POST", "/sendRegisterEmail", func(ctx *gin.Context) {
		resp, err := service.SendRegisterEmail(ctx, &v1.SendRegisterEmailRequest{Email: ctx.PostForm("email")})
		if err != nil {
			ctx.AbortWithStatusJSON(200, err)
			return
		}

		ctx.JSON(200, gin.H{"code": 0, "data": resp})
	})

	srv.HandlePrefix("/", router)
	return srv
}
