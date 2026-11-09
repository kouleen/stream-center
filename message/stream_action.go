package message

import (
	"context"
	"encoding/json"

	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/logutil"
	"github.com/kouleen/idl/kitex_gen/stream"
	"github.com/kouleen/stream-center/service"
)

type StreamAction struct{}

func (p *StreamAction) Queue() string {
	return "stream.sync.action.queue"
}

func (p *StreamAction) HandleMessage(ctx context.Context, msg *message.Message) error {
	defer func() {
		if err := recover(); err != nil {
			logutil.CtxError(ctx, "recover panic err: %#v", err)
		}
	}()
	var req stream.StreamSyncRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	if _, err := service.SyncStream(ctx, &req); err != nil {
		return err
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(StreamAction))
}
