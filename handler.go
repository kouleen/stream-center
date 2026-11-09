package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/handle"
)

// StreamServiceImpl implements the last service interface defined in the IDL.
type StreamServiceImpl struct{}

// SyncStream implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) SyncStream(ctx context.Context, req *stream.StreamSyncRequest) (resp bool, err error) {
	return handle.SyncStream(ctx, req)
}

// QueryStreamServerPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServerPage(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerPageResponse, err error) {
	return handle.QueryServerPage(ctx, req)
}

// QueryStreamServerList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServerList(ctx context.Context, req *stream.StreamServerRequest) (resp []*stream.StreamServerResponse, err error) {
	return handle.QueryServerList(ctx, req)
}

// QueryStreamServer implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServer(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerResponse, err error) {
	return handle.QueryServer(ctx, req)
}

// DeleteStreamServer implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamServer(ctx context.Context, req *stream.StreamServerRequest) (resp bool, err error) {
	return handle.DeleteServer(ctx, req)
}

// QueryStreamClientHeaderPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderPageResponse, err error) {
	return handle.QueryClientHeaderPage(ctx, req)
}

// QueryStreamClientHeaderList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*stream.StreamClientHeaderResponse, err error) {
	return handle.QueryClientHeaderList(ctx, req)
}

// QueryStreamClientHeader implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderResponse, err error) {
	return handle.QueryClientHeader(ctx, req)
}

// DeleteStreamClientHeader implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	return handle.DeleteClientHeader(ctx, req)
}

// QueryStreamClientLinePage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLinePageResponse, err error) {
	return handle.QueryClientLinePage(ctx, req)
}

// QueryStreamClientLineList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*stream.StreamClientLineResponse, err error) {
	return handle.QueryClientLineList(ctx, req)
}

// QueryStreamClientLine implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLineResponse, err error) {
	return handle.QueryClientLine(ctx, req)
}

// DeleteStreamClientLine implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	return handle.DeleteClientLine(ctx, req)
}

// QueryStreamClientLogPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogPageResponse, err error) {
	return handle.QueryClientLogPage(ctx, req)
}

// QueryStreamClientLogList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*stream.StreamClientLogResponse, err error) {
	return handle.QueryClientLogList(ctx, req)
}

// QueryStreamClientLog implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogResponse, err error) {
	return handle.QueryClientLog(ctx, req)
}

// DeleteStreamClientLog implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp bool, err error) {
	return handle.DeleteClientLog(ctx, req)
}
