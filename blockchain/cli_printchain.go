package blockchain

import (
	"fmt"
	"strconv"
)

func PrintChain(nodeID string) {
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		//for _, tx := range block.Transactions {
		//	fmt.Println(tx)
		//}
		for _,ge:=range block.Gecko{
			fmt.Println(ge)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
