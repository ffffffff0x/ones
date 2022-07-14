package engine

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	ones "ones/mod"
	"os"
	"reflect"
)

var chaosRecursion = 0
// https://github.com/projectdiscovery/chaos-client/blob/master/pkg/chaos/chaos.go

func TodoChaos() (string, []string) {
	chaosRecursion += 1
	if chaosRecursion % ones.Recursion == 0 {
		return "", nil
	}

	ChaosKeyValue := ones.GetToken("chaos")

	url := fmt.Sprintf("https://dns.projectdiscovery.io/dns/%s/subdomains", ones.Chaos)

	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", ChaosKeyValue)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}

	var obj map[string]interface{}
	_ = json.Unmarshal(resp.Body(), &obj)

	if obj["error"].(string) == "invalid token" {
		return TodoChaos()
	}

	v := reflect.ValueOf(obj["subdomains"])
	count := v.Len()
	for i := 0; i < count; i++ {
		resp2 = append(resp2, fmt.Sprint(v.Index(i))+"."+ones.Chaos)
	}

	//fmt.Println(string(resp.Body()))
	return string(resp.Body()), resp2

}
