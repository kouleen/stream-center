package service

import (
	"context"
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
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
		clientHeaderId := node.Generate().Int64()
		clientHeader := &modle.StreamClientHeader{
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
		if len(clientLineList) == 1 && clientLineList[0].ClientId == req.GetClientId() {
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
	return
}
