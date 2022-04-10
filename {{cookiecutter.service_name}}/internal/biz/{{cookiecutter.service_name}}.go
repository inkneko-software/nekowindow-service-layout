package biz

import (
	"context"
	v1 "nekowindow-backend/app/service/{{cookiecutter.department}}/{{cookiecutter.service_name}}/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)


type User struct {
	ID           int64 `gorm:"column:uid" gorm:"primaryKey"`
	RegisterDate time.Time
	Username     string
	AuthSalt     string
	AuthHash     string
	Nick         string
	Sign         string
	Email        string
	Phone        string
	Exp          int
	Avatar       string
	Followers    int64
	Subscribes   int64
}


type {{cookiecutter.serviceUpper}}Repo interface {
	//创建用户
	CreateUser(ctx context.Context, user *User) error

	//根据uid获取用户, 无记录返回nil, nil
	GetUser(ctx context.Context, uid int64) (*User, error)

	//根据邮箱查询用户, 无记录返回nil, nil
	GetUserByEmail(ctx context.Context, email string) (*User, error)

	//根据手机号查询用户，无记录返回nil, nil
	GetUserByPhone(ctx context.Context, phone string) (*User, error)

	//根据uid更新用户信息
	UpdateUser(ctx context.Context, user *User) error
}

type IdentifyRepo interface {
	//创建会话
	CreateSession(ctx context.Context, uid uint32, create int64, expire int64) (sessionKey string, err error)
}

type {{cookiecutter.serviceUpper}}Usecase struct {
	ar  {{cookiecutter.serviceUpper}}Repo
	ir  IdentifyRepo
	tm  Transaction
	log *log.Helper
}

func New{{cookiecutter.serviceUpper}}Usecase(ar {{cookiecutter.serviceUpper}}Repo, ir IdentifyRepo, tm  Transaction, logger log.Logger) *{{cookiecutter.serviceUpper}}Usecase {
	return &{{cookiecutter.serviceUpper}}Usecase{ar: ar, ir: ir, tm:tm, log: log.NewHelper(log.With(logger, "package", "biz"))}
}


func (uc *{{cookiecutter.serviceUpper}}Usecase) CreateUser(ctx context.Context, email string, password string, code string) (user *User, err error) {
	// 在此处实现业务逻辑
	return user, nil
}