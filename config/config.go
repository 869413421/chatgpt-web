package config

import (
	"encoding/json"
	"fmt"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"log"
	"os"
	"strconv"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// gpt apikey
	ApiKey string `json:"api_key"`
	Port   int    `json:"port"`
	//
	BotDesc string `json:"bot_desc"`
	// GPT请求最大字符数
	MaxTokens int `json:"max_tokens"`
	// GPT模型
	Model string `json:"model"`
	// 热度
	Temperature      float64 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	PresencePenalty  float32 `json:"presence_penalty"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			MaxTokens:        60,
			Port:             8080,
			Model:            "text-davinci-003",
			Temperature:      0.9,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.6,
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}
		// 有环境变量使用环境变量
		ApiKey := os.Getenv("APIKEY")
		Model := os.Getenv("MODEL")
		MaxTokens := os.Getenv("MAX_TOKENS")
		Temperature := os.Getenv("TEMPREATURE")
		TopP := os.Getenv("TOP_P")
		FrequencyPenalty := os.Getenv("FREQ")
		PresencePenalty := os.Getenv("PRES")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}

		if Model != "" {
			config.Model = Model
		}

		if MaxTokens != "" {
			max, err := strconv.Atoi(MaxTokens)
			if err != nil {
				logger.Danger(fmt.Sprintf("config MaxTokens err: %v ,get is %v", err, MaxTokens))
				return
			}
			config.MaxTokens = max
		}
		if Temperature != "" {
			temp, err := strconv.ParseFloat(Temperature, 64)
			if err != nil {
				logger.Danger(fmt.Sprintf("config Temperature err: %v ,get is %v", err, Temperature))
				return
			}
			config.Temperature = temp
		}
		if TopP != "" {
			temp, err := strconv.ParseFloat(TopP, 32)
			if err != nil {
				logger.Danger(fmt.Sprintf("config Temperature err: %v ,get is %v", err, TopP))
				return
			}
			config.TopP = float32(temp)
		}
		if FrequencyPenalty != "" {
			temp, err := strconv.ParseFloat(FrequencyPenalty, 32)
			if err != nil {
				logger.Danger(fmt.Sprintf("config Temperature err: %v ,get is %v", err, FrequencyPenalty))
				return
			}
			config.FrequencyPenalty = float32(temp)
		}
		if PresencePenalty != "" {
			temp, err := strconv.ParseFloat(PresencePenalty, 32)
			if err != nil {
				logger.Danger(fmt.Sprintf("config Temperature err: %v ,get is %v", err, PresencePenalty))
				return
			}
			config.PresencePenalty = float32(temp)
		}

	})
	if config.ApiKey == "" {
		logger.Danger("config err: api key required")
	}

	return config
}
