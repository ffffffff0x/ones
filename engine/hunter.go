package engine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	ones "ones/mod"
	"os"
	"strconv"
)

// # https://hunter.qianxin.com/home/helpCenter

func TodoHunter() (string, []string) {

	byte1 := []byte(ones.Hunter)
	base64str := base64.StdEncoding.EncodeToString(byte1)
	//fmt.Println(base64str)

	HunterKeyValue := string(ones.Confs["hunter_key"])
	HunterKeyValue = HunterKeyValue[1 : len(HunterKeyValue)-1]
	//fmt.Println(HunterKeyValue[1 : len(HunterKeyValue)-1])

	url := fmt.Sprintf("https://hunter.qianxin.com/openApi/search?api-key=%s&search=%s&page=1&page_size=%d", HunterKeyValue, base64str, ones.Num)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		os.Exit(3)
	}

	var hunter ones.HunterInfo
	err = json.Unmarshal(resp, &hunter)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range hunter.Data.Arr {
		resp2 = append(resp2, v.IP+":"+strconv.Itoa(v.Port))
	}

	//fmt.Println(hunter.Data.Arr)
	//fmt.Println(string(resp))
	return string(resp), resp2

}
