package blockchain

import (
	"encoding/hex"
	"fmt"
	"github.com/wdnb/gene/gecko"
	"github.com/wdnb/gene/wallet"
	"log"
)

//Transaction
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

func (bc *Blockchain)NewGecko(w *wallet.Wallet,msg *gecko.Gecko) *gecko.Gecko{
	from := fmt.Sprintf("%s", w.GetAddress())

	g:=bc.GetLastestBlock().Gecko[len(bc.GetLastestBlock().Gecko)-1]

	msg.Serial = g.Serial+1;

	txout:=NewTXOutput(1,from)
	geckos:=gecko.Gecko{msg.Gene,msg.Serial,msg.BMI,msg.Egg,msg.Name,msg.Sex,msg.Birth,msg.Death,txout.PubKeyHash}

	return  &geckos
}
