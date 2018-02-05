package miner

import "github.com/hoshsadiq/tescoin/block"

const CURRENT_BLOCK_REWARD = 12.5

func Mine(blockchain block.Blockchain, blockRewardReceiver string) *block.Block {
    lastBlock := blockchain.GetLastBlock()
    proof := lastBlock.ProofOfWork()

    transaction := block.NewTransaction("mined", blockRewardReceiver, 12.5)
    blockchain.AddTransaction(transaction)

    prevHash := lastBlock.GetHash()
    return blockchain.NewBlock(proof, prevHash)
}
