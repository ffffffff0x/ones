package utils

import (
	"bufio"
	"fmt"
	onese "ones/engine"
	ones "ones/mod"
	"os"
)

var AllJson = ""
var TmpSlice []string
var HostSlice []string

func OutputProcess() {

	if ones.Fofa != "" {
		AllJson, TmpSlice = onese.TodoFofa()
		HostSlice = append(TmpSlice)
	}
	if ones.Quake != "" {
		AllJson, TmpSlice = onese.TodoQuake()
		HostSlice = append(TmpSlice)
	}
	if ones.Hunter != "" {
		AllJson, TmpSlice = onese.TodoHunter()
		HostSlice = append(TmpSlice)
	}
	if ones.Shodan != "" {
		onese.TodoShodan()
	}
	if ones.Zoomeye != "" {
		onese.TodoZoomeye()
	}
	if ones.Chaos != "" {
		AllJson, TmpSlice = onese.TodoChaos()
		HostSlice = append(TmpSlice)
	}

	if ones.Json != "" {
		TodoJson()
	}
	if ones.Txt != "" {
		TodoTxt()
	}
	if ones.Json == "" && ones.Txt == "" {
		ToRaw()
	}

}
func TodoJson() {

	filePath := ones.Json
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(AllJson)
	write.Flush()

	fmt.Println("json 格式文件已输出到: ", ones.Json)
}

func TodoTxt() {

	filePath := ones.Txt
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	for _, v := range HostSlice {
		write.WriteString(v + "\n")
	}

	write.Flush()

	fmt.Println("txt 格式文件已输出到: ", ones.Txt)

}

func ToRaw() {
	for _, v := range HostSlice {
		fmt.Println(v)
	}

}
