package grpc

// import (
//     context "context"
//     proto "google.golang.org/protobuf/proto"
//     grpc "google.golang.org/grpc"
// )

// type WalletServiceClient interface {
//     CreateWallet(ctx context.Context, in *WalletRequest, opts ...grpc.CallOption) (*Response, error)
//     GetWalletBalance(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*BalanceResponse, error)
//     CreateTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Response, error)
//     GetWalletTransactionHistory(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*TransactionHistory, error)
// }

// type WalletServiceServer interface {
//     CreateWallet(context.Context, *WalletRequest) (*Response, error)
//     GetWalletBalance(context.Context, *UserID) (*BalanceResponse, error)
//     CreateTransaction(context.Context, *TransactionRequest) (*Response, error)
//     GetWalletTransactionHistory(context.Context, *UserID) (*TransactionHistory, error)
// }

// type WalletServiceClient struct {
//     cc grpc.ClientConnInterface
// }

// func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
//     return &WalletServiceClient{cc}
// }

// func (c *WalletServiceClient) CreateWallet(ctx context.Context, in *WalletRequest, opts ...grpc.CallOption) (*Response, error) {
//     out := new(Response)
//     err := c.cc.Invoke(ctx, "/wallet.WalletService/CreateWallet", in, out, opts...)
//     if err != nil {
//         return nil, err
//     }
//     return out, nil
// }

// func (c *WalletServiceClient) GetWalletBalance(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*BalanceResponse, error) {
//     out := new(BalanceResponse)
//     err := c.cc.Invoke(ctx, "/wallet.WalletService/GetWalletBalance", in, out, opts...)
//     if err != nil {
//         return nil, err
//     }
//     return out, nil
// }

// func (c *WalletServiceClient) CreateTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Response, error) {
//     out := new(Response)
//     err := c.cc.Invoke(ctx, "/wallet.WalletService/CreateTransaction", in, out, opts...)
//     if err != nil {
//         return nil, err
//     }
//     return out, nil
// }

// func (c *WalletServiceClient) GetWalletTransactionHistory(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*TransactionHistory, error) {
//     out := new(TransactionHistory)
//     err := c.cc.Invoke(ctx, "/wallet.WalletService/GetWalletTransactionHistory", in, out, opts...)
//     if err != nil {
//         return nil, err
//     }
//     return out, nil
// }

// type WalletServiceServer struct {
//     // Implement the server methods
// }

// func NewWalletServiceServer(srv WalletServiceServer) *grpc.Server {
//     return grpc.NewServer(grpc.UnaryInterceptor(nil))
// }

// func (s *WalletServiceServer) RegisterService(grpcService *grpc.ServiceDesc, impl interface{}) {
//     s.srv.RegisterService(grpcService, impl)
// }

// func RegisterWalletServiceServer(s *grpc.Server, srv WalletServiceServer) {
//     s.RegisterService(&WalletService_ServiceDesc, srv)
// }
