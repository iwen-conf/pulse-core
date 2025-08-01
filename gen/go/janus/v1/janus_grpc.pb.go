// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: janus/v1/janus.proto

package janusv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LedgerService_CreateAccount_FullMethodName            = "/janus.LedgerService/CreateAccount"
	LedgerService_GetAccount_FullMethodName               = "/janus.LedgerService/GetAccount"
	LedgerService_ListAccounts_FullMethodName             = "/janus.LedgerService/ListAccounts"
	LedgerService_UpdateAccount_FullMethodName            = "/janus.LedgerService/UpdateAccount"
	LedgerService_RecordJournalEntry_FullMethodName       = "/janus.LedgerService/RecordJournalEntry"
	LedgerService_GetJournalEntry_FullMethodName          = "/janus.LedgerService/GetJournalEntry"
	LedgerService_ListJournalEntries_FullMethodName       = "/janus.LedgerService/ListJournalEntries"
	LedgerService_ListJournalEntriesByUser_FullMethodName = "/janus.LedgerService/ListJournalEntriesByUser"
	LedgerService_ReverseJournalEntry_FullMethodName      = "/janus.LedgerService/ReverseJournalEntry"
)

// LedgerServiceClient is the client API for LedgerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ===================================================================
// LedgerService 定义
// ===================================================================
// LedgerService 是记账核心的统一接口。
type LedgerServiceClient interface {
	// 为一个已存在的实体创建一个会计科目。
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	// 获取一个会计科目的详细信息，包括实时余额。
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
	// 列出一个核算实体的所有会计科目（即其专属的会计科目表）。
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	// 更新一个会计科目的信息（例如：修改名称或禁用它）。
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	// 记录一笔记账分录 (核心写操作)。
	RecordJournalEntry(ctx context.Context, in *RecordJournalEntryRequest, opts ...grpc.CallOption) (*RecordJournalEntryResponse, error)
	// 获取一笔记账凭证的详细信息。
	GetJournalEntry(ctx context.Context, in *GetJournalEntryRequest, opts ...grpc.CallOption) (*JournalEntry, error)
	// 列出一个核算实体的记账凭证历史。
	ListJournalEntries(ctx context.Context, in *ListJournalEntriesRequest, opts ...grpc.CallOption) (*ListJournalEntriesResponse, error)
	// 列出一个用户处理过的所有记账凭证。
	ListJournalEntriesByUser(ctx context.Context, in *ListJournalEntriesByUserRequest, opts ...grpc.CallOption) (*ListJournalEntriesResponse, error)
	// 冲正一笔已存在的记账凭证，会创建一笔金额完全相反的新凭证。
	ReverseJournalEntry(ctx context.Context, in *ReverseJournalEntryRequest, opts ...grpc.CallOption) (*RecordJournalEntryResponse, error)
}

type ledgerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerServiceClient(cc grpc.ClientConnInterface) LedgerServiceClient {
	return &ledgerServiceClient{cc}
}

