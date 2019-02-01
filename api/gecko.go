package api

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/wdnb/gene/blockchain"
	"github.com/wdnb/gene/gecko"
	"net/http"
)





//func Test(c *gin.Context)  {
//	var msg gecko.BMI
//
//	if c.BindJSON(&msg) == nil {
//		fmt.Println(msg)
//		//message := c.Param("Sex")//获取url
//		c.JSON(http.StatusOK, gin.H{
//			"message":	msg.Height,
//		})
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"err_no":     "4001",
//			"message":    "request's params wrong!",
//			//"timestamp":  time.Now().Format(_timestamp_format),
//		})
//	}
//}

func GeckoList(c *gin.Context) {
	fmt.Println("showgecko")
	return
}

func LeopardCreate(c *gin.Context)  {
	var msg gecko.Gecko
	if c.BindJSON(&msg) != nil {
		response(c,http.StatusBadRequest,"request's params wrong!")
		return
	}
	if _,err := govalidator.ValidateStruct(msg); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	//getCurrentDirectory()
	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//println(path)
	//在线账户模式下根据id获取用户钱包
	//离线账户读取本地钱包
	blockchain.EntryGecko("19pUjZXFySwcYSkzy3o2E7hyKTmd6zTNnq","3000",true,msg)
	response(c,http.StatusOK,msg)
	return
}

func LeopardRetrieve(c *gin.Context)  {
	p:=blockchain.PrintGecko("3000")
	response(c,http.StatusOK,p)
}

func Leopard()  {

}

func BaseGeneLoad(c *gin.Context)  {

	response(c,http.StatusOK,"")
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
	//fmt.Println("ssss")
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
