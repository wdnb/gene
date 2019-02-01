package debug

import (
	"github.com/wdnb/blockchain-tutorial/gecko/wallet"
	"testing"
)

func TestWallet(t *testing.T) {
	//err := godotenv.Load()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//nodeID := os.Getenv("NODE_ID")
	//if nodeID == "" {
	//	fmt.Printf("NODE_ID env. var is not set!")
	//	os.Exit(1)
	//}
	wallet.CreateWallet("3000")
}
