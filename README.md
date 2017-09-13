# wxpay-reverse-proxy
支持客户端证书验证转发的tcp代理

基本就是重新打包了一下的 github.com/jpillora/go-tcp-proxy.

## 使用说明

go install github.com/jpillora/go-tcp-proxy

或者windows版直接下载

https://github.com/btbxbob/wxpay-reverse-proxy/releases

注意运行目录需要有config.json。

启动后像这样进行测试即可：

```
curl -k -d "@add_merchant2.txt" --key "apiclient_key.pem" --cert "apiclient_cert.pem"  -H "Host: api.mch.weixin.qq.com" https://localhost:9999/secapi/mch/submchmanage?action=add
```