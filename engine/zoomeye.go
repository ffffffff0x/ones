package engine

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	ones "ones/mod"
	"os"
	"strconv"
	"sync"
)

var TmpSlice []string
var SumSlice []string
var zoomeyeRecursion = 0

func TodoZoomeye() []string {

	ZoomKeyValue := ones.GetToken("zoom")
	ZoomKeyValue = ZoomKeyValue[1 : len(ZoomKeyValue)-1]
	//fmt.Println(ZoomKeyValue)

	num := 0
	if ones.Num <= 0 {
		num = 20
	} else {
		num = ones.Num
	}

	maxPage := num / 20

	if num%20 > 0 {
		maxPage++
	}

	wg := &sync.WaitGroup{}
	for keyword := 1; keyword <= maxPage; keyword++ {
		wg.Add(1)
		go func(keyword int, group *sync.WaitGroup) {
			TmpSlice := SendReq(ZoomKeyValue, keyword)
			SumSlice = append(TmpSlice)
			group.Done()
		}(keyword, wg)
	}
	wg.Wait()

	//fmt.Println(string(resp.Body()))
	return SumSlice

}

func SendReq(key string, num int) []string {
	zoomeyeRecursion += 1
	if zoomeyeRecursion % ones.Recursion == 0 {
		return nil
	}

	url := fmt.Sprintf("https://api.zoomeye.org/host/search?query=%s&page=%d", ones.Zoomeye, num)
	//fmt.Println(url)

	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Add("API-KEY", key)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		os.Exit(3)
	}

	var zoomeye ones.ZoomInfo
	_ = json.Unmarshal(resp.Body(), &zoomeye)

	for _, v := range zoomeye.Matches {
		resp2 = append(resp2, v.IP+":"+strconv.Itoa(v.Portinfo.Port))
	}

	if zoomeye.Matches == nil {
		return SendReq(ones.GetToken("zoom"), num)
	} else {
		return resp2
	}


}
