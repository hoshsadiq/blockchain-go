package server

import (
    "net/http"
    "github.com/hoshsadiq/blockchain-go/block"
)

type chainReplacedResponse struct {
    Message  string         `json:"message"`
    NewChain []*block.Block `json:"new_chain"`
}

func newChainReplacedResponse(blocks []*block.Block) *chainReplacedResponse {
    return &chainReplacedResponse{
        Message:  "Chain replaced",
        NewChain: blocks,
    }
}

type chainAuthoritativeResponse struct {
    Message string         `json:"message"`
    Chain   []*block.Block `json:"chain"`
}

func newChainAuthoritativeResponse(blocks []*block.Block) *chainAuthoritativeResponse {
    return &chainAuthoritativeResponse{
        Message: "Chain is authoritative",
        Chain:   blocks,
    }
}

func consensus(writer http.ResponseWriter, request *http.Request) {
    if blockchain.Consensus() {
        writeJsonResponse(writer, newChainReplacedResponse(blockchain.GetBlocks()), http.StatusOK)
    } else {
        writeJsonResponse(writer, newChainAuthoritativeResponse(blockchain.GetBlocks()), http.StatusOK)
    }
}
