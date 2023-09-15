// Name: Musaab Imran
// Student ID: 20i-1794
// Date Created: 2022-09-12

package assignment01bca1

// Importing packages
import (
	"crypto/sha256"
	"fmt"
)

// Single block struct that has: transaction, nonce, previous hash, current hash, previous block reference
type block struct {
	transaction   string
	nonce         int
	previousHash  string
	currentHash   string
	previousBlock *block
}

// The CreateHash() method will provide the block Hash value.
// Returns the string after calculating the Hash
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}

// The NewBlock() method will create a new block and assign it the values and return the block
// Previous Block parameter was added so that we can keep track of the previous block
// making a chain of blocks (link list)
func NewBlock(transaction string, nonce int, previousHash string, previousBlock *block) *block {
	b := new(block)
	b.transaction = transaction
	b.nonce = nonce
	b.previousHash = previousHash

	// Concatenating the transaction, nonce and previous hash to create a string for hash calculation
	stringToHash := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	b.currentHash = CalculateHash(stringToHash)
	b.previousBlock = previousBlock
	return b
}

// A method to print all the blocks in a nice format showing block data such as transaction, nonce, previous hash, current block hash
// Block parameter was added to print the blockchain from the genesis block
func ListBlocks(b *block) {
	blockNumber := 1
	for b != nil {
		fmt.Printf("Block Number: %d\n", blockNumber)
		fmt.Printf("Transaction: %s\n", b.transaction)
		fmt.Printf("Nonce: %d\n", b.nonce)
		fmt.Printf("Previous Hash: %s\n", b.previousHash)
		fmt.Printf("Current Hash: %s\n\n", b.currentHash)
		b = b.previousBlock
		blockNumber++
	}
}

// Function to change block transaction of the given block ref
// Block ref and transaction parameters were added to change the transaction value of that block
func ChangeBlock(b *block, transaction string) {
	if b != nil {
		b.transaction = transaction
	} else {
		fmt.Println("Block not created!!!")
	}
}

// Function to verify blockchain in case any changes are made.
// Block parameter was added to verify whether that block is valid or has been changed
// The function returns true if the blockchain is valid and false if it is not valid (has been changed)
func VerifyChain(chain *block) bool {
	for chain != nil {
		stringToHash := fmt.Sprintf("%s%d%s", chain.transaction, chain.nonce, chain.previousHash)
		Hash := CalculateHash(stringToHash)
		if Hash != chain.currentHash {
			return false
		}
		chain = chain.previousBlock
	}
	return true
}
