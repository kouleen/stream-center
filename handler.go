package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/stream"
)

// StreamServiceImpl implements the last service interface defined in the IDL.
type StreamServiceImpl struct{}

// SyncStream implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) SyncStream(ctx context.Context, req *stream.StreamSyncRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamServerPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServerPage(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamServerList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServerList(ctx context.Context, req *stream.StreamServerRequest) (resp []*stream.StreamServerResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamServer implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamServer(ctx context.Context, req *stream.StreamServerRequest) (resp *stream.StreamServerResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStreamServer implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamServer(ctx context.Context, req *stream.StreamServerRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientHeaderPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeaderPage(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientHeaderList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeaderList(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp []*stream.StreamClientHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientHeader implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp *stream.StreamClientHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStreamClientHeader implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientHeader(ctx context.Context, req *stream.StreamClientHeaderRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLinePage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLinePage(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLinePageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLineList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLineList(ctx context.Context, req *stream.StreamClientLineRequest) (resp []*stream.StreamClientLineResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLine implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp *stream.StreamClientLineResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStreamClientLine implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientLine(ctx context.Context, req *stream.StreamClientLineRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLogPage implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLogPage(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLogList implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLogList(ctx context.Context, req *stream.StreamClientLogRequest) (resp []*stream.StreamClientLogResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStreamClientLog implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) QueryStreamClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp *stream.StreamClientLogResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStreamClientLog implements the StreamServiceImpl interface.
func (s *StreamServiceImpl) DeleteStreamClientLog(ctx context.Context, req *stream.StreamClientLogRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
