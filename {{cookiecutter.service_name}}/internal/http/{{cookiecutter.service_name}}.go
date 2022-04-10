package http

import (
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/service"
	v1 "nekowindow-backend/app/service/{{cookiecutter.department}}/{{cookiecutter.service_name}}/api/v1"
	"nekowindow-backend/pkg/net/http/middleware/response"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// {{cookiecutter.serviceUpper}}FileSystemController is a http controller.
type {{cookiecutter.serviceUpper}}HttpController struct {
	s *service.{{cookiecutter.serviceUpper}}Service
	log     *log.Helper
}

func New{{cookiecutter.serviceUpper}}HttpController(service *service.{{cookiecutter.serviceUpper}}Service, logger log.Logger) *{{cookiecutter.serviceUpper}}HttpController {
	return &{{cookiecutter.serviceUpper}}HttpController{s: service, log: log.NewHelper(log.With(logger, "package", "http")))}
}

func (controller *{{cookiecutter.serviceUpper}}HttpController) ExampleHandler(ctx *gin.Context) {
	//Implement your handler here.
	req := v1.ExampleReq{}
	resp, err := service.ExampleService(ctx, req)
	if err != nil{
		response.AbortWithErrors(ctx, err)
	}
	response.Protobuf(ctx, resp)

}
