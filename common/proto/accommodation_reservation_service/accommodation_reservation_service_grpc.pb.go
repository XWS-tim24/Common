// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: accommodation_reservation_service/accommodation_reservation_service.proto

package accommodation_reservation

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

// AccommodationReservationServiceClient is the client API for AccommodationReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationReservationServiceClient interface {
	GetRequestById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	GetAllPendingForUser(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetAllPendingForUserResponse, error)
	GetAllPendingForAccomodation(ctx context.Context, in *GetAllPendingForAccRequest, opts ...grpc.CallOption) (*GetAllPendingForAccResponse, error)
	GetReservationById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetReservationByIdResponse, error)
	GetNumberOfCanceled(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetIntResponse, error)
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	CreateRequest(ctx context.Context, in *CreateReservationRequestRequest, opts ...grpc.CallOption) (*CreateReservationRequestResponse, error)
}

type accommodationReservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationReservationServiceClient(cc grpc.ClientConnInterface) AccommodationReservationServiceClient {
	return &accommodationReservationServiceClient{cc}
}

func (c *accommodationReservationServiceClient) GetRequestById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/GetRequestById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) GetAllPendingForUser(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetAllPendingForUserResponse, error) {
	out := new(GetAllPendingForUserResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/GetAllPendingForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) GetAllPendingForAccomodation(ctx context.Context, in *GetAllPendingForAccRequest, opts ...grpc.CallOption) (*GetAllPendingForAccResponse, error) {
	out := new(GetAllPendingForAccResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/GetAllPendingForAccomodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) GetReservationById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetReservationByIdResponse, error) {
	out := new(GetReservationByIdResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/GetReservationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) GetNumberOfCanceled(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetIntResponse, error) {
	out := new(GetIntResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/GetNumberOfCanceled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationReservationServiceClient) CreateRequest(ctx context.Context, in *CreateReservationRequestRequest, opts ...grpc.CallOption) (*CreateReservationRequestResponse, error) {
	out := new(CreateReservationRequestResponse)
	err := c.cc.Invoke(ctx, "/accommodation_reservation.AccommodationReservationService/CreateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationReservationServiceServer is the server API for AccommodationReservationService service.
// All implementations must embed UnimplementedAccommodationReservationServiceServer
// for forward compatibility
type AccommodationReservationServiceServer interface {
	GetRequestById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	GetAllPendingForUser(context.Context, *GetByUserIdRequest) (*GetAllPendingForUserResponse, error)
	GetAllPendingForAccomodation(context.Context, *GetAllPendingForAccRequest) (*GetAllPendingForAccResponse, error)
	GetReservationById(context.Context, *GetByIdRequest) (*GetReservationByIdResponse, error)
	GetNumberOfCanceled(context.Context, *GetByUserIdRequest) (*GetIntResponse, error)
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	CreateRequest(context.Context, *CreateReservationRequestRequest) (*CreateReservationRequestResponse, error)
	mustEmbedUnimplementedAccommodationReservationServiceServer()
}

// UnimplementedAccommodationReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationReservationServiceServer struct {
}

func (UnimplementedAccommodationReservationServiceServer) GetRequestById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestById not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) GetAllPendingForUser(context.Context, *GetByUserIdRequest) (*GetAllPendingForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPendingForUser not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) GetAllPendingForAccomodation(context.Context, *GetAllPendingForAccRequest) (*GetAllPendingForAccResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPendingForAccomodation not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) GetReservationById(context.Context, *GetByIdRequest) (*GetReservationByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationById not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) GetNumberOfCanceled(context.Context, *GetByUserIdRequest) (*GetIntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfCanceled not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) CreateRequest(context.Context, *CreateReservationRequestRequest) (*CreateReservationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedAccommodationReservationServiceServer) mustEmbedUnimplementedAccommodationReservationServiceServer() {
}

// UnsafeAccommodationReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationReservationServiceServer will
// result in compilation errors.
type UnsafeAccommodationReservationServiceServer interface {
	mustEmbedUnimplementedAccommodationReservationServiceServer()
}

func RegisterAccommodationReservationServiceServer(s grpc.ServiceRegistrar, srv AccommodationReservationServiceServer) {
	s.RegisterService(&AccommodationReservationService_ServiceDesc, srv)
}

func _AccommodationReservationService_GetRequestById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).GetRequestById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/GetRequestById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).GetRequestById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_GetAllPendingForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).GetAllPendingForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/GetAllPendingForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).GetAllPendingForUser(ctx, req.(*GetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_GetAllPendingForAccomodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPendingForAccRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).GetAllPendingForAccomodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/GetAllPendingForAccomodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).GetAllPendingForAccomodation(ctx, req.(*GetAllPendingForAccRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_GetReservationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).GetReservationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/GetReservationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).GetReservationById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_GetNumberOfCanceled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).GetNumberOfCanceled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/GetNumberOfCanceled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).GetNumberOfCanceled(ctx, req.(*GetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationReservationService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationReservationServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_reservation.AccommodationReservationService/CreateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationReservationServiceServer).CreateRequest(ctx, req.(*CreateReservationRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationReservationService_ServiceDesc is the grpc.ServiceDesc for AccommodationReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation_reservation.AccommodationReservationService",
	HandlerType: (*AccommodationReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequestById",
			Handler:    _AccommodationReservationService_GetRequestById_Handler,
		},
		{
			MethodName: "GetAllPendingForUser",
			Handler:    _AccommodationReservationService_GetAllPendingForUser_Handler,
		},
		{
			MethodName: "GetAllPendingForAccomodation",
			Handler:    _AccommodationReservationService_GetAllPendingForAccomodation_Handler,
		},
		{
			MethodName: "GetReservationById",
			Handler:    _AccommodationReservationService_GetReservationById_Handler,
		},
		{
			MethodName: "GetNumberOfCanceled",
			Handler:    _AccommodationReservationService_GetNumberOfCanceled_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _AccommodationReservationService_CreateReservation_Handler,
		},
		{
			MethodName: "CreateRequest",
			Handler:    _AccommodationReservationService_CreateRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_reservation_service/accommodation_reservation_service.proto",
}
