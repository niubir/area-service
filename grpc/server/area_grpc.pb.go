// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: area.proto

package server

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

// AreaServiceClient is the client API for AreaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaServiceClient interface {
	// 刷新区域
	FreshAreas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// 获取指定区域
	GetArea(ctx context.Context, in *GetAreaReq, opts ...grpc.CallOption) (*GetAreaAck, error)
	// 获取省份列表
	GetProvinces(ctx context.Context, in *GetProvincesReq, opts ...grpc.CallOption) (*GetProvincesAck, error)
	// 获取城市列表
	GetCities(ctx context.Context, in *GetCitiesReq, opts ...grpc.CallOption) (*GetCitiesAck, error)
	// 获取区县列表
	GetDistricts(ctx context.Context, in *GetDistrictsReq, opts ...grpc.CallOption) (*GetDistrictsAck, error)
	// 获取街道列表
	GetStreets(ctx context.Context, in *GetStreetsReq, opts ...grpc.CallOption) (*GetStreetsAck, error)
}

type areaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaServiceClient(cc grpc.ClientConnInterface) AreaServiceClient {
	return &areaServiceClient{cc}
}

func (c *areaServiceClient) FreshAreas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/area.AreaService/FreshAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) GetArea(ctx context.Context, in *GetAreaReq, opts ...grpc.CallOption) (*GetAreaAck, error) {
	out := new(GetAreaAck)
	err := c.cc.Invoke(ctx, "/area.AreaService/GetArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) GetProvinces(ctx context.Context, in *GetProvincesReq, opts ...grpc.CallOption) (*GetProvincesAck, error) {
	out := new(GetProvincesAck)
	err := c.cc.Invoke(ctx, "/area.AreaService/GetProvinces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) GetCities(ctx context.Context, in *GetCitiesReq, opts ...grpc.CallOption) (*GetCitiesAck, error) {
	out := new(GetCitiesAck)
	err := c.cc.Invoke(ctx, "/area.AreaService/GetCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) GetDistricts(ctx context.Context, in *GetDistrictsReq, opts ...grpc.CallOption) (*GetDistrictsAck, error) {
	out := new(GetDistrictsAck)
	err := c.cc.Invoke(ctx, "/area.AreaService/GetDistricts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaServiceClient) GetStreets(ctx context.Context, in *GetStreetsReq, opts ...grpc.CallOption) (*GetStreetsAck, error) {
	out := new(GetStreetsAck)
	err := c.cc.Invoke(ctx, "/area.AreaService/GetStreets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaServiceServer is the server API for AreaService service.
// All implementations must embed UnimplementedAreaServiceServer
// for forward compatibility
type AreaServiceServer interface {
	// 刷新区域
	FreshAreas(context.Context, *Empty) (*Empty, error)
	// 获取指定区域
	GetArea(context.Context, *GetAreaReq) (*GetAreaAck, error)
	// 获取省份列表
	GetProvinces(context.Context, *GetProvincesReq) (*GetProvincesAck, error)
	// 获取城市列表
	GetCities(context.Context, *GetCitiesReq) (*GetCitiesAck, error)
	// 获取区县列表
	GetDistricts(context.Context, *GetDistrictsReq) (*GetDistrictsAck, error)
	// 获取街道列表
	GetStreets(context.Context, *GetStreetsReq) (*GetStreetsAck, error)
	mustEmbedUnimplementedAreaServiceServer()
}

// UnimplementedAreaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAreaServiceServer struct {
}

func (UnimplementedAreaServiceServer) FreshAreas(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreshAreas not implemented")
}
func (UnimplementedAreaServiceServer) GetArea(context.Context, *GetAreaReq) (*GetAreaAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArea not implemented")
}
func (UnimplementedAreaServiceServer) GetProvinces(context.Context, *GetProvincesReq) (*GetProvincesAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvinces not implemented")
}
func (UnimplementedAreaServiceServer) GetCities(context.Context, *GetCitiesReq) (*GetCitiesAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCities not implemented")
}
func (UnimplementedAreaServiceServer) GetDistricts(context.Context, *GetDistrictsReq) (*GetDistrictsAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistricts not implemented")
}
func (UnimplementedAreaServiceServer) GetStreets(context.Context, *GetStreetsReq) (*GetStreetsAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreets not implemented")
}
func (UnimplementedAreaServiceServer) mustEmbedUnimplementedAreaServiceServer() {}

// UnsafeAreaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaServiceServer will
// result in compilation errors.
type UnsafeAreaServiceServer interface {
	mustEmbedUnimplementedAreaServiceServer()
}

func RegisterAreaServiceServer(s grpc.ServiceRegistrar, srv AreaServiceServer) {
	s.RegisterService(&AreaService_ServiceDesc, srv)
}

func _AreaService_FreshAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).FreshAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/FreshAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).FreshAreas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_GetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).GetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/GetArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).GetArea(ctx, req.(*GetAreaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_GetProvinces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProvincesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).GetProvinces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/GetProvinces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).GetProvinces(ctx, req.(*GetProvincesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_GetCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCitiesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).GetCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/GetCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).GetCities(ctx, req.(*GetCitiesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_GetDistricts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistrictsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).GetDistricts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/GetDistricts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).GetDistricts(ctx, req.(*GetDistrictsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaService_GetStreets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServiceServer).GetStreets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.AreaService/GetStreets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServiceServer).GetStreets(ctx, req.(*GetStreetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AreaService_ServiceDesc is the grpc.ServiceDesc for AreaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AreaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "area.AreaService",
	HandlerType: (*AreaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FreshAreas",
			Handler:    _AreaService_FreshAreas_Handler,
		},
		{
			MethodName: "GetArea",
			Handler:    _AreaService_GetArea_Handler,
		},
		{
			MethodName: "GetProvinces",
			Handler:    _AreaService_GetProvinces_Handler,
		},
		{
			MethodName: "GetCities",
			Handler:    _AreaService_GetCities_Handler,
		},
		{
			MethodName: "GetDistricts",
			Handler:    _AreaService_GetDistricts_Handler,
		},
		{
			MethodName: "GetStreets",
			Handler:    _AreaService_GetStreets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "area.proto",
}
