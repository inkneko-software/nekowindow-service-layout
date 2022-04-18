package data

import (
	"context"
	v1 "nekowindow-backend/app/service/main/identify/api/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type identifyRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewIdentifyRepo(data *Data, logger log.Logger) biz.IdentifyRepo {
	return &identifyRepoImpl{
		data: data,
		log:  log.NewHelper(log.With(logger, "package", "data")),
	}
}

func (r *identifyRepoImpl) CreateSession(ctx context.Context, uid uint32, create int64, expire int64) (sessionKey string, err error) {
	resp, err := r.data.ic.CreateSession(ctx, &v1.CreateSessionRequest{Uid: uid, Create: create, Expire: expire})

	if err != nil {
		r.log.Error("调用identify.CreateSession失败")
		return "", nil
	}

	r.log.Info("调用成功， key=", resp.SessionKey)

	return resp.SessionKey, nil
}
