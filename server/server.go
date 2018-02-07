package server

import (
    "net/http"
    "github.com/hoshsadiq/tescoin/block"
    "encoding/json"
    "io"
    "fmt"
)

var blockchain *block.Blockchain
var recipientMinerAddress string

type errorResponse struct {
    Message string `json:"message"`
}

func RunServer(bc *block.Blockchain, recipientMinerAddr string, port int) error {
    blockchain = bc
    recipientMinerAddress = recipientMinerAddr
    http.HandleFunc("/mine", post(mine))
    http.HandleFunc("/transactions/new", post(newTransaction))
    http.HandleFunc("/chain", get(getChain))
    http.HandleFunc("/nodes/register", post(registerNode))
    http.HandleFunc("/nodes", get(getNodes))
    http.HandleFunc("/consensus", post(consensus))
    if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil); err != nil {
        return err
    }

    return nil
}

func writeJsonResponse(w http.ResponseWriter, payload interface{}, statusCode int) {
    w.WriteHeader(statusCode)
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(payload)
    io.WriteString(w, string(result))
}

func get(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            h(w, r)
            return
        }
        http.Error(w, "get only", http.StatusMethodNotAllowed)
    }
}

func post(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "post only", http.StatusMethodNotAllowed)
    }
}
