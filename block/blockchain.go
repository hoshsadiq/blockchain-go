package block

import (
    "time"
    "github.com/hoshsadiq/tescoin/node"
    "github.com/hoshsadiq/tescoin/helper"
    "log"
)

type Blockchain struct {
    CurrentTransactions []*Transaction
    Blocks              []*Block
    Nodes               []*node.Node
}

func NewBlockchain() *Blockchain {
    blockchain := Blockchain{
        CurrentTransactions: []*Transaction{},
        Blocks:              []*Block{},
        Nodes:               []*node.Node{},
    }

    // genesis block!
    blockchain.NewBlock(100, "__genesis")

    return &blockchain
}

func (blockchain *Blockchain) NewBlock(nonce int, previousHash string) *Block {
    if previousHash == "" {
        panic("previous hash cannot be empty")
    }

    block := NewBlock(len(blockchain.Blocks), time.Now(), blockchain.CurrentTransactions, nonce, previousHash)

    blockchain.CurrentTransactions = []*Transaction{}
    blockchain.AddBlock(block)

    return block
}

func (blockchain *Blockchain) AddBlock(block *Block) {
    blockchain.Blocks = append(blockchain.Blocks, block)
}

func (blockchain *Blockchain) AddTransaction(tx *Transaction) int {
    blockchain.CurrentTransactions = append(blockchain.CurrentTransactions, tx)

    return blockchain.GetLastBlock().GetIndex() + 1
}

func (blockchain *Blockchain) GetLastBlock() *Block {
    return blockchain.Blocks[len(blockchain.Blocks)-1]
}

func (blockchain *Blockchain) RegisterNode(node node.Node) {
    blockchain.Nodes = append(blockchain.Nodes, &node)
}

func (blockchain *Blockchain) Consensus() (replaced bool) {
    var newBlocks []*Block
    neighbours := blockchain.Nodes

    maxLength := len(blockchain.Blocks)

    for _, n := range neighbours {
        response := helper.GetUrl(n.GetChainUrl())
        if response != nil {
            length := int(response["length"].(float64))
            var blocks []*Block
            helper.ConvertInterface(response["blocks"], &blocks)

            if length > maxLength && IsValidChain(blocks) {
                maxLength = length
                newBlocks = blocks
            }
        }
    }

    if newBlocks != nil {
        blockchain.Blocks = newBlocks
        return true
    }

    return false
}

func (blockchain *Blockchain) GetBlocks() []*Block {
    return blockchain.Blocks
}

func (blockchain *Blockchain) GetNodes() []*node.Node {
    return blockchain.Nodes
}

func IsValidChain(blocks []*Block) bool {
    var lastBlock *Block
    for idx, blk := range blocks {
        if idx == 0 {
            lastBlock = blk
            continue
        }

        previousHash := blk.GetPreviousHash()
        if previousHash != lastBlock.GetHash() {
            return false
        }

        if !helper.ValidNonce(lastBlock.GetNonce(), blk.GetNonce(), previousHash) {
            return false
        }

        lastBlock = blk
        idx++
    }

    return true
}

func (blockchain *Blockchain) ProofOfWork(lastBlock *Block) int {
    lastHash := lastBlock.GetHash()
    lastNonce := lastBlock.GetNonce()

    nonce := 0
    for {
        if helper.ValidNonce(lastNonce, nonce, lastHash) {
            log.Printf("nonce found %d", nonce)
            return nonce
        }

        nonce++
    }

    panic("something went wrong bro")
}
