package gecko

import (
	"bytes"
	"encoding/gob"
	"log"
)

type BMI struct {
	Weight		int `valid:"range(0|1000)"`
	Height		int `valid:"range(0|1000)"`
	//Unique      string `valid:"required"`
}



type Egg struct {
	Father string
	Mother string
	Temperature []string
	//Birth    int64
	//Death    int64
	//Unique      string `valid:"required"`
}

//发情时间，交配时间，生产时间
type Reproduction struct {
	//Estrus int64
	Mating []string
	Baby []string
	//Unique      string `valid:"required"`
}

type Care struct {
	//Lasteat int64
	//Unique      string `valid:"required"`
}

type Gecko struct {
	BMI
	Egg
	//Reproduction
	//Care
	Name		string
	Sex			string `valid:"in(male|female|unknown)"`
	//Geng
	Birth   int64
	Death  int64
	PubKeyHash []byte
	//Unique      string `valid:"required"`
}



//type Print struct {
//	Hash []byte
//	Height int
//	PrevBlockHash []byte
//	Gecko []*Gecko
//	Transactions  []*blockchain.Transaction
//}

//func (geckos Gecko) String() string {
//	//gecko.Gecko{}
//	var lines []string
//	//fmt.Println(tx)
//	//fmt.Println("ss")
//
//	lines = append(lines, fmt.Sprintf("--- Transaction %v:", geckos.Name))
//	lines = append(lines,fmt.Sprintf("		Weight:		%v",geckos.BMI.Weight))
//	lines = append(lines,fmt.Sprintf("		Height:		%v",geckos.BMI.Height))
//	lines = append(lines,fmt.Sprintf("		Birth:		%v",geckos.Birth))
//	lines = append(lines,fmt.Sprintf("		Death:		%v",geckos.Death))
//	//for i, input := range geckos.BMI {
//
//		//lines = append(lines, fmt.Sprintf("		Input %d:", i))
//		//lines = append(lines, fmt.Sprintf("       TXID:      %x", input.Txid))
//		//lines = append(lines, fmt.Sprintf("       Out:       %d", input.Weight))
//		//lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
//		//lines = append(lines, fmt.Sprintf("       PubKey:    %x", input.PubKey))
//	//}
//	//
//	//for i, output := range tx.Vout {
//	//	lines = append(lines, fmt.Sprintf("     Output %d:", i))
//	//	lines = append(lines, fmt.Sprintf("       Value:  %d", output.Value))
//	//	lines = append(lines, fmt.Sprintf("       Script: %x", output.PubKeyHash))
//	//}
//
//
//
//	return strings.Join(lines, "\n")
//}

// Serialize returns a serialized Transaction
func (tx Gecko) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}