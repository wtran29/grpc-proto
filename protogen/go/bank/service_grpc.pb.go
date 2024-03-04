// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/bank/service.proto

package bank

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

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error)
	FetchExchangeRates(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRatesClient, error)
	SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionsClient, error)
	TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (BankService_TransferMultipleClient, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error) {
	out := new(CurrentBalanceResponse)
	err := c.cc.Invoke(ctx, "/bank.BankService/GetCurrentBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) FetchExchangeRates(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[0], "/bank.BankService/FetchExchangeRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceFetchExchangeRatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BankService_FetchExchangeRatesClient interface {
	Recv() (*ExchangeRateResponse, error)
	grpc.ClientStream
}

type bankServiceFetchExchangeRatesClient struct {
	grpc.ClientStream
}

func (x *bankServiceFetchExchangeRatesClient) Recv() (*ExchangeRateResponse, error) {
	m := new(ExchangeRateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[1], "/bank.BankService/SummarizeTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceSummarizeTransactionsClient{stream}
	return x, nil
}

type BankService_SummarizeTransactionsClient interface {
	Send(*Transaction) error
	CloseAndRecv() (*TransactionSummary, error)
	grpc.ClientStream
}

type bankServiceSummarizeTransactionsClient struct {
	grpc.ClientStream
}

func (x *bankServiceSummarizeTransactionsClient) Send(m *Transaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionsClient) CloseAndRecv() (*TransactionSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransactionSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (BankService_TransferMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[2], "/bank.BankService/TransferMultiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceTransferMultipleClient{stream}
	return x, nil
}

type BankService_TransferMultipleClient interface {
	Send(*TransferRequest) error
	Recv() (*TransferResponse, error)
	grpc.ClientStream
}

type bankServiceTransferMultipleClient struct {
	grpc.ClientStream
}

func (x *bankServiceTransferMultipleClient) Send(m *TransferRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankServiceTransferMultipleClient) Recv() (*TransferResponse, error) {
	m := new(TransferResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/bank.BankService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error)
	FetchExchangeRates(*ExchangeRateRequest, BankService_FetchExchangeRatesServer) error
	SummarizeTransactions(BankService_SummarizeTransactionsServer) error
	TransferMultiple(BankService_TransferMultipleServer) error
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBalance not implemented")
}
func (UnimplementedBankServiceServer) FetchExchangeRates(*ExchangeRateRequest, BankService_FetchExchangeRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchExchangeRates not implemented")
}
func (UnimplementedBankServiceServer) SummarizeTransactions(BankService_SummarizeTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SummarizeTransactions not implemented")
}
func (UnimplementedBankServiceServer) TransferMultiple(BankService_TransferMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferMultiple not implemented")
}
func (UnimplementedBankServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_GetCurrentBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankService/GetCurrentBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, req.(*CurrentBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_FetchExchangeRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExchangeRateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankServiceServer).FetchExchangeRates(m, &bankServiceFetchExchangeRatesServer{stream})
}

type BankService_FetchExchangeRatesServer interface {
	Send(*ExchangeRateResponse) error
	grpc.ServerStream
}

type bankServiceFetchExchangeRatesServer struct {
	grpc.ServerStream
}

func (x *bankServiceFetchExchangeRatesServer) Send(m *ExchangeRateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BankService_SummarizeTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServiceServer).SummarizeTransactions(&bankServiceSummarizeTransactionsServer{stream})
}

type BankService_SummarizeTransactionsServer interface {
	SendAndClose(*TransactionSummary) error
	Recv() (*Transaction, error)
	grpc.ServerStream
}

type bankServiceSummarizeTransactionsServer struct {
	grpc.ServerStream
}

func (x *bankServiceSummarizeTransactionsServer) SendAndClose(m *TransactionSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionsServer) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BankService_TransferMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServiceServer).TransferMultiple(&bankServiceTransferMultipleServer{stream})
}

type BankService_TransferMultipleServer interface {
	Send(*TransferResponse) error
	Recv() (*TransferRequest, error)
	grpc.ServerStream
}

type bankServiceTransferMultipleServer struct {
	grpc.ServerStream
}

func (x *bankServiceTransferMultipleServer) Send(m *TransferResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankServiceTransferMultipleServer) Recv() (*TransferRequest, error) {
	m := new(TransferRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BankService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentBalance",
			Handler:    _BankService_GetCurrentBalance_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _BankService_CreateAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchExchangeRates",
			Handler:       _BankService_FetchExchangeRates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SummarizeTransactions",
			Handler:       _BankService_SummarizeTransactions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferMultiple",
			Handler:       _BankService_TransferMultiple_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/bank/service.proto",
}
