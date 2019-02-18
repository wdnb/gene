package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BaseGene struct {
	s map[string]map[string]bool
}

type nodes struct {
	s map[string]map[string]bool
}
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) ReadInJson(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadInFile(filename string) []byte{
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}