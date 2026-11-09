package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
)

func QueryClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLinePageResponse, err error) {
	return
}

func QueryClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*stream.StreamClientLineResponse, err error) {
	return
}

func QueryClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLineResponse, err error) {
	return
}

func DeleteClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	return
}
