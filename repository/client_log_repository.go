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

func QueryClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*modle.StreamClientLog, total int64, err error) {
	query := getClientLogQuery(ctx, req)
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

func getClientLogQuery(ctx context.Context, req *stream.StreamClientLogRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLog{}).Where("is_delete = 0")
	if req.GetClientId() != "" {
		query = query.Where("client_id = ?", req.GetClientId())
	}
	if req.GetAction() != "" {
		query = query.Where("action = ?", req.GetAction())
	}
	if req.GetClientIp() != "" {
		query = query.Where("client_ip like ?", "%"+req.GetClientIp()+"%")
	}
	return query
}

func QueryClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*modle.StreamClientLog, err error) {
	query := getClientLogQuery(ctx, req)
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

func QueryClientLogById(ctx context.Context, id int64) (resp *modle.StreamClientLog, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamClientLog{}).Take(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func CreateClientLog(ctx context.Context, entity *modle.StreamClientLog) error {
	return txctx.ExtractTx(ctx).Model(&modle.StreamClientLog{}).Create(entity).Error
}

func BatchDeleteClientLog(ctx context.Context, idList []int64) error {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(idList) > 0 {
			if err := tx.Where("id in (?)", idList).Delete(&modle.StreamClientLog{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
