package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/stream"
)

type StreamClientLine struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	ClientId   string     `json:"clientId" gorm:"column:client_id;not null"`
	StreamId   string     `json:"streamId" gorm:"column:stream_id;not null"`
	IoType     *int8      `json:"ioType" gorm:"column:io_type;not null;"`
	Status     *int8      `json:"status" gorm:"column:status;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *StreamClientLine) TableName() string {
	return "stream_client_line"
}

func (p *StreamClientLine) ConvertResp() *stream.StreamClientLineResponse {
	return &stream.StreamClientLineResponse{
		Id:         p.ID,
		ClientId:   p.ClientId,
		StreamId:   p.StreamId,
		IoType:     p.IoType,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}

var ioTypeMap = map[string]int8{
	"on_play":    int8(1),
	"on_publish": int8(2),
}

func (p *StreamClientLine) ToIoType(action string) *int8 {
	if ioType, ok := ioTypeMap[action]; ok {
		return &ioType
	}
	return nil
}
