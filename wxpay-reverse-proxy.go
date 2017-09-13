package main

import (
	"github.com/lumanetworks/go-tcp-proxy"
)

func main() {
	logger := proxy.ColorLogger{
		Verbose: true,
		Color:   true,
	}

	logger.Info("Proxying to api.mch.weixin.qq.com.")
}

func init() {
}
