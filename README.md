YourRansom
---

简单地用Golang写的“勒索软件”仿制品，仅用于学习和研究，实现了基本的加密、解密、改扩展名操作。支持自行设定readme文件、在线下载readme文件、自定义文件后缀名和提示信息。

加密后的key将存于同目录下`YourRansom.key`文件，该文件解密后得到`YourRansom.dkey`置于同目录，再次执行软件本体即可解密。

名字向前辈 [Tlaster/YourAV](https://github.com/Tlaster/YourAV) 致敬。

## TOC

- [使用指南](#使用指南)
- [注意事项](#注意事项)
  - [解密](#解密)
  - [RSA密钥对](#RSA密钥对)
- [LICENSE](#LICENSE)


## 使用指南

1. 下载并安装Golang环境：<https://golang.org/doc/install>

2. 使用go get下载该项目到本地：`go get github.com/YourRansom/YourRansom`

3. 修改YourRansom的参数，使用您的文本编辑器打开`$GOPATH/src/github.com/YourRansom/YourRansom/config.go`，如果您未设置`GOPATH`环境变量，
那么打开`$HOME/go/src/github.com/YourRansom/YourRansom/config.go`(Unix-like)或`%USERPROFILE%\go\src\github.com\YourRansom\YourRansom\config.go`(Windows)，将配置文件和加密密钥填入，配置即完成。

4. 编译YourRansom，执行：`go build -x github.com/YourRansom/YourRansom`，编译完成的可执行文件将存于当前工作目录。


## 注意事项

### 解密

解密需要将加密时生成的key文件使用[解密工具](https://goo.gl/J2HSk0)和配置中的公钥所对应的RSA私钥解密为dkey文件后与YourRansom程序置于同目录，并再次运行YourRansom。

解密key的步骤可以参考：<https://goo.gl/Z8uc5l>

### RSA密钥对

一个RSA密钥对由一个公钥和一个私钥组成，YourRansom参数中需要填入的是公钥，但请务必妥善保管私钥。

关于如何生成，可以使用我们提供的[生成工具](https://github.com/YourRansom/genKeypair)。

默认设置中的公钥对应的私钥为：
```
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDbdg+Rcnrsa/7uT9zoMNN+nyo8ajk9eqL6S12Q1gI21gkr9yDr
LzsDfBKssgf50iMbgE7UShEAm61qIFjMWG9XpVzxAgGz8el52j1yYTy/6XfmZVXe
Ycw6wosP/xD+peafeJGWenC5Qkw9ucFlakQFQ2QHr9tIxc/AifSAlGabHQIDAQAB
AoGBAM9RpX1ab4NetlK9AUwbrAAnLkgqdO5+Ju5aOga0FR1mbv2olOF4GcC9+gpI
mL5I5D97o3xqh8tSRa1G53wLYwniP+mZ7YodKX9w5SuuBFRZDmBcPQBAqYu6EAec
nN/wPwBZlbRqqXYDdHebCZmSuDCpjmrthpv1nxPJjpRiqJfFAkEA83cS/HxH1+44
KwsUxpfYko8yogU+w7PqlRtWauLYAUR5on0g9eD5Bp/T936ERfcCBFbUTZ9+yWYa
1lIOLJfhJwJBAObCm+iBalIPvzqHvOoQf4eO2xvX1b3V+FxT9b+LuiUqPEMLW6TQ
lZw16bertfWBofjjm/URdATgdsE5hKgAxBsCQFeiltz3R0z8XI9xz6qkYbpvfQRA
6xS6oEfHrVWQDbx3D2ljrQeUUU8HHN9LVQVyIfG553WBYbvQ2vwmUR/QE6UCQQC+
xXxnBzaCiQoqtTT0vJbx1qRFrHXD7zTX/4FWzYkiWHxhYO5unxJQhjGl6osPYBAr
1t+EBt3Heloy+/4zdg6pAkB7POl1VNVeWvhDP8Oh5EJPVRQXj6owM/qpE4PIkYjH
I9UW6c8ZiByQek6xXA419ML0ljeTIau3xwWnYpdHKag0
-----END RSA PRIVATE KEY-----
```


## LICENSE

基于GPLv3协议和[不作恶]附加协议开源。

[不作恶]附加协议内容在`LICENSE.additional`文件中。
