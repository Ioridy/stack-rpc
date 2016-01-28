package handler

import (
	log "github.com/golang/glog"
	example "github.com/micro/go-micro/examples/server/proto/example"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("Received Example.Call request with metadata: %v", md)
	rsp.Msg = server.DefaultOptions().Id + ": Hello " + req.Name
	return nil
}

func (e *Example) Stream(ctx context.Context, stream server.Streamer) error {
	log.Info("Executing streaming handler")
	req := &example.StreamingRequest{}

	// We just want to receive 1 request and then process here
	if err := stream.Recv(req); err != nil {
		log.Errorf("Error receiving streaming request: %v", err)
		return err
	}

	log.Infof("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)

		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e *Example) PingPong(ctx context.Context, stream server.Streamer) error {
	for {
		req := &example.Ping{}
		if err := stream.Recv(req); err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
	return nil
}
