package blockchain

import (
	"fmt"
	"github.com/wdnb/blockchain-tutorial/gecko/wallet"
	"log"
)

func GetGec(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewBlockchain(nodeID)
	GeckoSet := GeckoSet{bc}
	defer bc.db.Close()

	//balance := 0
	pubKeyHash := wallet.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	Geckos := GeckoSet.FindGecko(pubKeyHash)
	fmt.Println(Geckos)
	//
	//for _, out := range Geckos {
	//	balance += out.Value
	//}

	//fmt.Printf("Balance of '%s': %d\n", address, balance)
}

