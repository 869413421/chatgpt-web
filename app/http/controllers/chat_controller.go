package controllers

import (
	"context"
	"github.com/869413421/chatgpt-web/pkg/types"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	gogpt "github.com/sashabaranov/go-openai"
	"golang.org/x/net/proxy"

	"github.com/869413421/chatgpt-web/config"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"github.com/gin-gonic/gin"
)

var chatModels = []string{gogpt.GPT432K0314, gogpt.GPT4, gogpt.GPT40314, gogpt.GPT432K, gogpt.GPT3Dot5Turbo, gogpt.GPT3Dot5Turbo0301}

// ChatController 首页控制器
type ChatController struct {
	BaseController
}

// NewChatController 创建控制器
func NewChatController() *ChatController {
	return &ChatController{}
}

// Index 首页
func (c *ChatController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

// Completion 回复
func (c *ChatController) Completion(ctx *gin.Context) {
	var request gogpt.ChatCompletionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	logger.Info(request)
	if len(request.Messages) == 0 {
		c.ResponseJson(ctx, http.StatusBadRequest, "request messages required", nil)
		return
	}

	cnf := config.LoadConfig()
	gptConfig := gogpt.DefaultConfig(cnf.ApiKey)

	if cnf.Proxy != "" {
		transport := &http.Transport{}

		if strings.HasPrefix(cnf.Proxy, "socks5h://") {
			// 创建一个 DialContext 对象，并设置代理服务器
			dialContext, err := newDialContext(cnf.Proxy[10:])
			if err != nil {
				panic(err)
			}
			transport.DialContext = dialContext
		} else {
			// 创建一个 HTTP Transport 对象，并设置代理服务器
			proxyUrl, err := url.Parse(cnf.Proxy)
			if err != nil {
				panic(err)
			}
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
		// 创建一个 HTTP 客户端，并将 Transport 对象设置为其 Transport 字段
		gptConfig.HTTPClient = &http.Client{
			Transport: transport,
		}

	}

	// 自定义gptConfig.BaseURL
	if cnf.ApiURL != "" {
		gptConfig.BaseURL = cnf.ApiURL
	}

	client := gogpt.NewClientWithConfig(gptConfig)
	if request.Messages[0].Role != "system" {
		newMessage := append([]gogpt.ChatCompletionMessage{
			{Role: "system", Content: cnf.BotDesc},
		}, request.Messages...)
		request.Messages = newMessage
		logger.Info(request.Messages)
	}

	// cnf.Model 是否在 chatModels 中
	if types.Contains(chatModels, cnf.Model) {
		request.Model = cnf.Model
		resp, err := client.CreateChatCompletion(ctx, request)
		if err != nil {
			c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.ResponseJson(ctx, http.StatusOK, "", gin.H{
			"reply":    resp.Choices[0].Message.Content,
			"messages": append(request.Messages, resp.Choices[0].Message),
		})
	} else {
		prompt := ""
		for _, item := range request.Messages {
			prompt += item.Content + "/n"
		}
		prompt = strings.Trim(prompt, "/n")

		logger.Info("request prompt is %s", prompt)
		req := gogpt.CompletionRequest{
			Model:            cnf.Model,
			MaxTokens:        cnf.MaxTokens,
			TopP:             cnf.TopP,
			FrequencyPenalty: cnf.FrequencyPenalty,
			PresencePenalty:  cnf.PresencePenalty,
			Prompt:           prompt,
		}

		resp, err := client.CreateCompletion(ctx, req)
		if err != nil {
			c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.ResponseJson(ctx, http.StatusOK, "", gin.H{
			"reply": resp.Choices[0].Text,
			"messages": append(request.Messages, gogpt.ChatCompletionMessage{
				Role:    "assistant",
				Content: resp.Choices[0].Text,
			}),
		})
	}
}

type dialContextFunc func(ctx context.Context, network, address string) (net.Conn, error)

func newDialContext(socks5 string) (dialContextFunc, error) {
	baseDialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}

	if socks5 != "" {
		// split socks5 proxy string [username:password@]host:port
		var auth *proxy.Auth = nil

		if strings.Contains(socks5, "@") {
			proxyInfo := strings.SplitN(socks5, "@", 2)
			proxyUser := strings.Split(proxyInfo[0], ":")
			if len(proxyUser) == 2 {
				auth = &proxy.Auth{
					User:     proxyUser[0],
					Password: proxyUser[1],
				}
			}
			socks5 = proxyInfo[1]
		}

		dialSocksProxy, err := proxy.SOCKS5("tcp", socks5, auth, baseDialer)
		if err != nil {
			return nil, err
		}

		contextDialer, ok := dialSocksProxy.(proxy.ContextDialer)
		if !ok {
			return nil, err
		}

		return contextDialer.DialContext, nil
	} else {
		return baseDialer.DialContext, nil
	}
}
