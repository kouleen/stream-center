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

func QueryClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*modle.StreamClientLine, total int64, err error) {
	query := getClientLineQuery(ctx, req)
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
	query = query.Order("io_type desc,create_time desc,id desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func getClientLineQuery(ctx context.Context, req *stream.StreamClientLineRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLine{}).Where("is_delete = 0")
	if req.GetStreamId() != "" {
		query = query.Where("stream_id = ?", req.GetStreamId())
	}
	if req.GetClientId() != "" {
		query = query.Where("client_id = ?", req.GetClientId())
	}
	if req.IoType != nil {
		query = query.Where("io_type = ?", req.IoType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	return query
}

func QueryClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*modle.StreamClientLine, err error) {
	query := getClientLineQuery(ctx, req)
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

func QueryClientLineById(ctx context.Context, id int64) (resp *modle.StreamClientLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(resp).Take(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func QueryClientLineByIdList(ctx context.Context, ids []int64) (resp []*modle.StreamClientLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLine{}).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

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

// QueryClientLineByStreamIdList 查询是否还剩有正常在线
func QueryClientLineByStreamIdList(ctx context.Context, streamIdList []string) (resp []*modle.StreamClientLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLine{}).Where("stream_id in (?) and status = 1", streamIdList).Find(&resp).Error; err != nil {
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
