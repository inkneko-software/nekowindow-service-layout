package data

import (
	"context"
	identifyv1 "nekowindow-backend/app/service/main/identify/api/v1"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, New{{cookiecutter.serviceUpper}}Repo, NewIdentifyRepo, NewIdentifyClient, NewTransaction)

// Data .
type Data struct {
	// TODO wrapped database client
	mysql *gorm.DB
	ic    identifyv1.IdentifyClient
}

// NewData .
func NewData(
	c *conf.Data,
	ic identifyv1.IdentifyClient,
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	mysqldb, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Data{mysql: mysqldb, ic: ic}, cleanup, nil
}

func NewIdentifyClient(r registry.Discovery) identifyv1.IdentifyClient {

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://default/nekowindow.service.main.identify"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := identifyv1.NewIdentifyClient(conn)
	return c
}
