package service

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/google/uuid"
	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/common/pkg/logutil"
	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/modle"
	"github.com/kouleen/stream-center/repository"
)

func QueryClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLinePageResponse, err error) {
	list, total, err := repository.QueryClientLinePage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamClientLineResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &stream.StreamClientLinePageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*stream.StreamClientLineResponse, err error) {
	list, err := repository.QueryClientLineList(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamClientLineResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLineResponse, err error) {
	item, err := repository.QueryClientLineById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), err
}

func StopClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	itemList, err := repository.QueryClientLineByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return false, nil
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
		for _, line := range itemList {
			select {
			case lineCh <- line:
			case <-bgCtx.Done():
				return
			}
		}
	}()
	return true, nil
}

func sendKickStreamClient(ctx context.Context, ch <-chan *modle.StreamClientLine) {
	defer func() {
		if err := recover(); err != nil {
			logutil.CtxError(ctx, "send kick stream client error.err:%v", err)
		}
	}()
	for entity := range ch {
		messageByte, err := json.Marshal(entity)
		if err != nil {
			panic(err)
		}
		if msgErr := message.SendToQueue(context.Background(), "stream-client-kick-topic", &message.Message{
			TraceID:     ctxutil.GetTraceId(ctx),
			MessageID:   strconv.Itoa(int(uuid.New().ID())),
			ContentType: "application/json",
			Body:        messageByte,
			Headers:     make(map[string]any),
		}); msgErr != nil {
			logger.CtxErrorf(ctx, "send login log error: %v", err)
		}
	}
}
