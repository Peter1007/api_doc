package modules

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DocPath struct {
	Name  interface{}
	Url   string
	MS    interface{}
	Child []DocPath
}

type DocModule struct {
}

//遍历目录及子目录
func (docModule *DocModule) VistorFiles(dirPth string, suffix string) (files []DocPath, err error) {

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToLower(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		var childFile DocPath
		name := fi.Name()
		if fi.IsDir() {
			log.Println(name)
			childFile.Name = name
			childFile.Child, _ = docModule.VistorFiles(dirPth+PthSep+name, suffix)

		} else if strings.HasSuffix(strings.ToLower(name), suffix) { //匹配文件
			full_file_name := dirPth + PthSep + name
			file, err := os.Open(full_file_name)
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			b, err := ioutil.ReadAll(file)
			if err != nil {
				log.Panic(err)
			}

			var doc map[string]interface{}
			json.Unmarshal(b, &doc)

			childFile.Url = base64.StdEncoding.EncodeToString([]byte(full_file_name))
			childFile.MS, _ = doc["MS"]
			childFile.Name, _ = doc["service_name"]
		}

		files = append(files, childFile)
	}
	return
}

//api详细页面
func (docModule *DocModule) Detail(fileName string) (doc Doc, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	json.Unmarshal(b, &doc)

	file_info, _ := os.Stat(fileName)
	doc.ModTime = file_info.ModTime().Format("2006-01-02 15:04:05")

	b, err = json.Marshal(doc.Request.RequestExample)
	if err != nil {
		return doc, err
	}
	doc.Request.RequestExample = string(b)

	b, err = json.Marshal(doc.Response.ResponseExample)
	if err != nil {
		return doc, err
	}
	doc.Response.ResponseExample = string(b)

	switch doc.MS {
	case "SOA_AMS":
		doc.BaseMs = "ayi"
	case "SOA_BMS":
		doc.BaseMs = "base"
	case "SOA_CMS":
		doc.BaseMs = "coupon"
	case "SOA_DMS":
		doc.BaseMs = "data"
	case "SOA_MEMS":
		doc.BaseMs = "member"
	case "SOA_MOMS":
		doc.BaseMs = "money"
	case "SOA_OMS":
		doc.BaseMs = "order"
	}

	return
}

//js跨域请求
func (docModule *DocModule) RequestRemote(uri string, req_body []byte) (result []byte, err error) {
	//uri:beta_base/banner/get_list_by_platform
	path := strings.SplitN(uri, "/", 3)
	base_path := strings.Split(path[1], "_")

	var host string
	switch base_path[0] {
	case "beta":
		host = "http://139.196.111.219/" + base_path[1]
	case "gamma":
		host = "http://139.196.111.185/" + base_path[1]
	default:
		host = "http://" + base_path[1] + ".yunjiazheng.com"
	}
	remote_url := host + "/" + path[2]

	body := bytes.NewBuffer(req_body)
	res, err := http.Post(remote_url, "application/json;charset=utf-8", body)
	if err != nil {
		result = []byte("{\"meta\":\"\",\"status\":500,\"data\":\"远程接口调用异常\"}")
		err = nil
		return
	}
	result, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	return
}
