# openapi-scf-proxy-go
This repository is inspired by [Ice-Hazymoon/openai-scf-proxy](https://github.com/Ice-Hazymoon/openai-scf-proxy).

## 腾讯云函数SCF服务的设置
添加环境变量：
* TargetHostURL=https://api.openai.com
* ProxyListenOn=0.0.0.0:9000

编译压缩命令：
```shell
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main scf_bootstrap
```

函数服务部署成功后，通过APIGW开放域名访问。
