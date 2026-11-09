package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/txctx"
	"github.com/kouleen/stream-center/utils"
	"gorm.io/gorm"
)

func QueryClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*modle.StreamClientHeader, total int64, err error) {
	query := getClientHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc,id desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func getClientHeaderQuery(ctx context.Context, req *stream.StreamClientHeaderRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientHeader{}).Where("is_delete = 0")
	if req.GetStreamId() != "" {
		query = query.Where("stream_id = ?", req.GetStreamId())
	}
	if req.GetStreamName() != "" {
		query = query.Where("stream_name = ?", req.GetStreamName())
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	if req.GetServerId() != "" {
		query = query.Where("server_id = ?", req.GetServerId())
	}
	if req.GetServiceId() != "" {
		query = query.Where("service_id = ?", req.GetServiceId())
	}
	if req.GetApp() != "" {
		query = query.Where("app = ?", req.GetApp())
	}
	if req.GetVhost() != "" {
		query = query.Where("vhost = ?", req.GetVhost())
	}
	return query
}

func QueryClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*modle.StreamClientHeader, err error) {
	query := getClientHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	query = query.Order("create_time desc")
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryClientHeaderById(ctx context.Context, id int64) (resp *modle.StreamClientHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Take(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func QueryClientHeaderByIdList(ctx context.Context, ids []int64) (resp []*modle.StreamClientHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientHeader{}).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

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

func BatchUpdateClientHeader(ctx context.Context, entityList []*modle.StreamClientHeader) error {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, entity := range entityList {
			if err := tx.Model(&modle.StreamClientHeader{}).Where("id = ?", entity.ID).Save(entity).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
