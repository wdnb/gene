package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/wdnb/gene/blockchain"
	"github.com/wdnb/gene/gecko"
	"github.com/wdnb/gene/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func LeopardCreate(c *gin.Context)  {
	var msg gecko.Gecko
	//bindjson不能完整解析gene呀不知道为啥吖 只能自己动手了呀
	gene,_:=(ioutil.ReadAll(c.Request.Body))

	if err := json.Unmarshal(gene, &msg); err != nil {
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	if _,err := govalidator.ValidateStruct(msg); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	//在线账户模式下根据id获取用户钱包
	//离线账户读取本地钱包
	uname := GetTokenUserName(c)
	udb := NewUserDB()
	address,err:=udb.GetAddress(uname)
	if nil!=err {
		log.Panic(err)
	}
	//address将用来签名PubKeyHash
	//钱包文件名格式 wallet_uname
	//需要支持批量添加gecko
	blockchain.EntryGecko(string(address),uname,true,&msg)
	response(c,http.StatusOK,msg)
	return
}

func LeopardRetrieve(c *gin.Context)  {
	uname := GetTokenUserName(c)
	t:=blockchain.GeckoList(uname)
	response(c,http.StatusOK,t)
}

func LeopardPrintChain(c *gin.Context)  {
	uname := GetTokenUserName(c)
	p:=blockchain.PrintGecko(uname)
	response(c,http.StatusOK,p)
}

func BaseGeneLoad(c *gin.Context)  {
	var gene  []byte
	gene = utils.ReadInFile("./db/gene.json")
	responseWithRaw(c.Writer,http.StatusOK,gene)
}

func BMIRetrieve(c *gin.Context)  {
	return
}

func BMIUpdate(c *gin.Context)  {
	return
}

func BMIDelete(c *gin.Context)  {
	return
}

func CreateEgg()  {
	return
}

func RetrieveEgg()  {
	return
}

func UpdateEgg()  {
	return
}

func DeleteEgg()  {
	return
}

func BMICreate(c *gin.Context)  {
	var msg gecko.BMI
	if c.BindJSON(&msg) != nil {
		response(c,http.StatusBadRequest,"request's params wrong!")
		return
	}
	if _,err := govalidator.ValidateStruct(msg); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	response(c,http.StatusOK,"ok")
	return
}
