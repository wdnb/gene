package config

import (
	"fmt"
	"github.com/wdnb/gene/gecko"
	"time"
)

type Cgecko gecko.Gecko

type BMI struct {
	Weight		int
	Height		int
}



type Egg struct {
	Father string
	Mother string
	Temperature []string
	Birth     time.Time
	Death     time.Time
	Hash      string
}

//发情时间，交配时间，生产时间
type Reproduction struct {
	Estrus time.Time
	Mating []string
	Baby []string
}

type Care struct {
	Lasteat time.Time
}

type Gecko struct {
	BMI
	From Egg
	Reproduction
	Care
	name		string
	Sex			string
	Geng		[]string
	Birth     time.Time
	Death     time.Time
	Hash      string
}

func testgecko()  {
	fmt.Println("ss")
}