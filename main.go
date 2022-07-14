package main

import (
	"flag"
	"fmt"
	ones "ones/mod"
	"ones/utils"
	"os"
)

var InNum = 0

func main() {

	flag.Parse()

	if ones.V {
		fmt.Println(ones.Version)
		os.Exit(0) // 输出版本后停止
	}

	// 解析配置
	path := ones.ConfigPath()
	ones.Init(path)

	// 处理输入
	InputNum()

}

func InputNum() {

	if ones.Fofa != "" {
		InNum += 1
	}
	if ones.Quake != "" {
		InNum += 1
	}
	if ones.Hunter != "" {
		InNum += 1
	}
	if ones.Shodan != "" {
		InNum += 1
	}
	if ones.Zoomeye != "" {
		InNum += 1
	}
	if ones.Chaos != "" {
		InNum += 1
	}

	if InNum > 2 {
		if ones.Json != "" {
			fmt.Println("查询多个引擎时,不可导出为 json")
			os.Exit(3)
		} else {
			utils.OutputProcess()
		}
	} else if InNum == 0 { // 无参数输出Help
		flag.Usage()
	} else {
		utils.OutputProcess()
	}
}
