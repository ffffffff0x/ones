package engine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	ones "ones/mod"
	"strconv"
	"sync"
	"time"
)

var TmpSlice2 []string
var SumSlice2 []string
var hunterRecursion = 0

func TodoHunter() []string {

	HunterKeyValue := ones.GetToken("hunter")
	//fmt.Println(HunterKeyValue[1 : len(HunterKeyValue)-1])

	num := 0
	pagesize := 0
	if ones.Num <= 100 {
		num = 100
		pagesize = ones.Num
	} else {
		num = ones.Num
		pagesize = 100
	}

	maxPage := num / 100

	if num%100 > 0 {
		maxPage++
	}

	for keyword := 1; keyword <= maxPage; keyword++ {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// API 有速率限制
		time.Sleep(2 * time.Second)
		go func(keyword int, group *sync.WaitGroup) {
			TmpSlice2 := SendReq2(HunterKeyValue, keyword, pagesize)
			SumSlice2 = append(TmpSlice2)
			group.Done()
		}(keyword, wg)
		wg.Wait()
	}

	return SumSlice2

}

// # https://hunter.qianxin.com/home/helpCenter
func SendReq2(key string, page int, pagesize int) []string {

	if hunterRecursion-ones.Recursion == 0 {
		return nil
	}

	byte1 := []byte(ones.Hunter)
	base64str := base64.StdEncoding.EncodeToString(byte1)
	//fmt.Println(base64str)

	url := fmt.Sprintf("https://hunter.qianxin.com/openApi/search?api-key=%s&search=%s&page=%d&page_size=%d", key, base64str, page, pagesize)
	//fmt.Println(url)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		//os.Exit(3)
	}
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		//os.Exit(3)
	}

	var hunter ones.HunterInfo
	err = json.Unmarshal(resp, &hunter)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hunter.Code == 401 {
		log.Println("hunter token 疑似失效", key)
		hunterRecursion += 1
		return SendReq2(ones.GetToken("hunter"), page, pagesize)
	}

	for _, v := range hunter.Data.Arr {
		//fmt.Println(v.IP + ":" + strconv.Itoa(v.Port))
		resp2 = append(resp2, v.IP+":"+strconv.Itoa(v.Port))
	}

	//fmt.Println(hunter.Data.Arr)
	//fmt.Println(string(resp))
	return resp2
}
