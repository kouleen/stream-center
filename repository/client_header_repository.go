package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/txctx"
)

func QueryClientHeaderByStreamId(ctx context.Context, streamId string) (resp *modle.StreamClientHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientHeader{}).Where("stream_id = ?", streamId).Take(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateClientHeader(ctx context.Context, entity *modle.StreamClientHeader) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientHeader{}).Create(entity).Error
}

func UpdateClientHeader(ctx context.Context, entity *modle.StreamClientHeader) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientHeader{}).Where("id = ?", entity.ID).Save(entity).Error
}
