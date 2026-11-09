package repository

import (
	"context"

	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/txctx"
)

func CreateClientLog(ctx context.Context, entity *modle.StreamClientLog) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientLog{}).Create(entity).Error
}
