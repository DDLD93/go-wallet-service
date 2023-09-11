// paystack/paystack.go

package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
    paystackAPIURL = "https://api.paystack.co"
)

// TransactionRequest represents the request for initiating a Paystack transaction.
type TransactionRequest struct {
    Email      string  `json:"email"`
    Amount     float64 `json:"amount"`
    Reference  string  `json:"reference"`
    CallbackURL string `json:"callback_url"`
}

// TransactionResponse represents the response from Paystack when initiating a transaction.
type TransactionResponse struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        AuthorizationURL string `json:"authorization_url"`
        AccessCode       string `json:"access_code"`
    } `json:"data"`
}
type TransactionCallback struct {
    Event string `json:"event"`
    Data  struct {
        Reference      string `json:"reference"`
        TransactionID  int    `json:"transaction_id"`
        Amount         int    `json:"amount"`
        Currency       string `json:"currency"`
        Status         string `json:"status"`
        FailureReason  string `json:"failure_reason"`
        Customer       struct {
            Email string `json:"email"`
        } `json:"customer"`
    } `json:"data"`
}

// InitiateTransaction initiates a transaction with Paystack.
func InitiateTransaction(secretKey string, transactionReq TransactionRequest) (TransactionResponse, error) {
    // Convert the transaction request to JSON
    payload, err := json.Marshal(transactionReq)
    if err != nil {
        return TransactionResponse{}, err
    }

    // Create an HTTP client
    client := &http.Client{}

    // Create a POST request to Paystack's transaction initiation endpoint
    req, err := http.NewRequest("POST", paystackAPIURL+"/transaction/initialize", bytes.NewBuffer(payload))
    if err != nil {
        return TransactionResponse{}, err
    }

    // Set the Paystack secret key in the request headers
    req.Header.Set("Authorization", "Bearer "+secretKey)
    req.Header.Set("Content-Type", "application/json")

    // Send the HTTP request
    resp, err := client.Do(req)
    if err != nil {
        return TransactionResponse{}, err
    }
    defer resp.Body.Close()

    // Parse the response
    var transactionResp TransactionResponse
    if err := json.NewDecoder(resp.Body).Decode(&transactionResp); err != nil {
        return TransactionResponse{}, err
    }

    return transactionResp, nil
}

func HandlePaystackCallback(c *gin.Context) {
    // Parse the JSON data from the request body
    var callbackData TransactionCallback
    if err := c.BindJSON(&callbackData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding callback data"})
        return
    }

    // Verify the transaction status
    if callbackData.Event == "charge.success" {
        fmt.Printf("Transaction ID %s was successful\n", callbackData.Data.Reference)
        // You can update your database or perform other actions here.
    } else {
        fmt.Printf("Transaction ID %s failed with message: %s\n", callbackData.Data.Reference, callbackData.Data.FailureReason)
        // Handle the failure scenario here.
    }

    // Respond with a 200 OK status to acknowledge receipt of the callback
	fmt.Print(callbackData)
    c.JSON(http.StatusOK, gin.H{"message": "Callback received"})
}

