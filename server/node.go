package server

import (
    "net/http"
    "github.com/hoshsadiq/tescoin/node"
    "github.com/hoshsadiq/tescoin/block"
)

type nodesRegisteredResponse struct {
    Message    string `json:"message"`
    TotalNodes int    `json:"total_nodes"`
}

func newNodesRegisteredResponse(nodes []*node.Node) *nodesRegisteredResponse {
    return &nodesRegisteredResponse{
        Message:    "New nodes registered",
        TotalNodes: len(nodes),
    }
}

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

func registerNode(writer http.ResponseWriter, request *http.Request) {
    var nodes []node.Node
    if !requestBodyToPayload(writer, request, &nodes) {
        return
    }

    for _, n := range nodes {
        blockchain.RegisterNode(&n)
    }

    writeJsonResponse(writer, newNodesRegisteredResponse(blockchain.GetNodes()), http.StatusOK)
}

func consensus(writer http.ResponseWriter, request *http.Request) {
    if request.Method != "POST" {
        http.Error(writer, "Must use POST", 400)
    }

    if blockchain.Consensus() {
        writeJsonResponse(writer, newChainReplacedResponse(blockchain.GetBlocks()), http.StatusOK)
    } else {
        writeJsonResponse(writer, newChainAuthoritativeResponse(blockchain.GetBlocks()), http.StatusOK)
    }
}
