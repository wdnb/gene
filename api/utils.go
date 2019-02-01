package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func response(c *gin.Context,code int,message interface{})  {
	c.JSON(http.StatusOK, gin.H{
		"code":	code,
		"message":message,
	})
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