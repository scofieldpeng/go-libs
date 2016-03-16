//Copyright 2015 Author Scofield Peng
//curl 包用来方便快捷迅速简单粗暴的进行http请求
// 警告： 此版本不是很满意，会进行重写，使用需谨慎！！！！
//
// 使用方法：
//    curlIns := curl.NewCurl()
//    curlIns.SetUrl() //设置要请求的url
//    curlIns.SetParams() // 设置要发送的参数
//    curlIns.SetMethod() // 设置请求方式（目前支持post,get,postform,put,delete)
//    curlObj := curlIns.Send() // 发送请求
//    var res map[string]string
//    curlObj.Parse("json",&res) // 将请求回来的参数回写到res中，第一个参数为返回数据的格式，目前支持json和xml
//    resStr := curlObj.Parse2Str() // 返回string格式的结果
package curl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//curl
type curl struct {
	url     string                 //请求的url地址
	method  string                 //请求的方式
	params  map[string]interface{} //请求的参数值
	headers map[string]string      //请求的header值
	result  []byte                 //请求的原始结果集
	resp    http.Response          //请求返回的response对象
}

// NewCurl 返回一个NewCurl结构体对象
func NewCurl() *curl {
	return &curl{
		url:     "",
		method:  "get",
		params:  make(map[string]interface{}),
		headers: make(map[string]string),
	}
}

//SetUrl 设置请求的地址
func (this *curl) SetUrl(url string) *curl {
	this.url = url
	return this
}

//Url 返回本次请求的url地址
func (this *curl) Url() string {
	return this.url
}

//SetMethod 设置请求方式,支持"post","postform","get"
func (this *curl) SetMethod(method string) *curl {
	if method != "get" && method != "post" && method != "postform" && method != "put" {
		method = "get"
	}
	this.method = method
	return this
}

//Method 设置请求的方式
func (this *curl) Method() string {
	return this.method
}

//SetParams 设置请求参数值
func (this *curl) SetParams(params map[string]interface{}) *curl {
	this.params = params
	return this
}

//Params 获取请求参数
func (this *curl) Params() map[string]interface{} {
	return this.params
}

//SetHeader 设置header请求头
func (this *curl) SetHeader(key, value string) *curl {
	this.headers[key] = value
	return this
}

//SetHeaders 批量设置header请求头
func (this *curl) SetHeaders(headers map[string]string) *curl {
	for k, v := range headers {
		this.SetHeader(k, v)
	}
	return this
}

//获取请求头
func (this *curl) Header() map[string]string {
	return this.headers
}

//setResult 设置结果集
func (this *curl) setResult(result []byte) {
	this.result = result
}

//getResult 获取原始结果集内容
func (this *curl) getResult() []byte {
	return this.result
}

//Send 发送请求,从微信服务器获取数据
func (this *curl) Send() error {
	requestUrl := this.Url()
	var resp *http.Response
	var err error
	methodMap := map[string]string{
		"post":     "POST",
		"postform": "POST",
		"get":      "GET",
		"put":      "PUT",
		"delete":   "DELETE",
	}
	params := bytes.NewBuffer([]byte(""))
	switch this.Method() {
	case "get":
		tmp := url.Values{}
		for k, v := range this.Params() {
			switch t := v.(type) {
			case string:
				tmp.Set(k, fmt.Sprintf("%s", t))
			case int:
				tmp.Set(k, strconv.Itoa(t))
			case int32:
				tmp.Set(k, fmt.Sprintf("%d", t))
			case int64:
				tmp.Set(k, strconv.FormatInt(t, 10))
			case float64:
				tmp.Set(k, strconv.FormatFloat(t, 'f', -1, 64))
			case float32:
				tmp.Set(k, fmt.Sprintf("%g", t))
			case map[string]interface{}, []interface{}:
				if res, err := json.Marshal(t); err != nil {
					log.Println(err)
				} else {
					tmp.Set(k, string(res))
				}
			default:
				log.Printf("无法解析该类型参数!%s:%#v", k, t)
			}
		}
		requestUrl = requestUrl + "?" + tmp.Encode()
	case "post", "put", "delete":
		tmpParams, err := json.Marshal(this.Params())
		if err != nil {
			tmpParams = []byte("")
		}

		params.Write(tmpParams)
	case "postform":
		tmp := url.Values{}
		for k, v := range this.Params() {
			switch t := v.(type) {
			case string:
				tmp.Set(k, fmt.Sprintf("%s", t))
			case int:
				tmp.Set(k, strconv.Itoa(t))
			case int32:
				tmp.Set(k, fmt.Sprintf("%d", t))
			case int64:
				tmp.Set(k, strconv.FormatInt(t, 10))
			case float64:
				tmp.Set(k, strconv.FormatFloat(t, 'f', -1, 64))
			case float32:
				tmp.Set(k, fmt.Sprintf("%g", t))
			case map[string]interface{}, []interface{}:
				if res, err := json.Marshal(t); err != nil {
					log.Println(err)
				} else {
					tmp.Set(k, string(res))
				}
			default:
				log.Printf("无法解析该类型参数!%s:%#v", k, t)
			}
		}
		params.WriteString(tmp.Encode())
		this.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf8")
	}

	req, err := http.NewRequest(methodMap[this.Method()], requestUrl, params)

	for k, v := range this.Header() {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		log.Println("curl失败，请求url:", this.Url(), "请求参数", this.Params(), "请求头部：", this.Header())
		return err
	}

	this.resp = *resp
	if bufBytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		this.setResult(bufBytes)
	}

//	log.Println("请求url:", this.Url(), "请求参数", this.Params(), "请求头部：", this.Header())

	return nil
}

//Parse 解析结果集
func (this *curl) Parse(resType string, output interface{}) error {
	switch resType {
	case "string":
		output = string(this.getResult())
		return nil
	case "json":
		return json.Unmarshal(this.getResult(), output)
	case "xml":
		return xml.Unmarshal(this.getResult(), output)
	default:
		return errors.New("请传入正确的resType值,目前只支持string,json,xml")
	}

	return nil
}

//将结果解析成string
func (this *curl) Parse2Str() string {
	return string(this.getResult())
}

//获取请求的结果集
func (this curl) GetResp() http.Response {
	return this.resp
}
