package mod

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var configPath string = "ones-config.json"

type MainConfig struct {
	FofaKey   []string `json:"fofa_key"`	// fofa_email:fofa_key
	ZoomKey   []string `json:"zoom_key"`
	ShodanKey []string `json:"shodan_key"`
	QuakeKey  []string `json:"quake_key"`
	HunterKey []string `json:"hunter_key"`
	ChaosKey  []string `json:"chaos_key"`
}

var Confs MainConfig
// 统计当前使用的第x个api token
var fofaNumber = 0
var zoomNumber = 0
var shodanNumber = 0
var quakeNumber = 0
var hunterNumber = 0
var chaosNumber = 0

var instanceOnce sync.Once

// LoadConfig 从配置文件中载入json字符串
func LoadConfig(path string) MainConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return *mainConfig
}

// Init 初始化
func Init(path string) {
	_, err := os.Stat(path)
	if err != nil {
		// 没有配置文件，生成配置文件
		configBytes, _ := json.MarshalIndent(&Confs, "", "  ")
		ioutil.WriteFile(path, configBytes, 0666)
		log.Printf("generate config file at %s", path)
	} else {
		instanceOnce.Do(func() {
			Confs = LoadConfig(path)
			configPath = path
		})
	}
}

// ConfigPath 获取配置文件路径
func ConfigPath() string {
	return configPath
}

// GetToken 获取当前应该使用的token，循环遍历 fofa zoom shodan quake hunter chaos
func GetToken(types string) string {
	var token string
	switch types {
	case "fofa":
		{
			if Confs.FofaKey == nil {
				log.Panicln("fofaKey不能为空，格式 fofa_email:fofa_key")
			}
			number := fofaNumber % len(Confs.FofaKey)
			token = Confs.FofaKey[number]
			fofaNumber += 1
		}
	case "zoom":
		{
			if Confs.ZoomKey == nil {
				log.Panicln("ZoomKey不能为空，格式为 [1, 2, 3]")
			}
			number := zoomNumber % len(Confs.ZoomKey)
			token = Confs.ZoomKey[number]
			zoomNumber += 1
		}
	case "shodan":
		{
			if Confs.ShodanKey == nil {
				log.Panicln("ShodanKey不能为空，格式为 [1, 2, 3]")
			}
			number := shodanNumber % len(Confs.ShodanKey)
			token = Confs.ShodanKey[number]
			shodanNumber += 1
		}
	case "quake":
		{
			if Confs.QuakeKey == nil {
				log.Panicln("QuakeKey不能为空，格式为 [1, 2, 3]")
			}
			number := quakeNumber % len(Confs.QuakeKey)
			token = Confs.QuakeKey[number]
			quakeNumber += 1
		}
	case "hunter":
		{
			if Confs.HunterKey == nil {
				log.Panicln("QuakeKey不能为空，格式为 [1, 2, 3]")
			}
			number := hunterNumber % len(Confs.HunterKey)
			token = Confs.HunterKey[number]
			hunterNumber += 1
		}
	case "chaos":
		{
			if Confs.ChaosKey == nil {
				log.Panicln("ChaosKey不能为空，格式为 [1, 2, 3]")
			}
			number := chaosNumber % len(Confs.ChaosKey)
			token = Confs.ChaosKey[number]
			chaosNumber += 1
		}
	default:
		{
		}
	}
	return token
}