package server

import (
    "net/http"
    "github.com/hoshsadiq/tescoin/block"
)

type chainResponse struct {
    Blocks []*block.Block `json:"blocks"`
    Length int            `json:"length"`
}

func newChainResponse(blocks []*block.Block) *chainResponse {
    return &chainResponse{
        Blocks: blocks,
        Length: len(blocks),
    }
}

func getChain(writer http.ResponseWriter, r *http.Request) {
    writeJsonResponse(writer, newChainResponse(blockchain.GetBlocks()), http.StatusOK)
}
