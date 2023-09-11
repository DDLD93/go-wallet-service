 package grpc
// import (
//     "log"
//     "net"

//     "google.golang.org/grpc"
// )

// type server struct{}

// // func (s *server) CreateWallet(ctx context.Context, req *pb.WalletRequest) (*pb.Response, error) {
// //     // Implement logic to handle CreateWallet request
// // }

// // func (s *server) GetWalletBalance(ctx context.Context, req *pb.UserID) (*pb.BalanceResponse, error) {
// //     // Implement logic to handle GetWalletBalance request
// // }

// // func (s *server) CreateTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Response, error) {
// //     // Implement logic to handle CreateTransaction request
// // }

// // func (s *server) GetWalletTransactionHistory(ctx context.Context, req *pb.UserID) (*pb.TransactionHistory, error) {
// //     // Implement logic to handle GetWalletTransactionHistory request
// // }

// func StartGRPCServer() {
//     lis, err := net.Listen("tcp", ":50051")
//     if err != nil {
//         log.Fatalf("failed to listen: %v", err)
//     }
//     s := grpc.NewServer()
//     pb.RegisterWalletServiceServer(s, &server{})
//     if err := s.Serve(lis); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }
