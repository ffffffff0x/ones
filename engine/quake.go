package engine

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	ones "ones/mod"
	"os"
	"strconv"
)

var quakeRecursion = 0

func TodoQuake() (string, []string) {
	quakeRecursion += 1
	if quakeRecursion%ones.Recursion == 0 {
		return "", nil
	}

	QuakeKeyValue := ones.GetToken("quake")
	//fmt.Println(QuakeKeyValue)

	url := `https://quake.360.cn/api/v3/search/quake_service`

	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	requestBody := fmt.Sprintf(`{"query": %q,
"start": 0, 
"size": %d,
"include": ["ip","port"]}`, ones.Quake, ones.Num)

	//fmt.Println(requestBody)
	req.SetBody([]byte(requestBody))

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.Header.Add("X-QuakeToken", QuakeKeyValue)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}
	//fmt.Println(string(resp.Body()))
	var serviceInfo ones.ServiceInfo
	err := json.Unmarshal(resp.Body(), &serviceInfo)
	if err != nil {
		log.Println("quake token 疑似失效", QuakeKeyValue)
		return TodoQuake()
	}
	if serviceInfo.Code != 0 {
		os.Exit(3)
	}

	for _, value := range serviceInfo.Data {
		output := value.IP + ":" + strconv.Itoa(value.Port)
		resp2 = append(resp2, output)
	}

	//fmt.Println(string(resp.Body()))
	return string(resp.Body()), resp2

}
