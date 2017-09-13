package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/lumanetworks/go-tcp-proxy"
)

// Configuration 配置类
type Configuration struct {
	// IsLoad 加载标记
	IsLoad bool
	// ListenAddress 本地监听地址
	ListenAddress string `json:"listen_address"`
	// Verbose 日志是否显示详情
	Verbose bool `json:"verbose"`
	// Verbose 日志是否显示更多详情
	VeryVerbose bool `json:"very_verbose"`
	// Nagles 关闭Nagles算法
	Nagles bool `json:"nagles"`
	// Hex 是否以Hex形式显示数据
	Hex bool `json:"hex"`
}

var (
	config Configuration
	connid = uint64(0)
)

func main() {
	config.IsLoad = false

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("error:", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	err = decoder.Decode(&config)

	if err != nil {
		log.Println("error:", err)
	}

	config.IsLoad = true

	logger := proxy.ColorLogger{
		Verbose: config.Verbose,
		Color:   true,
	}

	logger.Info("Proxying to api.mch.weixin.qq.com:443.")
	laddr, err := net.ResolveTCPAddr("tcp", config.ListenAddress)
	if err != nil {
		logger.Warn("Failed to resolve local address: %s", err)
		os.Exit(1)
	}
	raddr, err := net.ResolveTCPAddr("tcp", "api.mch.weixin.qq.com:443")
	if err != nil {
		logger.Warn("Failed to resolve remote address: %s", err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		logger.Warn("Failed to open local port to listen: %s", err)
		os.Exit(1)
	}

	logger.Info("Config done. Start proxy, listening: %s", config.ListenAddress)
	for {

		conn, err := listener.AcceptTCP()
		if err != nil {
			logger.Warn("Failed to accept connection '%s'", err)
			continue
		}
		connid++
		var p *proxy.Proxy

		p = proxy.New(conn, laddr, raddr)
		p.Nagles = config.Nagles
		p.OutputHex = config.Hex
		p.Log = proxy.ColorLogger{
			Verbose:     config.Verbose,
			VeryVerbose: config.Verbose,
			Prefix:      fmt.Sprintf("Connection #%03d ", connid),
			Color:       true,
		}

		go p.Start()
	}
}
