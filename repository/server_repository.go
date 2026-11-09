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

func QueryServerPage(ctx context.Context, req *stream.StreamServerRequest) (resp []*modle.StreamServer, total int64, err error) {
	query := getServerQuery(ctx, req)
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

func getServerQuery(ctx context.Context, req *stream.StreamServerRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamServer{}).Where("is_delete = 0")
	if req.GetServerId() != "" {
		query.Where("server_id = ?", req.GetServerId())
	}
	if req.GetServiceId() != "" {
		query.Where("service_id = ?", req.GetServiceId())
	}
	if req.Status != nil {
		query.Where("status = ?", req.Status)
	}
	return query
}

func QueryServerList(ctx context.Context, req *stream.StreamServerRequest) (resp []*modle.StreamServer, err error) {
	query := getServerQuery(ctx, req)
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

func QueryServerByServerId(ctx context.Context, serverId string) (resp *modle.StreamServer, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(resp).Where("server_id = ?", serverId).Take(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func QueryServerById(ctx context.Context, id int64) (resp *modle.StreamServer, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(resp).Take(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func QueryServerByIdList(ctx context.Context, idList []int64) (resp []*modle.StreamServer, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.StreamServer{}).Where("id in (?)", idList).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func CreateServer(ctx context.Context, entity *modle.StreamServer) error {
	return txctx.ExtractTx(ctx).Create(entity).Error
}

func UpdateServer(ctx context.Context, entity *modle.StreamServer) error {
	return txctx.ExtractTx(ctx).Save(entity).Error
}
