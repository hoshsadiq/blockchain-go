package block

import (
    "time"
    "github.com/hoshsadiq/tescoin/node"
    "github.com/hoshsadiq/tescoin/helper"
)

type Blockchain struct {
    currentTransactions []*Transaction
    blocks              []*Block
    nodes               []*node.Node
}

func NewBlockchain() *Blockchain {
    blockchain := Blockchain{
        currentTransactions: []*Transaction{},
        blocks:              []*Block{},
        nodes:               []*node.Node{},
    }

    // genesis block!
    blockchain.NewBlock(100, "__genesis")

    return &blockchain
}

func (blockchain *Blockchain) NewBlock(proof int, previousHash string) *Block {
    if previousHash == "" {
        panic("previous hash cannot be empty")
    }

    block := NewBlock(len(blockchain.blocks), time.Now(), blockchain.currentTransactions, proof, previousHash)

    blockchain.currentTransactions = []*Transaction{}
    blockchain.AddBlock(block)

    return block
}

func (blockchain *Blockchain) AddBlock(block *Block) {
    blockchain.blocks = append(blockchain.blocks, block)
}

func (blockchain *Blockchain) AddTransaction(tx *Transaction) int {
    blockchain.currentTransactions = append(blockchain.currentTransactions, tx)

    return blockchain.GetLastBlock().GetIndex() + 1
}

func (blockchain *Blockchain) GetLastBlock() *Block {
    return blockchain.blocks[len(blockchain.blocks)-1]
}

func (blockchain *Blockchain) RegisterNode(node *node.Node) {
    blockchain.nodes = append(blockchain.nodes, node)
}

func (blockchain *Blockchain) Consensus() (replaced bool) {
    var newBlocks []*Block
    neighbours := blockchain.nodes

    maxLength := len(blockchain.blocks)

    for _, n := range neighbours {
        response := helper.GetUrl(n.GetChainUrl())
        if response != nil {
            length := response["length"].(int)
            var blocks []*Block
            helper.ConvertInterface(response["blocks"], &blocks)

            if length > maxLength && IsValidChain(blocks) {
                maxLength = length
                newBlocks = blocks
            }
        }
    }

    if newBlocks != nil {
        blockchain.blocks = newBlocks
        return true
    }

    return false
}

func (blockchain *Blockchain) GetBlocks() []*Block {
    return blockchain.blocks
}

func (blockchain *Blockchain) GetNodes() []*node.Node {
    return blockchain.nodes
}

func IsValidChain(blocks []*Block) bool {
    var lastBlock *Block
    for idx, blk := range blocks {
        if idx == 0 {
            lastBlock = blk
            continue
        }

        if blk.GetPreviousHash() != lastBlock.GetHash() {
            return false
        }

        if !helper.ValidProof(lastBlock.GetProof(), blk.GetProof(), lastBlock.GetPreviousHash()) {
            return false
        }

        lastBlock = blk
        idx++
    }

    return true
}
