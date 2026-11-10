package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/service"
)

func QueryClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderPageResponse, err error) {
	return service.QueryClientHeaderPage(ctx, req)
}

func QueryClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*stream.StreamClientHeaderResponse, err error) {
	return service.QueryClientHeaderList(ctx, req)
}

func QueryClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderResponse, err error) {
	return service.QueryClientHeader(ctx, req)
}

func StopClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	return service.StopClientHeader(ctx, req)
}
