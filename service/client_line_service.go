package service

import (
	"context"

	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/stream"
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

func DeleteClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	itemList, err := repository.QueryClientLineByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return false, err
	}
	isDelete := int8(1)
	for _, item := range itemList {
		item.IsDelete = &isDelete
		item.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateClientLine(ctx, itemList); err != nil {
		return false, err
	}
	return true, nil
}
