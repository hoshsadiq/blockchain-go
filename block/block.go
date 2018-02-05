package block

import (
    "time"
    "github.com/hoshsadiq/tescoin/helper"
    "fmt"
    "github.com/cloudflare/cfssl/log"
)

type Block struct {
    Index        int            `json:"index"`
    Timestamp    time.Time      `json:"timestamp"`
    Transactions []*Transaction `json:"transactions"`
    Proof        int            `json:"proof"`
    PreviousHash string         `json:"previous_hash"`
}

func NewBlock(index int, timestamp time.Time, transactions []*Transaction, proof int, previousHash string) *Block {
    return &Block{
        Index:        index,
        Timestamp:    timestamp,
        Transactions: transactions,
        Proof:        proof,
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

func (block *Block) GetProof() int {
    return block.Proof
}

func (block *Block) GetPreviousHash() string {
    return block.PreviousHash
}

func (block *Block) GetHash() string {
    return helper.GetHash(fmt.Sprintf("%v", block))
}

func (block *Block) ProofOfWork() int {
    lastHash := block.GetHash()

    proof := 0
    for {
        if helper.ValidProof(block.GetProof(), proof, lastHash) {
            return proof
        }

        proof++
    }

    log.Infof("Proof found: %d", proof)

    return proof
}
