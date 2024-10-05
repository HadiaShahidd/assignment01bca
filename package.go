package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block structure to hold information about each block in the blockchain
type block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
	Timestamp    time.Time
}

// Blockchain slice to store all blocks
var Blockchain []block

// Function to calculate the hash for a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Function to create a new block and add it to the blockchain
func NewBlock(transaction string, nonce int) *block {
	var previousHash string
	if len(Blockchain) == 0 {
		previousHash = "" // Genesis block has no previous hash
	} else {
		previousHash = Blockchain[len(Blockchain)-1].CurrentHash // Get the hash of the last block
	}

	blockData := strings.Join([]string{transaction, strconv.Itoa(nonce), previousHash, time.Now().String()}, ":")
	blockHash := CalculateHash(blockData)

	newBlock := block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  blockHash,
		Timestamp:    time.Now(),
	}
	Blockchain = append(Blockchain, newBlock)
	return &newBlock
}

// Function to list all blocks in the blockchain
func ListBlocks() {
	for i, blk := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Transaction: %s\n", blk.Transaction)
		fmt.Printf("  Nonce: %d\n", blk.Nonce)
		fmt.Printf("  Previous Hash: %s\n", blk.PreviousHash)
		fmt.Printf("  Current Hash: %s\n", blk.CurrentHash)
		fmt.Printf("  Timestamp: %s\n", blk.Timestamp)
		fmt.Println("---------------------------")
	}
}

// Function to change the transaction of a block (for simplicity, change by block index)
func ChangeBlock(index int, newTransaction string) {
	if index < len(Blockchain) && index >= 0 {
		Blockchain[index].Transaction = newTransaction
		// Update the current hash after modifying the transaction
		blockData := strings.Join([]string{Blockchain[index].Transaction, strconv.Itoa(Blockchain[index].Nonce), Blockchain[index].PreviousHash, Blockchain[index].Timestamp.String()}, ":")
		Blockchain[index].CurrentHash = CalculateHash(blockData)
		fmt.Println("Block updated successfully!")
	} else {
		fmt.Println("Invalid block index")
	}
}

// Function to verify the integrity of the blockchain
func VerifyChain() {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].CurrentHash {
			fmt.Printf("Blockchain compromised at block %d!\n", i)
			return
		}
	}
	fmt.Println("Blockchain is valid.")
}
