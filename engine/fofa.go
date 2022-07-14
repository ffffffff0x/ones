package engine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	ones "ones/mod"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// https://fofa.info/api
var resp2 []string
var fofaRecursion = 0

func TodoFofa() (string, []string) {
	fofaRecursion += 1
	if fofaRecursion% ones.Recursion == 0 { // 到达最大递归层数，直接退出
		return "", nil
	}

	byte1 := []byte(ones.Fofa)
	base64str := base64.StdEncoding.EncodeToString(byte1)
	//fmt.Println(base64str)

	fofaValue := ones.GetToken("fofa")
	fofaEmailValue := strings.Split(fofaValue, ":")[0]
	fofaKeyValue := strings.Split(fofaValue, ":")[1]

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
	hasError, _ := strconv.ParseBool(fmt.Sprint(obj["error"]))
	if hasError == false {
		v := reflect.ValueOf(obj["results"])
		count := v.Len()
		for i := 0; i < count; i++ {
			resp2 = append(resp2, fmt.Sprint(v.Index(i)))
		}
		return string(resp), resp2
	} else { // 有错误递归，但是要考虑到无限递归的情况
		return TodoFofa()
	}
	return "", nil
}
