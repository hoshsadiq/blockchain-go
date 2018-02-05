package server

import (
    "net/http"
    "github.com/hoshsadiq/tescoin/block"
)

type transactionCreatedResponse struct {
    Message     string            `json:"message"`
    Index       int               `json:"index"`
    Transaction block.Transaction `json:"transaction"`
}

func newTransactionCreatedReponse(index int, tx block.Transaction) *transactionCreatedResponse {
    return &transactionCreatedResponse{
        Message:     "Transaction added to blockchain",
        Index:       index,
        Transaction: tx,
    }
}

func newTransaction(writer http.ResponseWriter, request *http.Request) {
    var tx block.Transaction
    if !requestBodyToPayload(writer, request, &tx) {
        return
    }

    index := blockchain.AddTransaction(&tx)

    writeJsonResponse(writer, newTransactionCreatedReponse(index, tx), http.StatusOK)
}
