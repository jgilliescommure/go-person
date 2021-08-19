// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package person

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PersonClient is the client API for Person service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonReply, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*PersonReply, error)
	GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*PersonReply, error)
	ListPersons(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (*ListPersonReply, error)
}

type personClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonClient(cc grpc.ClientConnInterface) PersonClient {
	return &personClient{cc}
}

func (c *personClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonReply, error) {
	out := new(CreatePersonReply)
	err := c.cc.Invoke(ctx, "/api.person.Person/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := c.cc.Invoke(ctx, "/api.person.Person/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) GetPerson(ctx context.Context, in *GetPersonRequest, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := c.cc.Invoke(ctx, "/api.person.Person/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) ListPersons(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (*ListPersonReply, error) {
	out := new(ListPersonReply)
	err := c.cc.Invoke(ctx, "/api.person.Person/ListPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServer is the server API for Person service.
// All implementations must embed UnimplementedPersonServer
// for forward compatibility
type PersonServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonReply, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*PersonReply, error)
	GetPerson(context.Context, *GetPersonRequest) (*PersonReply, error)
	ListPersons(context.Context, *ListPersonRequest) (*ListPersonReply, error)
	mustEmbedUnimplementedPersonServer()
}

// UnimplementedPersonServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServer struct {
}

func (UnimplementedPersonServer) CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServer) UpdatePerson(context.Context, *UpdatePersonRequest) (*PersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedPersonServer) GetPerson(context.Context, *GetPersonRequest) (*PersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedPersonServer) ListPersons(context.Context, *ListPersonRequest) (*ListPersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPersons not implemented")
}
func (UnimplementedPersonServer) mustEmbedUnimplementedPersonServer() {}

// UnsafePersonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServer will
// result in compilation errors.
type UnsafePersonServer interface {
	mustEmbedUnimplementedPersonServer()
}

func RegisterPersonServer(s grpc.ServiceRegistrar, srv PersonServer) {
	s.RegisterService(&Person_ServiceDesc, srv)
}

func _Person_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.person.Person/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.person.Person/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.person.Person/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).GetPerson(ctx, req.(*GetPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_ListPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).ListPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.person.Person/ListPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).ListPersons(ctx, req.(*ListPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Person_ServiceDesc is the grpc.ServiceDesc for Person service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Person_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.person.Person",
	HandlerType: (*PersonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _Person_CreatePerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _Person_UpdatePerson_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _Person_GetPerson_Handler,
		},
		{
			MethodName: "ListPersons",
			Handler:    _Person_ListPersons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/person/person.proto",
}