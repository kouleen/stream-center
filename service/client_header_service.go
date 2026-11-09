package service

import (
	"context"

	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/stream"
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

func DeleteClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	itemList, err := repository.QueryClientHeaderByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return false, err
	}
	isDelete := int8(1)
	for _, item := range itemList {
		item.IsDelete = &isDelete
		item.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateClientHeader(ctx, itemList); err != nil {
		return false, err
	}
	return true, nil
}
