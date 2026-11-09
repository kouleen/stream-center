package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/service"
)

func SyncStream(ctx context.Context, req *stream.StreamSyncRequest) (resp bool, err error) {
	return service.SyncStream(ctx, req)
}

func QueryServerPage(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerPageResponse, err error) {
	return
}

func QueryServerList(ctx context.Context, req *stream.StreamServerRequest) (resp []*stream.StreamServerResponse, err error) {
	return
}

func QueryServer(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerResponse, err error) {
	return
}

func DeleteServer(ctx context.Context, req *stream.StreamServerRequest) (resp bool, err error) {
	return
}
