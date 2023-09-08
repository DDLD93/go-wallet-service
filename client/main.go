package main

import (
    "log"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "github.com/DDLD93/go-wallet-service/src/models"
    "wallet"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := wallet.NewWalletServiceClient(conn)

    // Example: Call gRPC methods using the client.
    walletReq := &wallet.WalletRequest{
        UserId:   "123",
        Email:    "user@example.com",
        Balance:  100.0,
        Currency: "USD",
    }

    createResp, err := client.CreateWallet(context.Background(), walletReq)
    if err != nil {
        log.Fatalf("CreateWallet failed: %v", err)
    }
    log.Printf("CreateWallet response: %+v", createResp)

    // Implement calls to other gRPC methods here.
}
