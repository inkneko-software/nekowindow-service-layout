package biz

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
	v1 "nekowindow-backend/api/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/v1"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

/**
#用户信息
CREATE TABLE user(
    uid BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户uid',
    register_date TIMESTAMP NOT NULL COMMENT '注册日期',
    username VARCHAR(20) NOT NULL COMMENT '用户名',
    auth_salt VARCHAR(32)  NOT NULL COMMENT '盐',
    auth_hash CHAR(40) NOT NULL COMMENT 'sha1(auth_salt + password)',
    nick VARCHAR(20) NOT NULL DEFAULT '' COMMENT '昵称',
    sign VARCHAR(50) NOT NULL DEFAULT '' COMMENT '个性签名',
    email VARCHAR(50)  COMMENT '邮箱',
    phone CHAR(11) COMMENT '手机号，不带区号',
    exp INT NOT NULL DEFAULT 0 COMMENT '经验',
    avatar VARCHAR(100)  COMMENT '头像地址',
    followers BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    subscribes INT NOT NULL DEFAULT 0 COMMENT '关注数量'
)ENGINE=InnoDB DEFAULT CHARSET utf8mb4;
*/
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


type AccountRepo interface {
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

	//保存邮箱注册验证码
	SaveEmailRegisterCode(ctx context.Context, code *RegisterCode) error

	//查询最新的邮件验证码，无记录返回nil, nil
	GetEmailRegisterCode(ctx context.Context, email string) (*RegisterCode, error)
}

type IdentifyRepo interface {
	//创建会话
	CreateSession(ctx context.Context, uid int64, create int64, expire int64) (sessionKey string, err error)
}

type AccountUsecase struct {
	ar  AccountRepo
	ir  IdentifyRepo
	log *log.Helper
}

func NewAccountUsecase(ar AccountRepo, ir IdentifyRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{ar: ar, ir: ir, log: log.NewHelper(log.With(logger, "package", "biz"))}
}

var (
	vaildEmail = regexp.MustCompile(`^[0-9a-zA-Z_-]+?@(qq\.com|vip\.qq\.com|163\.com|gmail\.com|foxmail\.com|126\.com|sina\.com)$`)
	vaildPass  = regexp.MustCompile("^[0-9a-zA-Z_-]+$")
)

func (uc *AccountUsecase) CreateUser(ctx context.Context, email string, password string, code string) (user *User, err error) {
	//邮箱检查
	if vaildEmail.MatchString(email) != true || utf8.RuneCountInString(email) < 5 || utf8.RuneCountInString(email) > 40 {
		return nil, v1.AccountEmailFormatError
	}

	//密码长度在8-16字符之间，
	if vaildPass.MatchString(password) != true || utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 16 {
		return nil, v1.AccountPasswordFormatError
	}

	//检查邮箱是否注册
	user, err = uc.ar.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, v1.AccountOperationFailed
	}

	if user != nil {
		return nil, v1.AccountEmailExist
	}

	//检查验证码是否匹配
	lastCode, err := uc.ar.GetEmailRegisterCode(ctx, email)
	if err != nil {
		return nil, v1.AccountOperationFailed
	}

	if lastCode == nil || time.Now().Unix()-lastCode.CreateTimestamp.Unix() > 60 {
		return nil, v1.AccountEmailRegisterCodeError
	}

	//生成新用户
	authSalt := fmt.Sprintf("%x", md5.Sum([]byte(uuid.NewString())))
	authHashByte := sha1.Sum([]byte(authSalt + password))
	authHash := fmt.Sprintf("%x", authHashByte)

	user = &User{RegisterDate: time.Now(), AuthSalt: authSalt, AuthHash: authHash, Email: email, Avatar: "/picture/avatar/default.jpg"}
	err = uc.ar.CreateUser(ctx, user)
	if err != nil {
		return nil, v1.AccountOperationFailed
	}

	return user, nil
}