package api

import (
	"github.com/asaskevich/govalidator"
)
func verificationGecko(msg interface{})  error{
	////自定义tag验证函数
	//govalidator.TagMap["machine_id"] = govalidator.Validator(func(str string) bool {
	//	return strings.HasPrefix(str, "IX")
	//})

	if _, err := govalidator.ValidateStruct(msg); err != nil {
		return  err
	}
	return nil
}
