package server

import (
    "net/http"
    "github.com/hoshsadiq/tescoin/block"
    "github.com/hoshsadiq/tescoin/miner"
)

type minedResponse struct {
    Message      string               `json:"message"`
    Index        int                  `json:"index"`
    Transactions []*block.Transaction `json:"transactions"`
    Proof        int                  `json:"proof"`
    PreviousHash string               `json:"previous_hash"`
}

func newMinedResponse(blk *block.Block) *minedResponse {
    return &minedResponse{
        Message:      "Block mined",
        Index:        blk.GetIndex(),
        Transactions: blk.GetTransactions(),
        Proof:        blk.GetProof(),
        PreviousHash: string(blk.GetPreviousHash()),
    }
}

func mine(writer http.ResponseWriter, request *http.Request) {

    newBlock := miner.Mine(blockchain, recipientMinerAddress)

    writeJsonResponse(writer, newMinedResponse(newBlock), http.StatusOK)
}
