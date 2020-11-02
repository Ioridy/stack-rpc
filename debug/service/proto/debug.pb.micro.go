// Code generated by protoc-gen-stack. DO NOT EDIT.
// source: micro/stack-rpc/debug/service/proto/debug.proto

package debug

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/stack-labs/stack-rpc/client"
	server "github.com/stack-labs/stack-rpc/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Debug service

type DebugService interface {
	Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error)
}

type debugService struct {
	c    client.Client
	name string
}

func NewDebugService(name string, c client.Client) DebugService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "debug"
	}
	return &debugService{
		c:    c,
		name: name,
	}
}

func (c *debugService) Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Health", in)
	out := new(HealthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error) {
	req := c.c.NewRequest(c.name, "Debug.Log", &LogRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &debugServiceLog{stream}, nil
}

type Debug_LogService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Record, error)
}

type debugServiceLog struct {
	stream client.Stream
}

func (x *debugServiceLog) Close() error {
	return x.stream.Close()
}

func (x *debugServiceLog) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugServiceLog) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugServiceLog) Recv() (*Record, error) {
	m := new(Record)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Debug service

type DebugHandler interface {
	Health(context.Context, *HealthRequest, *HealthResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
	Log(context.Context, *LogRequest, Debug_LogStream) error
}

func RegisterDebugHandler(s server.Server, hdlr DebugHandler, opts ...server.HandlerOption) error {
	type debug interface {
		Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error
		Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error
		Log(ctx context.Context, stream server.Stream) error
	}
	type Debug struct {
		debug
	}
	h := &debugHandler{hdlr}
	return s.Handle(s.NewHandler(&Debug{h}, opts...))
}

type debugHandler struct {
	DebugHandler
}

func (h *debugHandler) Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error {
	return h.DebugHandler.Health(ctx, in, out)
}

func (h *debugHandler) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.DebugHandler.Stats(ctx, in, out)
}

func (h *debugHandler) Log(ctx context.Context, stream server.Stream) error {
	m := new(LogRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Log(ctx, m, &debugLogStream{stream})
}

type Debug_LogStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Record) error
}

type debugLogStream struct {
	stream server.Stream
}

func (x *debugLogStream) Close() error {
	return x.stream.Close()
}

func (x *debugLogStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugLogStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugLogStream) Send(m *Record) error {
	return x.stream.Send(m)
}
