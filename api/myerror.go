package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func DataVerification(msg interface{})  error{
	if _,err := govalidator.ValidateStruct(msg); err!=nil{
		return err
	}
	return nil
}

func ErrorRepsonse(code int,msg string) interface{} {
	data := map[string]interface{}{
		"code":code,
		"msg":msg,
	}
	return data
}

