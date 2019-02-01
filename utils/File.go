package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BaseGene struct {
//	id: int,
//	"tree": {
//	"id": 1,
//	"desc": "预设基因",
//	"level": "level1",
//	"logo": null,
//	"popular": null,
//	"nodes": [{
//	"id": 1,
//	"tree": {
//	"id": 1,
//	"desc": "牛奶粉",
//	"level": "level2",
//	"logo": "",
//	"popular": null,
//	"nodes": [{
//	"id": 157,
//	"tree": {
//	"id": 157,
//	"desc": "超级雪花恶魔白酒图片 设置url点击进入详细说明",
//	"desc2": "SSDB",
//	"logo": ""
//}
//},
//]
//}
//}]
//}

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