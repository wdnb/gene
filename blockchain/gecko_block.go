package blockchain

import (
	"fmt"
	"github.com/wdnb/gene/gecko"
	"github.com/wdnb/gene/wallet"
	"log"
)

//block

type Print struct {
	Hash []byte
	Height int
	PrevBlockHash []byte
	Gecko []*gecko.Gecko
	Transactions  []*Transaction
}

type List struct {
	Node []Node
}

type Node struct {
	title string
	Base []Base
}

type Base struct {
	Gene map[string]interface{}
	Serial int
	Name		string
}

//type title struct {
//	Group string
//}

func EntryGecko(from,nodeID string, mineNow bool,msg *gecko.Gecko) {
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
	gtx :=bc.NewGecko(&w,msg)//填充结构体
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
//小壁虎列表
func GeckoList(nodeID string)  interface{}{
	p:= PrintGecko(nodeID)
	var tmp []*gecko.Gecko
	var tmpbb []Base
	var tmpll []Node
	tmpG := make(map[string]bool)//用于记录group字段
	//tmpL := make(map[string]map[string]string)
	//s := []tmpL

	for _,value:=range p {
		for _,g:=range value.Gecko {
			//已有group字段 添加base
			//无group字段 添加List 注册group
			_, ok := tmpG[g.Group]
			if ok{
				tmpbb = append(tmpbb,Base{g.Gene,g.Serial,g.Name})
				//tmpll = append(tmpll,List{g.Name,tmpbb})
				fmt.Println("base")
				//tmpL = ["b"]["f"]["s"]
			}else {
				tmpll = append(tmpll,Node{g.Name,tmpbb})
				//List{g.Group}
				fmt.Println("list")
			}
			tmp = append(tmp,g)
			//fmt.Println(g.Group)
			tmpG[g.Group] = true

		}
	}

	//fmt.Println(tmpbb)
	//fmt.Println(tmpll)
	//fmt.Println(tmpG)
	return tmp
}

//填充list结构体
func ListInit(g interface{})  interface{}{
	return nil
}
