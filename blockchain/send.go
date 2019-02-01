package blockchain

import (
	"fmt"
	"github.com/wdnb/blockchain-tutorial/gecko/wallet"
	"log"
)
func Send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	//GECKOset := GeckoSet{bc}
	defer bc.db.Close()

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	w := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&w, to, amount, &UTXOSet)
	//gtx :=NewGecko(&w,&GECKOset)//填充结构体
	if mineNow {
		//cbTx := NewCoinbaseTX(from, "")
		//fmt.Println(cbTx)
		//txs := []*Transaction{cbTx, tx}
		//geckos := []*gecko.Gecko{gtx}
		//fmt.Println(geckos)
		//fmt.Println(txs)
		//gecko.
		//newBlock := bc.MineBlock(geckos)
		//UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
