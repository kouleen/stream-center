package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/repository"
)

func QueryClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogPageResponse, err error) {
	list, total, err := repository.QueryClientLogPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*stream.StreamClientLogResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &stream.StreamClientLogPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*stream.StreamClientLogResponse, err error) {
	list, err := repository.QueryClientLogList(ctx, req)
	if err != nil {
		return
	}
	records := make([]*stream.StreamClientLogResponse, len(list))
	for index, lis := range list {
		records[index] = lis.ConvertResp()
	}
	return records, nil
}

func QueryClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogResponse, err error) {
	item, err := repository.QueryClientLogById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), nil
}

func DeleteClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp bool, err error) {
	if err = repository.BatchDeleteClientLog(ctx, req.GetIdList()); err != nil {
		return
	}
	return true, nil
}
