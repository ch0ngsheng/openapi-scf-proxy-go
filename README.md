# openapi-scf-proxy-go
This repository is inspired by [Ice-Hazymoon/openai-scf-proxy](https://github.com/Ice-Hazymoon/openai-scf-proxy).

## 腾讯云/阿里云函数SCF服务的设置
添加环境变量：
* TargetHostURL=https://api.openai.com
* ProxyListenOn=0.0.0.0:9000

需要上传的压缩包，可以使用以下命令制作，或从[Releases](https://github.com/ch0ngsheng/openapi-scf-proxy-go/releases)处获取。
```shell
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main scf_bootstrap
```

函数服务部署成功后，通过APIGW开放域名（公网访问地址）访问。

云函数相关的详细操作步骤，可以参考[这里](https://github.com/Ice-Hazymoon/openai-scf-proxy)。
