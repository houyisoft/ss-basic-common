package utils

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type PostData struct {
	Type int //0: form, 1: body
	Ctx  *iris.Context
	Data *[]byte
	//body string
	//Items map[string]interface{}
}

func (_this *PostData) Get(field string) string {
	if _this.Type == 0 {
		return (*_this.Ctx).FormValue(field)
	}

	v, _, _, err := jsonparser.Get(*_this.Data, field)
	if err != nil { //再次尝试从 FormValues当中读取
		return ""
	}

	return string(v)
}



func GetParam(ctx *iris.Context, field string, isInt bool, args ...bool) interface{} {
	data := GetPostData(ctx)
	checkIsExists := (len(args) > 0) && args[0]
	var value interface{} = nil
	if _, ok := data.GetMap()[field]; ok {
		if isInt {
			value = int(data.GetInt(field))
		} else {
			value = data.Get(field)
		}
	} else if (*ctx).URLParamExists(field) {
		if isInt {
			value = int((*ctx).URLParamIntDefault(field, 0))
		} else {
			value = (*ctx).URLParam(field)
		}
	}
	if (value == nil) && (!checkIsExists) {
		if isInt {
			value = 0
		} else {
			value = ""
		}
	}
	return value
}


func GetPostData(ctx *iris.Context) PostData {
	contentType := (*ctx).GetHeader("content-type")
	postType := 0 //默认是form
	if strings.Index(contentType, "/json") > 0 {
		postType = 1
	}

	if postType == 0 {
		return PostData{
			Type: postType,
			Ctx:  ctx,
			Data: &[]byte{},
		}
	}

	body := (*ctx).Params().Get("PostBody")
	if strings.Compare(body, "") == 0 {
		data, _ := ioutil.ReadAll((*ctx).Request().Body)
		postBody := string(data)
		(*ctx).Params().SetImmutable("PostBody", postBody)
		return PostData{
			Type: postType,
			Ctx:  ctx,
			Data: &data,
		}
	}

	bytes := []byte(body)
	return PostData{
		Type: postType,
		Ctx:  ctx,
		Data: &bytes,
	}
}

func (_this *PostData) GetMap() map[string]string {
	if _this.Type == 1 { //json-body模式
		data := map[string]string{}
		tmp := map[string]interface{}{}
		json.Unmarshal(*_this.Data, &tmp)

		for k, v := range tmp {
			value := ""
			switch reflect.TypeOf(v).String() {
			case "int32":
				value = strconv.Itoa(int(v.(int32)))
			case "int64":
				value = strconv.Itoa(int(v.(int64)))
			case "float32":
				value = decimal.NewFromFloat32(v.(float32)).String()
			case "float64":
				value = decimal.NewFromFloat(v.(float64)).String()
			case "bool":
				value = "true"
				if v.(bool) {
					value = "false"
				}
			case "string":
				value = v.(string)
			default:
				if strings.HasPrefix(reflect.TypeOf(v).String(), "map[") {
					bs, err := json.Marshal(v)
					if err == nil {
						value = string(bs)
					}
				}
			}
			data[k] = value
		}
		return data
	}

	//form模式
	values := (*_this.Ctx).FormValues()
	data := map[string]string{}
	for k, v := range values {
		data[k] = strings.Join(v, ",")
	}

	return data
}

func (_this *PostData) GetInt(field string) int64 {
	if _this.Type == 0 {
		value := (*_this.Ctx).FormValue(field)
		if strings.Compare(value, "") == 0 {
			return 0
		}

		if num, err := strconv.ParseInt(value, 10, 64); err == nil {
			return num
		}
		return 0
	}

	v, _, _, err := jsonparser.Get(*_this.Data, field)
	if err != nil {
		return 0
	}
	if num, numErr := strconv.Atoi(string(v)); numErr == nil {
		return int64(num)
	}

	return 0
}
