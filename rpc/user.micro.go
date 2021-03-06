// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

/*
Package qcode is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	Void
	Error
*/
package qcode

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	GetAll(ctx context.Context, in *Void, opts ...client.CallOption) (UserService_GetAllService, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	PostMuch(ctx context.Context, opts ...client.CallOption) (UserService_PostMuchService, error)
	Post(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error)
	Put(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "qcode"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetAll(ctx context.Context, in *Void, opts ...client.CallOption) (UserService_GetAllService, error) {
	req := c.c.NewRequest(c.name, "UserService.GetAll", &Void{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceGetAll{stream}, nil
}

type UserService_GetAllService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*User, error)
}

type userServiceGetAll struct {
	stream client.Stream
}

func (x *userServiceGetAll) Close() error {
	return x.stream.Close()
}

func (x *userServiceGetAll) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceGetAll) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceGetAll) Recv() (*User, error) {
	m := new(User)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) PostMuch(ctx context.Context, opts ...client.CallOption) (UserService_PostMuchService, error) {
	req := c.c.NewRequest(c.name, "UserService.PostMuch", &User{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userServicePostMuch{stream}, nil
}

type UserService_PostMuchService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*User) error
}

type userServicePostMuch struct {
	stream client.Stream
}

func (x *userServicePostMuch) Close() error {
	return x.stream.Close()
}

func (x *userServicePostMuch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePostMuch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePostMuch) Send(m *User) error {
	return x.stream.Send(m)
}

func (c *userService) Post(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "UserService.Post", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Put(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "UserService.Put", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	GetAll(context.Context, *Void, UserService_GetAllStream) error
	Get(context.Context, *User, *User) error
	PostMuch(context.Context, UserService_PostMuchStream) error
	Post(context.Context, *User, *Error) error
	Put(context.Context, *User, *Error) error
	Delete(context.Context, *User, *Error) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	type userService interface {
		GetAll(ctx context.Context, stream server.Stream) error
		Get(ctx context.Context, in *User, out *User) error
		PostMuch(ctx context.Context, stream server.Stream) error
		Post(ctx context.Context, in *User, out *Error) error
		Put(ctx context.Context, in *User, out *Error) error
		Delete(ctx context.Context, in *User, out *Error) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetAll(ctx context.Context, stream server.Stream) error {
	m := new(Void)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserServiceHandler.GetAll(ctx, m, &userServiceGetAllStream{stream})
}

type UserService_GetAllStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*User) error
}

type userServiceGetAllStream struct {
	stream server.Stream
}

func (x *userServiceGetAllStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceGetAllStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceGetAllStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceGetAllStream) Send(m *User) error {
	return x.stream.Send(m)
}

func (h *userServiceHandler) Get(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) PostMuch(ctx context.Context, stream server.Stream) error {
	return h.UserServiceHandler.PostMuch(ctx, &userServicePostMuchStream{stream})
}

type UserService_PostMuchStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*User, error)
}

type userServicePostMuchStream struct {
	stream server.Stream
}

func (x *userServicePostMuchStream) Close() error {
	return x.stream.Close()
}

func (x *userServicePostMuchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePostMuchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePostMuchStream) Recv() (*User, error) {
	m := new(User)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *userServiceHandler) Post(ctx context.Context, in *User, out *Error) error {
	return h.UserServiceHandler.Post(ctx, in, out)
}

func (h *userServiceHandler) Put(ctx context.Context, in *User, out *Error) error {
	return h.UserServiceHandler.Put(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *User, out *Error) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}
