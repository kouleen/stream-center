package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/service"
)

func QueryClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogPageResponse, err error) {
	return service.QueryClientLogPage(ctx, req)
}

func QueryClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*stream.StreamClientLogResponse, err error) {
	return service.QueryClientLogList(ctx, req)
}

func QueryClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogResponse, err error) {
	return service.QueryClientLog(ctx, req)
}

func DeleteClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp bool, err error) {
	return service.DeleteClientLog(ctx, req)
}
