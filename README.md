# chatgpt-web
[![Release](https://img.shields.io/github/v/release/869413421/chatgpt-web.svg?style=flat-square)](https://github.com/869413421/wechatbot/releases/tag/v1.1.3)
![Github stars](https://img.shields.io/github/stars/869413421/chatgpt-web.svg)
![Forks](https://img.shields.io/github/forks/869413421/chatgpt-web.svg?style=flat-square)
> 本项目可以一键部署属于自己定制化的 chatgpt web 程序(兼容gpt3.5)，
> 只需下载release中对应平台的项目文件，修改配置后执行，打开 http://127.0.0.1:8080 ，便可以获得属于自己的chatgpt网站。
>
> 参考项目：[codegen](https://github.com/git-cloner/codegen)

> 项目当前默认为示例中AI聊天机器人参数，可以根据自己需求定制化。
>
> **注意，每个参数都可能影响你得到不一样的聊天效果,改变一个参数你就可能得到另一种回答，所以请自己尝试去调试，不要上来就抱怨人工智障。文档中有二十多中参数示例，如AI聊天机器人
> ，产品名称生成，python代码修复器等等等...**
>
> 详情参考官方详细[参数示例](https://beta.openai.com/examples)

# 更新记录
- [x] fix: 支持gpt-4模型，修改前端空白BUG。 2023-03-30
- [x] fix: 增加用户模块，认证页面，接口jwt验证。 2023-03-27
- [x] fix: 修复前端富文本显示问题，优化dockerfile。 2023-03-27
- [x] fix: 优化前端显示界面。 2023-03-20
- [x] feat: 增加接口代理配置。 2023-03-20
- [x] fix: 修复前端部分BUG，优化富文本代码格式。 2023-03-13
- [x] feat: 增加socsk5代理的支持，命令行参数配置。2023-03-13
- [x] feat: 增加docker-compose.yaml。2023-03-08
- [x] fix: 修复basic auth 。 2023-03-08
- [x] feat：修改为默认不开启代理。2023-03-06
- [x] feat：增加代理配置，解决国内无法使用。2023-03-04

# 项目功能
* 请求openai增加代理（防墙）
* AI性格设定
* 兼容3.0和3.5API
* 基本问答界面
* 参数可配置
* markdown语法
* 提问上下文
# 使用前提
> 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。


# 快速开始

`第一种：直接下载二进制(适合对编程不了解的同学)`

> 非技术人员请直接下载release中的[压缩包](https://github.com/869413421/chatgpt-web/releases) ，请根据自己系统以及架构选择合适的压缩包，下载之后直接解压运行。

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

# 使用docker运行
你可以使用docker快速运行本项目。
`第一种：基于环境变量运行`

```sh
# 运行项目，环境变量参考下方配置说明
$ docker run -itd --name chatgpt-web --restart=always \
 -e APIKEY=换成你的key \
 -e APIURL= \
 -e MODEL=gpt-3.5-turbo-0301 \
 -e BOT_DESC=你是一个AI助手,我需要你模拟一名温柔贴心的女朋友来回答我的问题. \
 -e MAX_TOKENS=512 \
 -e TEMPREATURE=0.9 \
 -e TOP_P=1 \
 -e FREQ=0.0 \
 -e PRES=0.6 \
 -e PROXY=http://host.docker.internal:10809 \
 -e AUTH_USER= \
 -e AUTH_PASSWORD= \
 -p 8080:8080 \
 --add-host="host.docker.internal:host-gateway" \
 qingshui869413421/chatgpt-web:latest
```

`注意`：`host.docker.internal`会指向容器所在宿主机的IP，因此只需要更改端口为你的代理端口即可。

运行命令中映射的配置文件参考下边的配置文件说明。

`第二种：基于配置文件挂载运行`

```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
$ cp config.dev.json config.json  # 其中 config.dev.json 从项目的根目录获取

# 运行项目
$ docker run -itd --name chatgpt-web -v `pwd`/config.json:/app/config.json -p 8080:8080 qingshui869413421/chatgpt-web:latest
```

其中配置文件参考下边的配置文件说明。

# 使用docker-docompose 运行

``docker compose up -d``


# 配置文件说明

```json
{
  "api_key": "your api key",
  "api_url": "",
  "port": 8080,
  "listen": "",
  "bot_desc": "你是一个AI助手，我需要你模拟一名温柔贴心的女朋友来回答我的问题。",
  "proxy": "http://host.docker.internal:10809",
  "model": "gpt-3.5-turbo-0301",
  "max_tokens": 512,
  "temperature": 0.9,
  "top_p": 1,
  "frequency_penalty": 0.0,
  "presence_penalty": 0.6,
  "auth_user": "",
  "auth_password": ""
}
```

````
api_key：openai api_key
api_url: openai api接口地址 不填使用默认 https://api.openai.com/v1 注，该服务的提供者可以看到你的明文请求(包括你在OpenAI的key)，建议自建或使用可信来源
port: http服务端口
listen: http服务监听地址，不填默认监听0.0.0.0
proxy: openai请求代理，防墙。 例如 http://127.0.0.1:7890 socks5://127.0.0.1:7890
bot_desc：AI特征，非常重要，功能等同给与AI一个身份设定
max_tokens: GPT响应字符数，最大2048，默认值512。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认text-davinci-003，具体选项参考官网训练场
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
top_p: 使用温度采样的替代方法称为核心采样，其中模型考虑具有top_p概率质量的令牌的结果。因此，0.1 意味着只考虑包含前 10% 概率质量的代币。
frequency_penalty:
presence_penalty:
auth_user": http基本认证用户名(空表示不开启验证)
auth_password": http基本认证密码
````

# NGINX反向代理配置样例

这里提供一份使用NGINX反向代理该软件的样例配置，方便集成于现有的站点，添加用户认证，套TLS等，该文件一般对应于`/etc/nginx/sites-available/default`文件，需要自行修改。

```nginx
# 监听80端口，跳转https
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    location / {
        return 301 https://$host$request_uri;
    }
}
# 监听443端口，使用https提供服务
server {
    # SSL相关配置来自 https://ssl-config.mozilla.org/
    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    # 证书路径，建议Fullchain
    ssl_certificate /path/to/your/cert.pem;
    # 私钥路径
    ssl_certificate_key /path/to/your/key.pem;
    ssl_session_timeout 1d;
    ssl_session_cache shared:MozSSL:10m;
    ssl_session_tickets off;
    # 执行下面的命令下载dhparam
    # curl https://ssl-config.mozilla.org/ffdhe2048.txt > /path/to/dhparam
    ssl_dhparam /path/to/dhparam;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384;
    ssl_prefer_server_ciphers off;
    # HSTS (ngx_http_headers_module is required) (63072000 seconds)
    add_header Strict-Transport-Security "max-age=63072000" always;
    # SSL配置结束

    server_name _;
    charset utf-8;
    client_max_body_size 5m;

    # 如果需要将chatgpt-web置于某一路径下，使用这个location配置
    location /your/path/ {
        # 基本身份认证 设定
        # 提示语
        auth_basic "Auth Require";
        # 认证配置文件 格式请参考 https://nginx.org/en/docs/http/ngx_http_auth_basic_module.html
        auth_basic_user_file /path/to/passwd;

        # 反向代理 假设chatgpt-web监听端口为8080
        proxy_pass http://127.0.0.1:8080/;
        proxy_http_version 1.1;
        # 反向代理超时时间设定(OpenAI的反应比较慢，设定为120秒后才超时)
        proxy_read_timeout 120s;
    }

    # 如果chatgpt-web放置于根路径，使用这个location配置
    location / {
        auth_basic "Auth Require";
        auth_basic_user_file /etc/nginx/passwd;

        proxy_pass http://127.0.0.1:8080/;
        proxy_http_version 1.1;
        proxy_read_timeout 120s;

        # 位于根路径时不需要修改index.html
    }

}
```

# Linux系统systemd服务配置

可以使用`systemd`配置`chatgpt-web`开机自启，假设可执行文件和相关资源文件放置在`/var/www/chatgpt-web/`目录下，`chatgpt-web`二进制文件需要其他用户可读可执行权限，其余资源文件需要其他用户可读权限，并且已经配置好`config.json`。

在目录`/etc/systemd/system/`下新建文件`chatgpt-web.service`，以下是文件样例。

```ini
[Unit]
Description=chatgpt-web
Documentation=https://github.com/869413421/chatgpt-web
# 在网络启动完成后运行
After=network.target nss-lookup.target

[Service]
# 使用随机用户执行该服务
DynamicUser=yes
# 指定工作目录
WorkingDirectory=/var/www/chatgpt-web/
# 执行程序
ExecStart=/var/www/chatgpt-web/chatgpt-web

[Install]
WantedBy=multi-user.target
```
保存后使用`systemctl daemon-reload`更新systemd配置文件，使用`systemctl start/stop chatgpt-web`启动/停止服务，使用`systemctl enable/disable chatgpt-web`启用/禁用服务开机自启。

可以使用`journalctl --unit chatgpt-web.service`查看程序日志。

# 免责声明 Disclaimers
The code is for demo and testing only. 代码仅用于演示和测试。

⚠⚠⚠请勿将本系统代码用于商业用途！

仿冒或冒用ChatGPT、OpenAI名义开展经营活动，可能构成《商标法》、《反不正当竞争法》下的一系列侵权行为； 以之牟利造成消费者损失的，可能产生《商标法》、《反不正当竞争法》、《消费者权益保护法》下的民事或行政责任，情节严重并造成重大损失的，还有可能构成刑事犯罪； 如果提供这种跨境经营服务存在私自搭建国际信道的情形，还有可能违反《网络安全法》、《刑法》的相关规定，承担行政责任或构成刑事犯罪。
