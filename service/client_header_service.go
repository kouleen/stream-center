package service

import (
	"context"
	"time"

	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/repository"
)

func QueryClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderPageResponse, err error) {
	list, total, err := repository.QueryClientHeaderPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamClientHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &stream.StreamClientHeaderPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*stream.StreamClientHeaderResponse, err error) {
	list, err := repository.QueryClientHeaderList(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamClientHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderResponse, err error) {
	item, err := repository.QueryClientHeaderById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), err
}

func StopClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	itemList, err := repository.QueryClientHeaderByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return false, err
	}
	clientStreamIdList := make([]string, len(itemList))
	for index, item := range itemList {
		clientStreamIdList[index] = item.StreamId
	}
	lineList, err := repository.QueryClientLineByStreamIdList(ctx, clientStreamIdList)
	if err != nil {
		return false, err
	}
	traceId := ctxutil.GetTraceId(ctx)
	userId := ctxutil.GetUserId(ctx)
	bgCtx := context.Background()
	bgCtx, bgCancel := context.WithTimeout(bgCtx, 60*time.Second)
	bgCtx = ctxutil.SetMeta(bgCtx, traceId, userId)
	lineCh := make(chan *modle.StreamClientLine, 1)
	go sendKickStreamClient(bgCtx, lineCh)
	go func() {
		defer func() {
			close(lineCh)
			bgCancel()
		}()
		for _, line := range lineList {
			select {
			case lineCh <- line:
			case <-bgCtx.Done():
				return
			}
		}
	}()
	return true, nil
}
