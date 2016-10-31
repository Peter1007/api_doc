package controllers

import (
	"net/http"
)

func Index(resp http.ResponseWriter, req *http.Request) string {
	return "hello go!"
	//	body, _ := ioutil.ReadAll(req.Body)
	//	fmt.Println(string(body))
	//	var b = []byte(`dddd`)
	//	resp.Write(b)
}
