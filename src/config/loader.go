package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/**
* config 初始化
 */
func init() {
	env := make(map[string]string)
	config, err := ioutil.ReadFile("config/index.json")
	if err != nil {
		log.Fatal("找不到index.json")
	}

	err = json.Unmarshal(config, &env)
	if err != nil {
		log.Println(err.Error())
		return
	}

	config, err = ioutil.ReadFile("config/" + env["app_env"] + ".json")
	if err != nil {
		log.Fatal("找不到" + env["app_env"] + ".json")
	}

	results := make(map[string]string)
	err = json.Unmarshal(config, &results)
	if err != nil {
		log.Fatalln("載入json, 發生未知錯誤")
	}

	_ = os.Setenv("app_env", env["app_env"])
	_ = os.Setenv("http_port", env["http_port"])
	_ = os.Setenv("maintain", env["maintain"])
	for k, v := range results {
		_ = os.Setenv(k, v)
	}
}
