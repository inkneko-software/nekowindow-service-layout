package data

import (
	"context"
	"errors"
	"fmt"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type {{cookiecutter.service_name}}RepoImpl struct {
	data *Data
	log  *log.Helper
}

// New{{cookiecutter.serviceUpper}}Repo .
func New{{cookiecutter.serviceUpper}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.serviceUpper}}Repo {
	return &{{cookiecutter.service_name}}RepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{cookiecutter.service_name}}RepoImpl) CreateUser(ctx context.Context, user *biz.User) error {
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