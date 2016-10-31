package controllers

import (
	"code.oa.com/sean/api-doc/modules"
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type DocsController struct {
	resp   http.ResponseWriter
	req    *http.Request
	module *modules.DocModule
}

func (docs *DocsController) Index(resp http.ResponseWriter, req *http.Request) {
	doc, _ := docs.module.VistorFiles("document\\api_description", ".json")
	//	doc, _ := docs.module.VistorFiles("D:\\xhj\\project\\new\\api_doc_local\\api_description", ".json")

	t, err := template.ParseFiles("views/index.tmpl")
	if err != nil {
		log.Panic(err)
	}

	t.ExecuteTemplate(resp, "index", doc)
}

func (docs *DocsController) Detail(resp http.ResponseWriter, req *http.Request) {
	name_encode, _ := url.QueryUnescape(req.URL.RawQuery)
	name, _ := base64.StdEncoding.DecodeString(name_encode)

	doc, _ := docs.module.Detail(string(name))

	t, err := template.ParseFiles("views/header.tmpl", "views/body.tmpl", "views/footer.tmpl")
	if err != nil {
		log.Panic(err)
	}

	t.ExecuteTemplate(resp, "body", doc)
}

//js跨域请求
func (docs *DocsController) RequestRemote(req *http.Request) string {
	uri := req.RequestURI
	req_body, _ := ioutil.ReadAll(req.Body)

	result, err := docs.module.RequestRemote(uri, req_body)
	if err != nil {
		log.Panic(err)
	}

	return string(result)
}
