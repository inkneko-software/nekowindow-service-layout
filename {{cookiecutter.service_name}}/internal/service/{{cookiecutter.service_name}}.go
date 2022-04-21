package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "nekowindow-backend/app/service/{{cookiecutter.department}}/{{cookiecutter.service_name}}/api/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// {{cookiecutter.serviceUpper}}Service is a greeter service.
type {{cookiecutter.serviceUpper}}Service struct {
	v1.Unimplemented{{cookiecutter.serviceUpper}}Server

	uc  *biz.{{cookiecutter.serviceUpper}}Usecase
	log *log.Helper
}

func New{{cookiecutter.serviceUpper}}Service(uc *biz.{{cookiecutter.serviceUpper}}Usecase, logger log.Logger) *{{cookiecutter.serviceUpper}}Service {
	return &{{cookiecutter.serviceUpper}}Service{uc: uc, log: log.NewHelper(log.With(logger, "package", "service"))}
}

// {{cookiecutter.serviceUpper}}Service impl example
func (s *{{cookiecutter.serviceUpper}}Service) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user, err := s.uc.CreateUser(ctx, req.Email, req.Password, req.Code)
	if err != nil {
		return nil, err
	}

	uid, sessionKey, err := s.uc.EmailLogin(ctx, user.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.CreateUserResponse{
		Uid:        uid,
		SessionKey: sessionKey,
	}, nil
}