func (c *ledgerServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, LedgerService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, LedgerService_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, LedgerService_ListAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, LedgerService_UpdateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) RecordJournalEntry(ctx context.Context, in *RecordJournalEntryRequest, opts ...grpc.CallOption) (*RecordJournalEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordJournalEntryResponse)
	err := c.cc.Invoke(ctx, LedgerService_RecordJournalEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) GetJournalEntry(ctx context.Context, in *GetJournalEntryRequest, opts ...grpc.CallOption) (*JournalEntry, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JournalEntry)
	err := c.cc.Invoke(ctx, LedgerService_GetJournalEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListJournalEntries(ctx context.Context, in *ListJournalEntriesRequest, opts ...grpc.CallOption) (*ListJournalEntriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListJournalEntriesResponse)
	err := c.cc.Invoke(ctx, LedgerService_ListJournalEntries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ListJournalEntriesByUser(ctx context.Context, in *ListJournalEntriesByUserRequest, opts ...grpc.CallOption) (*ListJournalEntriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListJournalEntriesResponse)
	err := c.cc.Invoke(ctx, LedgerService_ListJournalEntriesByUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) ReverseJournalEntry(ctx context.Context, in *ReverseJournalEntryRequest, opts ...grpc.CallOption) (*RecordJournalEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordJournalEntryResponse)
	err := c.cc.Invoke(ctx, LedgerService_ReverseJournalEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerServiceServer is the server API for LedgerService service.
// All implementations must embed UnimplementedLedgerServiceServer
// for forward compatibility.
//
// ===================================================================
// LedgerService 定义
// ===================================================================
// LedgerService 是记账核心的统一接口。
type LedgerServiceServer interface {
	// 为一个已存在的实体创建一个会计科目。
	CreateAccount(context.Context, *CreateAccountRequest) (*Account, error)
	// 获取一个会计科目的详细信息，包括实时余额。
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
	// 列出一个核算实体的所有会计科目（即其专属的会计科目表）。
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	// 更新一个会计科目的信息（例如：修改名称或禁用它）。
	UpdateAccount(context.Context, *UpdateAccountRequest) (*Account, error)
	// 记录一笔记账分录 (核心写操作)。
	RecordJournalEntry(context.Context, *RecordJournalEntryRequest) (*RecordJournalEntryResponse, error)
	// 获取一笔记账凭证的详细信息。
	GetJournalEntry(context.Context, *GetJournalEntryRequest) (*JournalEntry, error)
	// 列出一个核算实体的记账凭证历史。
	ListJournalEntries(context.Context, *ListJournalEntriesRequest) (*ListJournalEntriesResponse, error)
	// 列出一个用户处理过的所有记账凭证。
	ListJournalEntriesByUser(context.Context, *ListJournalEntriesByUserRequest) (*ListJournalEntriesResponse, error)
	// 冲正一笔已存在的记账凭证，会创建一笔金额完全相反的新凭证。
	ReverseJournalEntry(context.Context, *ReverseJournalEntryRequest) (*RecordJournalEntryResponse, error)
	mustEmbedUnimplementedLedgerServiceServer()
}

// UnimplementedLedgerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLedgerServiceServer struct{}

func (UnimplementedLedgerServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedLedgerServiceServer) GetAccount(context.Context, *GetAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedLedgerServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedLedgerServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedLedgerServiceServer) RecordJournalEntry(context.Context, *RecordJournalEntryRequest) (*RecordJournalEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordJournalEntry not implemented")
}
func (UnimplementedLedgerServiceServer) GetJournalEntry(context.Context, *GetJournalEntryRequest) (*JournalEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournalEntry not implemented")
}
func (UnimplementedLedgerServiceServer) ListJournalEntries(context.Context, *ListJournalEntriesRequest) (*ListJournalEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJournalEntries not implemented")
}
func (UnimplementedLedgerServiceServer) ListJournalEntriesByUser(context.Context, *ListJournalEntriesByUserRequest) (*ListJournalEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJournalEntriesByUser not implemented")
}
func (UnimplementedLedgerServiceServer) ReverseJournalEntry(context.Context, *ReverseJournalEntryRequest) (*RecordJournalEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseJournalEntry not implemented")
}
func (UnimplementedLedgerServiceServer) mustEmbedUnimplementedLedgerServiceServer() {}
func (UnimplementedLedgerServiceServer) testEmbeddedByValue()                       {}

// UnsafeLedgerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerServiceServer will
// result in compilation errors.
type UnsafeLedgerServiceServer interface {
	mustEmbedUnimplementedLedgerServiceServer()
}

func RegisterLedgerServiceServer(s grpc.ServiceRegistrar, srv LedgerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLedgerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LedgerService_ServiceDesc, srv)
}

func _LedgerService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_ListAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_RecordJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).RecordJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_RecordJournalEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).RecordJournalEntry(ctx, req.(*RecordJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_GetJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).GetJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_GetJournalEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).GetJournalEntry(ctx, req.(*GetJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListJournalEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJournalEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListJournalEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_ListJournalEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListJournalEntries(ctx, req.(*ListJournalEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ListJournalEntriesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJournalEntriesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ListJournalEntriesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_ListJournalEntriesByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ListJournalEntriesByUser(ctx, req.(*ListJournalEntriesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_ReverseJournalEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseJournalEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).ReverseJournalEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LedgerService_ReverseJournalEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).ReverseJournalEntry(ctx, req.(*ReverseJournalEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LedgerService_ServiceDesc is the grpc.ServiceDesc for LedgerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LedgerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "janus.LedgerService",
	HandlerType: (*LedgerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _LedgerService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _LedgerService_GetAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _LedgerService_ListAccounts_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _LedgerService_UpdateAccount_Handler,
		},
		{
			MethodName: "RecordJournalEntry",
			Handler:    _LedgerService_RecordJournalEntry_Handler,
		},
		{
			MethodName: "GetJournalEntry",
			Handler:    _LedgerService_GetJournalEntry_Handler,
		},
		{
			MethodName: "ListJournalEntries",
			Handler:    _LedgerService_ListJournalEntries_Handler,
		},
		{
			MethodName: "ListJournalEntriesByUser",
			Handler:    _LedgerService_ListJournalEntriesByUser_Handler,
		},
		{
			MethodName: "ReverseJournalEntry",
			Handler:    _LedgerService_ReverseJournalEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "janus/v1/janus.proto",
}
