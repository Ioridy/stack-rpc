// Code generated by protoc-gen-stack. DO NOT EDIT.
// source: router.proto

package stack_rpc_router

import (
	fmt "fmt"
	math "math"

	context "context"

	proto "github.com/golang/protobuf/proto"

	api "github.com/stack-labs/stack-rpc/api"

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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Router service

func NewRouterEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Router service

type RouterService interface {
	Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Router_WatchService, error)
	Advertise(ctx context.Context, in *Request, opts ...client.CallOption) (Router_AdvertiseService, error)
	Solicit(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Process(ctx context.Context, in *Advert, opts ...client.CallOption) (*ProcessResponse, error)
	Status(ctx context.Context, in *Request, opts ...client.CallOption) (*StatusResponse, error)
}

type routerService struct {
	c    client.Client
	name string
}

func NewRouterService(name string, c client.Client) RouterService {
	return &routerService{
		c:    c,
		name: name,
	}
}

func (c *routerService) Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error) {
	req := c.c.NewRequest(c.name, "Router.Lookup", in)
	out := new(LookupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Router_WatchService, error) {
	req := c.c.NewRequest(c.name, "Router.Watch", &WatchRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &routerServiceWatch{stream}, nil
}

type Router_WatchService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Event, error)
}

type routerServiceWatch struct {
	stream client.Stream
}

func (x *routerServiceWatch) Close() error {
	return x.stream.Close()
}

func (x *routerServiceWatch) Context() context.Context {
	return x.stream.Context()
}

func (x *routerServiceWatch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routerServiceWatch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routerServiceWatch) Recv() (*Event, error) {
	m := new(Event)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerService) Advertise(ctx context.Context, in *Request, opts ...client.CallOption) (Router_AdvertiseService, error) {
	req := c.c.NewRequest(c.name, "Router.Advertise", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &routerServiceAdvertise{stream}, nil
}

type Router_AdvertiseService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Advert, error)
}

type routerServiceAdvertise struct {
	stream client.Stream
}

func (x *routerServiceAdvertise) Close() error {
	return x.stream.Close()
}

func (x *routerServiceAdvertise) Context() context.Context {
	return x.stream.Context()
}

func (x *routerServiceAdvertise) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routerServiceAdvertise) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routerServiceAdvertise) Recv() (*Advert, error) {
	m := new(Advert)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerService) Solicit(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Router.Solicit", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Process(ctx context.Context, in *Advert, opts ...client.CallOption) (*ProcessResponse, error) {
	req := c.c.NewRequest(c.name, "Router.Process", in)
	out := new(ProcessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerService) Status(ctx context.Context, in *Request, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "Router.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Router service

type RouterHandler interface {
	Lookup(context.Context, *LookupRequest, *LookupResponse) error
	Watch(context.Context, *WatchRequest, Router_WatchStream) error
	Advertise(context.Context, *Request, Router_AdvertiseStream) error
	Solicit(context.Context, *Request, *Response) error
	Process(context.Context, *Advert, *ProcessResponse) error
	Status(context.Context, *Request, *StatusResponse) error
}

func RegisterRouterHandler(s server.Server, hdlr RouterHandler, opts ...server.HandlerOption) error {
	type router interface {
		Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error
		Watch(ctx context.Context, stream server.Stream) error
		Advertise(ctx context.Context, stream server.Stream) error
		Solicit(ctx context.Context, in *Request, out *Response) error
		Process(ctx context.Context, in *Advert, out *ProcessResponse) error
		Status(ctx context.Context, in *Request, out *StatusResponse) error
	}
	type Router struct {
		router
	}
	h := &routerHandler{hdlr}
	return s.Handle(s.NewHandler(&Router{h}, opts...))
}

type routerHandler struct {
	RouterHandler
}

func (h *routerHandler) Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error {
	return h.RouterHandler.Lookup(ctx, in, out)
}

func (h *routerHandler) Watch(ctx context.Context, stream server.Stream) error {
	m := new(WatchRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RouterHandler.Watch(ctx, m, &routerWatchStream{stream})
}

type Router_WatchStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Event) error
}

type routerWatchStream struct {
	stream server.Stream
}

func (x *routerWatchStream) Close() error {
	return x.stream.Close()
}

func (x *routerWatchStream) Context() context.Context {
	return x.stream.Context()
}

func (x *routerWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routerWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routerWatchStream) Send(m *Event) error {
	return x.stream.Send(m)
}

func (h *routerHandler) Advertise(ctx context.Context, stream server.Stream) error {
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RouterHandler.Advertise(ctx, m, &routerAdvertiseStream{stream})
}

type Router_AdvertiseStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Advert) error
}

