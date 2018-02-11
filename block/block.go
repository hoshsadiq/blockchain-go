package block

import (
    "time"
    "github.com/hoshsadiq/blockchain-go/helper"
    "encoding/json"
)

type Block struct {
    Index        int            `json:"index"`
    Timestamp    time.Time      `json:"timestamp"`
    Transactions []*Transaction `json:"transactions"`
    Nonce        int            `json:"nonce"`
    PreviousHash string         `json:"previous_hash"`
}

func NewBlock(index int, timestamp time.Time, transactions []*Transaction, nonce int, previousHash string) *Block {
    return &Block{
        Index:        index,
        Timestamp:    timestamp,
        Transactions: transactions,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
}

func (block *Block) GetIndex() int {
    return block.Index
}

func (block *Block) GetTimestamp() time.Time {
    return block.Timestamp
}

func (block *Block) GetTransactions() []*Transaction {
    return block.Transactions
}

func (block *Block) GetNonce() int {
    return block.Nonce
}

func (block *Block) GetPreviousHash() string {
    return block.PreviousHash
}

func (block *Block) GetHash() string {
    data, _ := json.Marshal(block)
    return helper.GetHash(string(data))
}

