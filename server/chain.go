package server

import (
    "net/http"
    "github.com/hoshsadiq/blockchain-go/block"
)

type chainResponse struct {
    CurrentTransaction []*block.Transaction `json:"current_transaction"`
    Blocks             []*block.Block       `json:"blocks"`
    Length             int                  `json:"length"`
}

func newChainResponse(blockchain *block.Blockchain) *chainResponse {
    return &chainResponse{
        CurrentTransaction: blockchain.CurrentTransactions,
        Blocks:             blockchain.GetBlocks(),
        Length:             len(blockchain.GetBlocks()),
    }
}

func getChain(writer http.ResponseWriter, r *http.Request) {
    writeJsonResponse(writer, newChainResponse(blockchain), http.StatusOK)
}
