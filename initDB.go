package main

import (
	"fmt"
	"github.com/wdnb/gene/api"
	"log"
)

//初始化数据库

func main() {
	udb:=api.CreateUserDB()
	err := udb.CreateMiniProgramUserBucket()
	if nil!=err {
		log.Panic("未知错误")
	}
	fmt.Println("小程序用户数据库创建成功")
}
