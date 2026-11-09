package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/txctx"
)

func QueryClientLineByClientId(ctx context.Context, clientId string) (resp *modle.StreamClientLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLine{}).Where("client_id = ?", clientId).Take(&resp).Error; err != nil {
		return nil, err
	}
	return
}

// QueryClientLineByStreamId 查询是否还剩有正常在线
func QueryClientLineByStreamId(ctx context.Context, streamId string) (resp []*modle.StreamClientLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLine{}).Where("stream_id = ? and status = 1", streamId).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func CreateClientLine(ctx context.Context, entity *modle.StreamClientLine) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientLine{}).Create(entity).Error
}

func UpdateClientLine(ctx context.Context, entity *modle.StreamClientLine) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientLine{}).Where("id = ?", entity.ID).Save(entity).Error
}
