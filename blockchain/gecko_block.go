package blockchain

import (
	"fmt"
	"github.com/wdnb/gene/gecko"
	"github.com/wdnb/gene/wallet"
	"log"
)

//录入
//发送
//更新
//删除

type Print struct {
	Hash []byte
	Height int
	PrevBlockHash []byte
	Gecko []*gecko.Gecko
	Transactions  []*Transaction
}

func EntryGecko(from,nodeID string, mineNow bool,msg gecko.Gecko) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}

	bc := NewBlockchain(nodeID)
	//UTXOSet := UTXOSet{bc}
	//GECKOset := GeckoSet{bc}
	defer bc.db.Close()

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	w := wallets.GetWallet(from)

	//tx := NewUTXOTransaction(&w, to, amount, &UTXOSet)
	gtx :=NewGecko(&w,msg)//填充结构体
	//gecko获取方式分为转账和自己创建
	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		//fmt.Println(cbTx)
		//txs := []*Transaction{cbTx, tx}
		txs := []*Transaction{cbTx}
		geckos := []*gecko.Gecko{gtx}
		//fmt.Println(geckos)
		//fmt.Println(txs)
		//gecko.
		//fmt.Println(geckos)
		bc.MineBlock(txs,geckos)
		//更新gecko block
		//UTXOSet.Update(newBlock)
	} else {
		//sendTx(knownNodes[0], gtx)
		//send gtx
		//sendTx(knownNodes[0], gtx)
	}

	fmt.Println("Success!")
}
//
func PrintGecko(nodeID string)  []Print{
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	bci := bc.Iterator()
	//var
	p := []Print{}
	//g:=[]*gecko.Gecko{}
	//s:=[]byte
	for {
		block := bci.Next()
		p=append(p,Print{block.Hash,block.Height,block.PrevBlockHash,block.Gecko,	block.Transactions})
		//fmt.Printf("============ Block %x ============\n", block.Hash)
		//fmt.Printf("Height: %d\n", block.Height)
		//fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		//
		//for _,ge:=range block.Gecko{
		//	fmt.Println(ge)
		//	//g=append(g,ge)
		//}
		//return p
		//fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}

	}
	//fmt.Println(g)
	return p
}

//func SendGecko(from, to string, amount int, nodeID string, mineNow bool) {
//	if !wallet.ValidateAddress(from) {
//		log.Panic("ERROR: Sender address is not valid")
//	}
//	if !wallet.ValidateAddress(to) {
//		log.Panic("ERROR: Recipient address is not valid")
//	}
//
//	bc := NewBlockchain(nodeID)
//	UTXOSet := UTXOSet{bc}
//	GECKOset := GeckoSet{bc}
//	defer bc.db.Close()
//
//	wallets, err := wallet.NewWallets(nodeID)
//	if err != nil {
//		log.Panic(err)
//	}
//	w := wallets.GetWallet(from)
//
//	tx := NewUTXOTransaction(&w, to, amount, &UTXOSet)
//	gtx :=NewGecko(&w,&GECKOset)//填充结构体
//	if mineNow {
//		cbTx := NewCoinbaseTX(from, "")
//		//fmt.Println(cbTx)
//		//txs := []*Transaction{cbTx, tx}
//		//geckos := []*gecko.Gecko{gtx}
//		//fmt.Println(geckos)
//		//fmt.Println(txs)
//		//gecko.
//		//newBlock := bc.MineBlock(geckos,txs)
//		UTXOSet.Update(newBlock)
//	} else {
//		sendTx(knownNodes[0], tx)
//	}
//
//	fmt.Println("Success!")
//}
