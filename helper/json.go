package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromReqBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(result)
	PanicIE(err)
}
func WriteToRespBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	PanicIE(err)
}
