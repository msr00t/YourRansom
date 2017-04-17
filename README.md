YourRansom
---

YourRansom 是使用 Golang 编写的一个加密工具，具有加密参数内置、密钥随机生成、在线下载readme文件等功能，设计上参考了【勒索软件】的设计思想。或者换句话说这就是一个自制的学习用的【勒索软件】仿制品。

在下只是个萌新，程序写的也很 naive，请各位大佬轻些疼♂爱。

## 自行编译

因为 YourRansom 将加密的配置内置在编译后的二进制文件中，所以如果您想要使用自己的 YourRansom ，就需要调整参数并编译一个自己的 YourRansom 。

### 准备环境

YourRansom 使用 Golang 编写，编译前也需要准备对应的 [Golang 环境](https://golang.org/doc/install)，或使用 cloud9 等在线环境编译（通常也需要自己装 Golang 环境就是了）。

之后即可通过 go get 从 GitHub 获取 YourRansom 的源文件：`go get github.com/YourRansom/YourRansom`

### 配置参数

YourRansom 将配置存储在二进制文件中，但并不是直接存储原数据（曾经是，后来我感觉那样太年轻太简单有时天真，于是改成了如今的模式），而是存储 JSON 格式的配置文件使用 DES 加密后又 base64 得到的字符串。我专门为此提供了一个配置生成器与样板文件：[YourRansom/confGen](https://github.com/YourRansom/confGen)，只需要对着表格将数据填写完成，再执行 confGen 即可获得配置信息。

配置项表：

| 配置项名 | 配置说明 | 备注 |
| --- | --- | --- |
| PubKey | RSA公钥 | 可使用[生成工具]([https://github.com/YourRansom/genKeypair](https://github.com/YourRansom/genKeypair))生成一个RSA密钥对，**请注意自行将所有换行替换为\n**。 |
| Filesuffix | 加在被加密文件后的后缀名 |
| KeyFilename | 加密后存储Key的文件名 |
| DkeyFilename | 解密时用于读取解密Key的文件名 |
| Alert | 程序启动时显示的提示信息 |
| Readme | 离线readme的内容 | 仅当在线readme下载失败时生效 |
| ReadmeFilename | 离线readme的存储文件名 | 同上 |
| ReadmeUrl | 在线readme的下载地址 | 留空表示不启用在线readme下载 |
| ReadmeNetFilename | 在线readme的存储文件名 |
| EncSuffix | 指定要被加密的文件后缀 | 格式为`后缀1\|后缀2\|后缀3` |

执行：

```
$ ./confGen MyRansom example.json

```

如不报错则说明配置文件已经加密完成，打开`原文件名.enc`即可得到加密后的配置文件。

使用您的文本编辑器打开`$GOPATH/src/github.com/YourRansom/YourRansom/config.go`，如果您未设置`GOPATH`环境变量，那么打开`$HOME/go/src/github.com/YourRansom/YourRansom/config.go`(Unix-like)或`%USERPROFILE%\go\src\github.com\YourRansom\YourRansom\config.go`(Windows)，将配置文件和加密密钥填入，配置即完成。

### 编译生成

直接执行命令，生成的二进制文件会直接输出至当前目录：

```
$ GOOS=[windows|linux|darwin] GOARCH=[386|amd64] go build github.com/YourRansom/YourRansom

```

考虑到兼容性问题，如果您要为使用 Windows 系统的用户提供服务，建议编译为 win32 程序，如果面向 Linux ，您可能需要编译32位和64位两个版本，而 macOS（darwin）只需要64位版本就够了。

### 加密解密

直接执行生成的二进制文件即可加密，而解密要复杂得多。

解密需要获得加密是生成的密钥文件，具体文件名取决于配置项`KeyFilename`，然后使用 PubKey 所对应的私钥和 [AES 密钥解密工具](https://github.com/YourRansom/YourRansom-keyDecryptor) 解密该文件，得到一个`YourRansom.dkey`文件，将其更名为你设置的`DkeyFilename`配置项。

将解密后的 Dkey 文件至于 YourRansom 同目录下，再次执行 YourRansom 即可解密。

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

### 使用在线IDE

如果您不喜欢在本地安装 Golang 环境，只是想简单地将它编译出来尝尝鲜的话，使用 Cloud9 、 Wide 之类的在线环境大概会是一个好选择。

首先在 [https://c9.io](https://c9.io) 注册一个账户并登陆，然后新建一个 Workspace 后即得到了一个在线的编译环境，剩下的操作参考前面的说明即可。

## LICENSE

基于GPLv3协议和[不作恶]附加协议开源。

[不作恶]附加协议内容在`LICENSE.additional`文件中。
