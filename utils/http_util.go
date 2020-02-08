package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type helper struct{}

//单例helper对象
var Helper *helper

func init() {
	Helper = new(helper)
}

//路径拼接
func (this *helper) HelperGetUrl(urlPath string, data map[string]string, isEncode bool) string {
	//urlPath = strings.Replace(urlPath,`/`,`\/`,-1)
	urlPath += "?"
	for k, v := range data {
		if isEncode {
			v = url.QueryEscape(v)
		}
		urlPath += k + "=" + v + "&"
	}
	urlPath = strings.TrimRight(urlPath, "&")

	return urlPath
}

//路径拼接,http表单处理
func (this *helper) HelperGetUrlHttps(urlPath string, data map[string]string, isEncode bool, isHttp bool) string {

	if isHttp {
		urlPath = strings.Replace(urlPath, "https://", "http://", 1)
	}
	urlPath += "?"
	for k, v := range data {
		if isEncode {
			v = url.QueryEscape(v)
		}
		urlPath += k + "=" + v + "&"
	}
	urlPath = strings.TrimRight(urlPath, "&")

	return urlPath
}

func (this *helper) HelperCurlGetHeader(urlPath string, timeout time.Duration, header map[string]string) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", urlPath, strings.NewReader(""))
	if err != nil {
		return []byte("")
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlGet(urlPath string, timeout time.Duration) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", urlPath, strings.NewReader(""))
	if err != nil {
		return []byte("")
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlPostRSA(urlPath string, httpBuildQuery string, timeout time.Duration) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(httpBuildQuery))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlPostHeader(urlPath string, data map[string]string, timeout time.Duration, header map[string]string) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	httpBuildQuery := ""
	for k, v := range data {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(data) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(httpBuildQuery))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlPost(urlPath string, data map[string]string, timeout time.Duration) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	httpBuildQuery := ""
	for k, v := range data {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(data) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(httpBuildQuery))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlPostJson(urlPath string, data string, timeout time.Duration) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (this *helper) HelperCurlPostJsonHeader(urlPath string, data string, timeout time.Duration, header map[string]string) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
