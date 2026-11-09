package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/service"
)

func QueryClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLinePageResponse, err error) {
	return service.QueryClientLinePage(ctx, req)
}

func QueryClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*stream.StreamClientLineResponse, err error) {
	return service.QueryClientLineList(ctx, req)
}

func QueryClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLineResponse, err error) {
	return service.QueryClientLine(ctx, req)
}

func DeleteClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	return service.DeleteClientLine(ctx, req)
}
