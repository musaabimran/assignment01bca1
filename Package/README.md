# Sample Main for Testing
package main

import (
	"fmt"
	"strings"

	assignment01bca "github.com/musaabimran/assignment01bca1"
)

// Main function
func main() {

	// Creating Blockchain
	genesisBlock := assignment01bca.NewBlock("Genesis", 0, nil)
	Block1 := assignment01bca.NewBlock("Alice to Bob", 1, genesisBlock)
	Block2 := assignment01bca.NewBlock("Bob to Jack", 2, Block1)
	Block3 := assignment01bca.NewBlock("Jack to Alice", 3, Block2)

	// Printing the blockchain
	fmt.Printf("\n%s Blockchain %s\n\n", strings.Repeat("=", 33), strings.Repeat("=", 33))
	assignment01bca.ListBlocks(Block3)

	// Changing transaction of Block2
	assignment01bca.ChangeBlock(Block2, "Bob to Alice")

	// Checking if a transaction has been changed
	if assignment01bca.VerifyChain(Block2) {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
	fmt.Println()
}
