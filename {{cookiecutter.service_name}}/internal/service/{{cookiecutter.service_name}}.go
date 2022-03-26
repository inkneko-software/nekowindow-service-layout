package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "nekowindow-backend/api/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// AccountService is a greeter service.
type AccountService struct {
	v1.UnimplementedAccountServer

	uc  *biz.AccountUsecase
	log *log.Helper
}

func NewAccountService(uc *biz.AccountUsecase, logger log.Logger) *AccountService {
	return &AccountService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AccountService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
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

func (s *AccountService) SendRegisterEmail(ctx context.Context, req *v1.SendRegisterEmailRequest) (*v1.SendRegisterEmailResponse, error) {
	s.log.Infof("req->%v", req)
	err := s.uc.SendRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &v1.SendRegisterEmailResponse{Delay: 60}, nil
}

func (s *AccountService) SendRecoveryEmail(ctx context.Context, req *v1.SendRecoveryEmailRequest) (*v1.SendRecoveryEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRecoveryEmail not implemented")
}
func (s *AccountService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*v1.UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (s *AccountService) EmailLogin(ctx context.Context, req *v1.EmailLoginRequest) (*v1.EmailLoginResponse, error) {
	s.log.Infof("")
	uid, sessionKey, err := s.uc.EmailLogin(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.EmailLoginResponse{
		Uid:        uid,
		SessionKey: sessionKey,
	}, nil

}

func (s *AccountService) ChangeNick(context.Context, *v1.ChangeNickRequest) (*v1.ChangeNickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNick not implemented")
}

func (s *AccountService) ChangeAvatar(context.Context, *v1.ChangeAvatarReq) (*v1.ChangeAvatarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAvatar not implemented")
}
func (s *AccountService) ChangeSign(context.Context, *v1.ChangeSignReq) (*v1.ChangeSignReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSign not implemented")
}
func (s *AccountService) ChangeUserInfo(context.Context, *v1.ChangeUserInfoReq) (*v1.ChangeUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserInfo not implemented")
}

func (s *AccountService) GetUserCard(ctx context.Context, req *v1.GetUserCardReq) (*v1.GetUserCardReply, error) {
	user, err := s.uc.GetUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserCardReply{
		Uid:    user.ID,
		Nick:   user.Nick,
		Sign:   user.Sign,
		Exp:    int32(user.Exp),
		Avatar: user.Avatar,
	}, nil
}
