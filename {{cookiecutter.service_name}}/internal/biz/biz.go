package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(New{{cookiecutter.serviceUpper}}Usecase)

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
