package mod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type Configs map[string]json.RawMessage

var configPath string = "ones.conf"

type MainConfig struct {
	FofaEmail string `json:"fofa_email"`
	FofaKey   string `json:"fofa_key"`
	ZoomKey   string `json:"zoom_key"`
	ShodanKey string `json:"shodan_key"`
	QuakeKey  string `json:"quake_key"`
	HunterKey string `json:"hunter_key"`
	ChaosKey  string `json:"chaos_key"`
}

var conf *MainConfig
var Confs Configs

var instanceOnce sync.Once

// LoadConfig 从配置文件中载入json字符串
func LoadConfig(path string) (Configs, *MainConfig) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs, mainConfig
}

// Init 初始化
func Init(path string) *MainConfig {
	if conf != nil && path != configPath {
		log.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
	}
	instanceOnce.Do(func() {
		allConfigs, mainConfig := LoadConfig(path)
		configPath = path
		conf = mainConfig
		Confs = allConfigs
	})

	return conf
}

// ConfigPath 获取配置文件路径
func ConfigPath() string {
	return configPath
}

// GetConfig 根据key获取对应的值，如果值为struct，则继续反序列化
func (cfg Configs) GetConfig(key string, config interface{}) error {
	c, ok := cfg[key]
	if ok {
		return json.Unmarshal(c, config)
	} else {
		return fmt.Errorf("fail to get cfg with key: %s", key)
	}
}
