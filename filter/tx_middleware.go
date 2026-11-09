package filter

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/stream-center/txctx"
	"gorm.io/gorm"
)

func TxServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			txCtx := txctx.InjectTx(ctx, tx)
			return next(txCtx, req, resp)
		})
	}
}
