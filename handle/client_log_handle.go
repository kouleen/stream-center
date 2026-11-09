package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
)

func QueryClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogPageResponse, err error) {
	return
}

func QueryClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*stream.StreamClientLogResponse, err error) {
	return
}

func QueryClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogResponse, err error) {
	return
}

func DeleteClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp bool, err error) {
	return
}
