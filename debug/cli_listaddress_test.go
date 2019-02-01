package debug

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/wdnb/blockchain-tutorial/gecko/wallet"
	"log"
	"os"
	"testing"
)

func TestListAddresses(t *testing.T){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env. var is not set!")
		os.Exit(1)
	}
	ListAddresses(nodeID)
}

func ListAddresses(nodeID string) {
	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
