package data

import (
	"context"
	"errors"
	"fmt"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type accountRepoImpl struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepoImpl) CreateUser(ctx context.Context, user *biz.User) error {
	err := r.data.mysql.Transaction(func(tx *gorm.DB) error {
		if err := r.data.mysql.Create(user).Error; err != nil {
			log.Info("创建用户失败")
			tx.Rollback()
			return err
		}
		user.Username = fmt.Sprintf("neko_%d", user.ID)
		if err := r.data.mysql.Save(user).Error; err != nil {
			log.Info("创建用户失败, 更新用户名失败")
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *accountRepoImpl) UpdateUser(ctx context.Context, user *biz.User) error {
	if err := r.data.mysql.Save(user).Error; err != nil {
		log.Info("更新用户失败")
		return err
	}
	return nil
}

//根据uid获取用户
func (r *accountRepoImpl) GetUser(ctx context.Context, uid int64) (user *biz.User, err error) {
	if err = r.data.mysql.Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	return user, err
}

//根据邮箱查询用户
func (r *accountRepoImpl) GetUserByEmail(ctx context.Context, email string) (user *biz.User, err error) {
	if err = r.data.mysql.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, err
}

//根据手机号查询用户
func (r *accountRepoImpl) GetUserByPhone(ctx context.Context, phone string) (user *biz.User, err error) {
	if err = r.data.mysql.Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, err
}

//保存邮箱注册验证码
func (r *accountRepoImpl) SaveEmailRegisterCode(ctx context.Context, code *biz.RegisterCode) error {
	if err := r.data.mysql.Create(code).Error; err != nil {
		return err
	}

	return nil
}

//查询最新的邮件验证码
func (r *accountRepoImpl) GetEmailRegisterCode(ctx context.Context, email string) (code *biz.RegisterCode, err error) {
	if err = r.data.mysql.Where("email = ?", email).Order("create_timestamp DESC").First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return code, err
}
