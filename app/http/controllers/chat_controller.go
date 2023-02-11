package controllers

import (
	"github.com/869413421/chatgpt-web/config"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
)

// ChatController 首页控制器
type ChatController struct {
	BaseController
}

// NewChatController 创建控制器
func NewChatController() *ChatController {
	return &ChatController{}
}

//Index 首页
func (c *ChatController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

type Question struct {
	Text string `json:"text"`
}

//Completion 回复
func (c *ChatController) Completion(ctx *gin.Context) {
	question := &Question{}
	err := ctx.BindJSON(question)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if question.Text == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request text is empty", nil)
		return
	}

	cnf := config.LoadConfig()
	client := gogpt.NewClient(cnf.ApiKey)
	prompt := cnf.BotDesc + "\n" + question.Text
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

	c.ResponseJson(ctx, http.StatusOK, "", resp.Choices[0].Text)
}
