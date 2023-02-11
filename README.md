# wechatbot

> 最近ChatGPT异常火爆，本项目可以将个人微信化身GPT机器人，
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。

[![Release](https://img.shields.io/github/v/release/869413421/chatgpt.svg?style=flat-square)](https://github.com/869413421/wechatbot/releases/tag/v1.1.3)
![Github stars](https://img.shields.io/github/stars/869413421/chatgpt.svg)
![Forks](https://img.shields.io/github/forks/869413421/chatgpt.svg?style=flat-square)


> 本项目基于第二种方式实现，模型之间具体差异可以参考[官方文档](https://beta.openai.com/docs/models/overview), 详细[参数示例](https://beta.openai.com/examples) 。

# 使用前提
> * 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。

# 使用docker运行
你可以使用docker快速运行本项目。

`第一种：基于环境变量运行`

```sh
# 运行项目，环境变量参考下方配置说明
$ docker run -itd --name chatgpt-web --restart=always \
 -e APIKEY=换成你的key \
 -e MODEL=text-davinci-003 \
 -e MAX_TOKENS=512 \
 -e TEMPREATURE=0.9 \
 -e TOP_P=1 \
 -e FREQ=0.0 \
 -e PRES=0.6 \
 -p 8080:8080 \
 docker.mirrors.sjtug.sjtu.edu.cn/qingshui869413421/chatgpt-web:latest
```

运行命令中映射的配置文件参考下边的配置文件说明。

`第二种：基于配置文件挂载运行`

```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
$ cp config.dev.json config.json  # 其中 config.dev.json 从项目的根目录获取

# 运行项目
$ docker run -itd --name chatgpt-web -v `pwd`/config.json:/app/config.json -p 8080:8080 docker.mirrors.sjtug.sjtu.edu.cn/qingshui869413421/chatgpt-web:latest
```

其中配置文件参考下边的配置文件说明。

# 快速开始

`第一种：直接下载二进制(适合对编程不了解的同学)`

> 非技术人员请直接下载release中的[压缩包](https://github.com/869413421/wechatbot/releases) ，请根据自己系统以及架构选择合适的压缩包，下载之后直接解压运行。

下载之后，在本地解压，即可看到可执行程序，与配置文件：

```
# windows
1.下载压缩包解压
2.复制文件中config.dev.json更改为config.json
3.将config.json中的api_key替换为自己的
4.双击exe运行，启动服务

# linux
$ tar xf chatgpt-web-v0.0.2-darwin-arm64.tar.gz # 解压
$ cd chatgpt-web-v0.0.2-darwin-arm64
$ cp config.dev.json # 根据情况调整配置文件内容
$ ./chatgpt-web  # 直接运行

# 如果要守护在后台运行
$ nohup ./chatgpt-web &> run.log &
$ tail -f run.log
```

`第二种：基于源码运行(适合了解go语言编程的同学)`

````
# 获取项目
$ git clone https://github.com/869413421/chatgpt-web.git

# 进入项目目录
$ cd chatgpt-web

# 复制配置文件
$ copy config.dev.json config.json

# 启动项目
$ go run main.go
````

# 配置文件说明

````
{
  "api_key": "your api key",
  "auto_pass": true,
  "session_timeout": 60,
  "max_tokens": 1024,
  "model": "text-davinci-003",
  "temperature": 1,
  "reply_prefix": "来自机器人回复：",
  "session_clear_token": "清空会话"
}

api_key：openai api_key
auto_pass:是否自动通过好友添加
session_timeout：会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文。
max_tokens: GPT响应字符数，最大2048，默认值512。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认text-davinci-003，具体选项参考官网训练场
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
reply_prefix: 私聊回复前缀
session_clear_token: 会话清空口令，默认`下一个问题`
````
