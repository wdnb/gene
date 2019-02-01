package config

import (
	"github.com/wdnb/gene/gecko"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

//检测数据格式
func main() {
	gecko := gecko.Gecko{}
	//var gecko gecko.Gecko
	JsonParse := NewJsonStruct()
	//v := Config{}
	//下面使用的是相对路径
	JsonParse.Load("./config/dataVerification.json", &gecko)
	fmt.Println(gecko)
	//fmt.Println(gecko.Mongo.MongoDb)
}
//如果数据不为空就检测数据是否异常
func checkRange(){

}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
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
