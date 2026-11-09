package service

import (
	"context"
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/repository"
	"gorm.io/gorm"
)

func SyncStream(ctx context.Context, req *stream.StreamSyncRequest) (resp bool, err error) {
	server, err := repository.QueryServerByServerId(ctx, req.ServerId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	now := time.Now()
	if server == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			return false, err
		}
		serverId := node.Generate().Int64()
		server = &modle.StreamServer{
			ID:        serverId,
			ServerId:  req.ServerId,
			ServiceId: req.ServiceId,
		}
		if err = repository.CreateServer(ctx, server); err != nil {
			return false, err
		}
	} else {
		if req.ServiceId != server.ServiceId {
			server.Remark = server.ServiceId + " 已下线"
			server.ServiceId = req.ServiceId

			server.UpdateTime = &now
			if err = repository.UpdateServer(ctx, server); err != nil {
				return false, err
			}
		}
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	switch req.GetAction() {
	case "on_publish", "on_play":
		clientHeader, err := repository.QueryClientHeaderByStreamId(ctx, req.GetStreamId())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		if clientHeader == nil {
			clientHeaderId := node.Generate().Int64()
			clientHeader = &modle.StreamClientHeader{
				ID:         clientHeaderId,
				StreamId:   req.StreamId,
				StreamName: req.Stream,
				StreamUrl:  req.StreamUrl,
				ServerId:   req.ServerId,
				ServiceId:  req.ServiceId,
				App:        req.App,
				Vhost:      req.Vhost,
			}
			if err = repository.CreateClientHeader(ctx, clientHeader); err != nil {
				return false, err
			}
		}
		clientLineId := node.Generate().Int64()
		clientLine := &modle.StreamClientLine{
			ID:       clientLineId,
			ClientId: req.ClientId,
			StreamId: req.StreamId,
		}
		if err = repository.CreateClientLine(ctx, clientLine); err != nil {
			return false, err
		}
		clientLogId := node.Generate().Int64()
		clientLog := &modle.StreamClientLog{
			ID:       clientLogId,
			ClientId: req.GetClientId(),
			ClientIp: req.GetIp(),
			Action:   req.GetAction(),
		}
		if err = repository.CreateClientLog(ctx, clientLog); err != nil {
			return false, err
		}
	case "on_unpublish", "on_stop":
		status := int8(0)
		clientLine, err := repository.QueryClientLineByClientId(ctx, req.ClientId)
		if err != nil {
			return false, err
		}
		clientLine.Status = &status
		clientLine.UpdateTime = &now
		if err = repository.UpdateClientLine(ctx, clientLine); err != nil {
			return false, err
		}
		clientLogId := node.Generate().Int64()
		clientLog := &modle.StreamClientLog{
			ID:       clientLogId,
			ClientId: req.GetClientId(),
			ClientIp: req.GetIp(),
			Action:   req.GetAction(),
		}
		if err = repository.CreateClientLog(ctx, clientLog); err != nil {
			return false, err
		}
		clientLineList, err := repository.QueryClientLineByStreamId(ctx, req.GetStreamId())
		if err != nil {
			return false, err
		}
		// 只有一个并且这一个是本次
		if len(clientLineList) == 0 || (len(clientLineList) == 1 && clientLineList[0].ClientId == req.GetClientId()) {
			clientHeader, err := repository.QueryClientHeaderByStreamId(ctx, req.GetStreamId())
			if err != nil {
				return false, err
			}
			clientHeader.Status = &status
			clientHeader.UpdateTime = &now
			if err = repository.UpdateClientHeader(ctx, clientHeader); err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func QueryServerPage(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerPageResponse, err error) {
	list, total, err := repository.QueryServerPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamServerResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &stream.StreamServerPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryServerList(ctx context.Context, req *stream.StreamServerRequest) (resp []*stream.StreamServerResponse, err error) {
	list, err := repository.QueryServerList(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamServerResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryServer(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerResponse, err error) {
	item, err := repository.QueryServerById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), err
}

func DeleteServer(ctx context.Context, req *stream.StreamServerRequest) (resp bool, err error) {
	itemList, err := repository.QueryServerByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return false, err
	}
	isDelete := int8(1)
	for _, item := range itemList {
		item.IsDelete = &isDelete
		item.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateServer(ctx, itemList); err != nil {
		return false, err
	}
	return true, nil
}
