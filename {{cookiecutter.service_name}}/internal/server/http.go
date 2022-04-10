package server

import (
	v1 "nekowindow-backend/app/service/{{cookiecutter.department}}/{{cookiecutter.service_name}}/api/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/conf"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/service"
	v1http "nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/http"
	"nekowindow-backend/pkg/net/http/middleware/auth"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"nekowindow-backend/pkg/net/http/middleware/cors"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, authMiddleware *auth.AuthMiddleware, con *v1http.{{cookiecutter.serviceUpper}}HttpController, logger log.Logger) *http.Server {
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
	router.Use(cors.NewCorsMiddleware())

	{{cookiecutter.service_name}} := router.Group("/x/web-interface/{{cookiecutter.service_name}}/")

	{{cookiecutter.service_name}}.Handle("GET", "/example", con.ExampleHandler)
	{{cookiecutter.service_name}}.Handle("GET", "/auth_or_exit", authMiddleware.UserAuth, con.ExampleHandler)
	

	srv.HandlePrefix("/", router)
	return srv
}
