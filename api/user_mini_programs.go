package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type User struct {
	UserName string    `json:"username" valid:"required~请输入用户名,uname~用户名可以为 英文字符区分大小写 数字 下划线 长度3到28之间"`
	Password string `json:"password" valid:"required~请输入密码,upass~密码长度为8到16之间"`
}
//微信作为可信端 不设置密码
type UserName struct {
	UserName string    `json:"username" valid:"required~请输入用户名,uname~用户名可以为 英文字符区分大小写 数字 下划线 长度3到28之间"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func custom_valid()  {
	govalidator.TagMap["uname"] = govalidator.Validator(func(str string) bool {
		s:=len(str)
		if s>28||s<3{
			return false
		}
		if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", str); ok {
			return true
		}
		return false
	})

	govalidator.TagMap["upass"] = govalidator.Validator(func(str string) bool {
		s:=len(str)
		if s>16||s<8{
			return false
		}
		return true
	})
}

func Register(c *gin.Context) {
	var user UserName
	if c.BindJSON(&user) != nil {
		response(c,http.StatusBadRequest,"request's params wrong!")
		return
	}
	custom_valid()
	if _,err := govalidator.ValidateStruct(user); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	udb:=NewUserDB()
	err := udb.IsExist(user)
	defer udb.db.Close()
	if err == nil {
		response(c, http.StatusForbidden, "the user was exist")
		return
	}
	//用户注册
	err = udb.Insert(user)
	if err != nil {
		response(c, http.StatusInternalServerError, "internal error")
		return
	}
	response(c,http.StatusOK,"注册成功")
	//c.Handler()
	return
}

func Login(c *gin.Context) {
	var user UserName
	if c.BindJSON(&user) != nil {
		response(c,http.StatusBadRequest,"request's params wrong!")
		return
	}
	custom_valid()
	if _,err := govalidator.ValidateStruct(user); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}
	//用户认证
	//exist := IsExist(db, collection, bson.M{"username": user.UserName})
	udb := NewUserDB()
	exist := udb.IsExist(user)
	defer udb.db.Close()
	if exist==nil {
		token, _ := GenerateToken(&user)
		response(c,http.StatusOK,JwtToken{Token: token})
	} else {
		response(c, http.StatusForbidden, "the user not exist")
	}
	return
}

//查询用户信息 返回用户信息
func Inquire(c *gin.Context) {
	var openid UserName
	if c.BindJSON(&openid) != nil {
		response(c,http.StatusBadRequest,"request's params wrong!")
		return
	}
	custom_valid()
	if _,err := govalidator.ValidateStruct(openid); err!=nil{
		response(c,http.StatusBadRequest,err.Error())
		return
	}

	udb := NewUserDB()
	err := udb.IsExist(openid)
	udb.db.Close()

	if err != nil {
		response(c, http.StatusForbidden, "the user not exist")
		return
	}
	response(c,http.StatusOK,"用户已注册")
	return
}
