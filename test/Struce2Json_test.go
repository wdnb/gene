package test_test

import (
	"fmt"
	"testing"
"encoding/json"
)

type Person struct {
	Name string
	Age int
	Gene []*Gene
}

//var sb

type Gene struct {
	 bool
	//SDB bool
	//NM bool
	//DB bool
	//RS bool
	//SB bool
}

var (
	txt = []byte(`{"gene":{"SSDB":false,"SDB":true,"NM":false,"DB":false,"rs":false,"SB":false}}`)
)

func TestStruct2Json(t *testing.T) {
	var m map[string]map[string]bool
	if err := json.Unmarshal(txt, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m["gene"]["SB"])
}
