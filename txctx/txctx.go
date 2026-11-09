package txctx

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"gorm.io/gorm"
)

type txKey struct{}

// InjectTx 把tx放进ctx
func InjectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// ExtractTx 优先取ctx里的事务tx，没有就返回传入的全局db
func ExtractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return mysql.GetWriteMysqlDDB().WithContext(ctx)
}