type routerAdvertiseStream struct {
	stream server.Stream
}

func (x *routerAdvertiseStream) Close() error {
	return x.stream.Close()
}

func (x *routerAdvertiseStream) Context() context.Context {
	return x.stream.Context()
}

func (x *routerAdvertiseStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routerAdvertiseStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routerAdvertiseStream) Send(m *Advert) error {
	return x.stream.Send(m)
}

func (h *routerHandler) Solicit(ctx context.Context, in *Request, out *Response) error {
	return h.RouterHandler.Solicit(ctx, in, out)
}

func (h *routerHandler) Process(ctx context.Context, in *Advert, out *ProcessResponse) error {
	return h.RouterHandler.Process(ctx, in, out)
}

func (h *routerHandler) Status(ctx context.Context, in *Request, out *StatusResponse) error {
	return h.RouterHandler.Status(ctx, in, out)
}

// Api Endpoints for Table service

func NewTableEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Table service

type TableService interface {
	Create(ctx context.Context, in *Route, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *Route, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *Route, opts ...client.CallOption) (*UpdateResponse, error)
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*ListResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type tableService struct {
	c    client.Client
	name string
}

func NewTableService(name string, c client.Client) TableService {
	return &tableService{
		c:    c,
		name: name,
	}
}

func (c *tableService) Create(ctx context.Context, in *Route, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Table.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Delete(ctx context.Context, in *Route, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Table.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Update(ctx context.Context, in *Route, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Table.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Table.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.name, "Table.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Table service

type TableHandler interface {
	Create(context.Context, *Route, *CreateResponse) error
	Delete(context.Context, *Route, *DeleteResponse) error
	Update(context.Context, *Route, *UpdateResponse) error
	List(context.Context, *Request, *ListResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
}

func RegisterTableHandler(s server.Server, hdlr TableHandler, opts ...server.HandlerOption) error {
	type table interface {
		Create(ctx context.Context, in *Route, out *CreateResponse) error
		Delete(ctx context.Context, in *Route, out *DeleteResponse) error
		Update(ctx context.Context, in *Route, out *UpdateResponse) error
		List(ctx context.Context, in *Request, out *ListResponse) error
		Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error
	}
	type Table struct {
		table
	}
	h := &tableHandler{hdlr}
	return s.Handle(s.NewHandler(&Table{h}, opts...))
}

type tableHandler struct {
	TableHandler
}

func (h *tableHandler) Create(ctx context.Context, in *Route, out *CreateResponse) error {
	return h.TableHandler.Create(ctx, in, out)
}

func (h *tableHandler) Delete(ctx context.Context, in *Route, out *DeleteResponse) error {
	return h.TableHandler.Delete(ctx, in, out)
}

func (h *tableHandler) Update(ctx context.Context, in *Route, out *UpdateResponse) error {
	return h.TableHandler.Update(ctx, in, out)
}

func (h *tableHandler) List(ctx context.Context, in *Request, out *ListResponse) error {
	return h.TableHandler.List(ctx, in, out)
}

func (h *tableHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.TableHandler.Query(ctx, in, out)
}