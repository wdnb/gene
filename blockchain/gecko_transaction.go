package blockchain

import (
	"encoding/hex"
	"fmt"
	"github.com/wdnb/blockchain-tutorial/gecko/gecko"
	"github.com/wdnb/blockchain-tutorial/gecko/wallet"
	"log"
)

func NewGeckoTransaction(w *wallet.Wallet, to string, amount int, UTXOSet *UTXOSet) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	pubKeyHash := wallet.HashPubKey(w.PublicKey)
	acc, validOutputs := UTXOSet.FindSpendableOutputs(pubKeyHash, amount)
	//fmt.Println(pubKeyHash)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, nil, w.PublicKey}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	from := fmt.Sprintf("%s", w.GetAddress())
	outputs = append(outputs, *NewTXOutput(amount, to))
	if acc > amount {
		outputs = append(outputs, *NewTXOutput(acc-amount, from)) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()
	UTXOSet.Blockchain.SignTransaction(&tx, w.PrivateKey)

	return &tx
}

func NewGecko(w *wallet.Wallet,msg gecko.Gecko) *gecko.Gecko{
	from := fmt.Sprintf("%s", w.GetAddress())
	fmt.Println(from)
	//bmi:=gecko.BMI{}
	//egg:=gecko.Egg{}
	//reproduction:=gecko.Reproduction{time.Now().Unix(),}
	//care:=gecko.Care{time.Now().Unix()}
	//name:="First generation"
	//sex:="man"
	//geng:=[]{"s","f"}
	//birth:=time.Now()
	//var death int64
	//death:=time.Now()
	//fmt.Println(death)
	txout:=NewTXOutput(1,from)
	geckos:=gecko.Gecko{msg.BMI,msg.Egg,msg.Name,msg.Sex,msg.Birth,msg.Death,txout.PubKeyHash}

	return  &geckos
}

//func (tx Transaction) String() string {
//	var lines []string
//	//fmt.Println(tx)
//
//	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))
//
//	for i, input := range tx.Vin {
//
//		lines = append(lines, fmt.Sprintf("     Input %d:", i))
//		lines = append(lines, fmt.Sprintf("       TXID:      %x", input.Txid))
//		lines = append(lines, fmt.Sprintf("       Out:       %d", input.Vout))
//		lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
//		lines = append(lines, fmt.Sprintf("       PubKey:    %x", input.PubKey))
//	}
//
//	for i, output := range tx.Vout {
//		lines = append(lines, fmt.Sprintf("     Output %d:", i))
//		lines = append(lines, fmt.Sprintf("       Value:  %d", output.Value))
//		lines = append(lines, fmt.Sprintf("       Script: %x", output.PubKeyHash))
//	}
//
//	return strings.Join(lines, "\n")
//}
