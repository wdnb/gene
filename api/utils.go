package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
//
//type Raw struct {
//	Code int         `json:"code"`
//	Message  []byte      `json:"message"`
//}

//var Raw map[string]interface{}

func response(c *gin.Context,code int,message interface{})  {
	c.JSON(http.StatusOK, gin.H{
		"code":	code,
		"message":message,
	})
}

func responseWithRaw(w http.ResponseWriter, code int, payload []byte)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_,err:=w.Write(payload)
	if nil!=err {
		log.Panic(err)
	}
}

//type Response struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}
//
//func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
//	response, _ := json.Marshal(payload)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//	w.Write(response)
//}