package mod

import "flag"

var (
	V       bool
	Fofa    string
	Quake   string
	Shodan  string
	Chaos   string
	Zoomeye string
	Hunter  string
	Num     int
	Version = "v1.0.0"
	Json    string
	Txt     string
)

func init() {

	flag.BoolVar(&V, "version", false, "显示版本号")
	flag.StringVar(&Fofa, "fofa", "", "fofa 查询")
	flag.StringVar(&Quake, "quake", "", "quake 查询")
	//flag.StringVar(&Shodan, "shodan", "", "shodan 查询")
	flag.StringVar(&Chaos, "chaos", "", "chaos 查询")
	//flag.StringVar(&Zoomeye, "zoomeye", "", "zoomeye 查询")
	flag.StringVar(&Hunter, "hunter", "", "hunter 查询")
	flag.IntVar(&Num, "num", 100, "数量")
	flag.StringVar(&Json, "json", "", "导出json格式")
	flag.StringVar(&Txt, "txt", "", "导出txt格式")

}
