package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/stream"
)

type StreamServer struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	ServerId   string     `json:"serverId" gorm:"column:server_id;not null"`
	ServiceId  string     `json:"serviceId" gorm:"column:service_id;not null"`
	Status     *int8      `json:"status" gorm:"column:status;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (*StreamServer) TableName() string {
	return "stream_server"
}

func (p *StreamServer) ConvertResp() *stream.StreamServerResponse {
	return &stream.StreamServerResponse{
		Id:         p.ID,
		ServerId:   p.ServerId,
		ServiceId:  p.ServiceId,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
