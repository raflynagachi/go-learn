package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	PanicIfError(err)
	PanicIfError(json.Unmarshal([]byte(body), x))
}
