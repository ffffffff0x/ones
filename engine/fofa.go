package engine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	ones "ones/mod"
	"os"
	"reflect"
)

// https://fofa.info/api
var resp2 []string

func TodoFofa() (string, []string) {

	byte1 := []byte(ones.Fofa)
	base64str := base64.StdEncoding.EncodeToString(byte1)
	//fmt.Println(base64str)

	fofaEmailValue := string(ones.Confs["fofa_email"])
	fofaEmailValue = fofaEmailValue[1 : len(fofaEmailValue)-1]
	//fmt.Println(fofaEmailValue)
	fofaKeyValue := string(ones.Confs["fofa_key"])
	fofaKeyValue = fofaKeyValue[1 : len(fofaKeyValue)-1]
	//fmt.Println(fofaKeyValue)

	url := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%d&fields=host", fofaEmailValue, fofaKeyValue, base64str, ones.Num)
	//fmt.Println(url)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		os.Exit(3)
	}

	var obj map[string]interface{}
	err = json.Unmarshal(resp, &obj)

	v := reflect.ValueOf(obj["results"])
	count := v.Len()
	for i := 0; i < count; i++ {
		resp2 = append(resp2, fmt.Sprint(v.Index(i)))
	}

	//fmt.Println(string(resp))
	return string(resp), resp2

}
