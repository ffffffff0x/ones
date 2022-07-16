package engine

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	ones "ones/mod"
	"os"
	"strconv"
	"sync"
)

var TmpSlice1 []string
var SumSlice1 []string
var shodanRecursion = 0

func TodoShodan() []string {

	ShodanKeyValue := ones.GetToken("shodan")
	//fmt.Println(ShodanKeyValue[1 : len(ShodanKeyValue)-1])

	num := 0
	if ones.Num <= 0 {
		num = 100
	} else {
		num = ones.Num
	}

	maxPage := num / 100

	if num%100 > 0 {
		maxPage++
	}

	wg := &sync.WaitGroup{}
	for keyword := 1; keyword <= maxPage; keyword++ {
		wg.Add(1)
		go func(keyword int, group *sync.WaitGroup) {
			TmpSlice := SendReq1(ShodanKeyValue, keyword)
			SumSlice1 = append(TmpSlice)
			group.Done()
		}(keyword, wg)
	}
	wg.Wait()

	return SumSlice1

}

func SendReq1(key string, num int) []string {
	shodanRecursion += 1
	if shodanRecursion % ones.Recursion == 0 {
		return nil
	}

	url := fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&page=%d&query=%s", key, num, ones.Shodan)
	//fmt.Println(url)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		log.Println("shodan token 疑似失效", key)
		return SendReq1(ones.GetToken("shodan"), num)
	}

	//fmt.Println(string(resp))

	var shodan ones.ShodanInfo
	_ = json.Unmarshal(resp, &shodan)

	for _, v := range shodan.Matches {
		resp2 = append(resp2, v.IPStr+":"+strconv.Itoa(v.Port))
	}
	return resp2

}
