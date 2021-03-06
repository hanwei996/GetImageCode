// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetImageCode.proto

package GetImageCode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for GetImageCode service

func NewGetImageCodeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetImageCode service

type GetImageCodeService interface {
	GetImageCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type getImageCodeService struct {
	c    client.Client
	name string
}

func NewGetImageCodeService(name string, c client.Client) GetImageCodeService {
	return &getImageCodeService{
		c:    c,
		name: name,
	}
}

func (c *getImageCodeService) GetImageCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImageCode.GetImageCode", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetImageCode service

type GetImageCodeHandler interface {
	GetImageCode(context.Context, *Request, *Response) error
}

func RegisterGetImageCodeHandler(s server.Server, hdlr GetImageCodeHandler, opts ...server.HandlerOption) error {
	type getImageCode interface {
		GetImageCode(ctx context.Context, in *Request, out *Response) error
	}
	type GetImageCode struct {
		getImageCode
	}
	h := &getImageCodeHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImageCode{h}, opts...))
}

type getImageCodeHandler struct {
	GetImageCodeHandler
}

func (h *getImageCodeHandler) GetImageCode(ctx context.Context, in *Request, out *Response) error {
	return h.GetImageCodeHandler.GetImageCode(ctx, in, out)
}
