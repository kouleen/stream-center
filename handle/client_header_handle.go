package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
)

func QueryClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderPageResponse, err error) {
	return
}

func QueryClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*stream.StreamClientHeaderResponse, err error) {
	return
}

func QueryClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderResponse, err error) {
	return
}

func DeleteClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	return
}
