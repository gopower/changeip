package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func getIp(c *gin.Context) {
	// 打开yaml文件
	yamlConfig, err := os.ReadFile("interfaces.yaml")
	if err != nil {
		slog.Error(" os.ReadFile error:", err)
		return
	}

	// 转换 the YAML string to a struct
	var config NetworkConfig
	err = yaml.Unmarshal([]byte(yamlConfig), &config)
	if err != nil {
		slog.Error("yaml.Unmarshal error:", err)
		return
	}

	// 打印网卡名称 ip 子网掩码
	var addrs []address
	for key, val := range config.Network.Ethernets {
		ipSlice := strings.Split(val.Addresses[0], "/")
		netmask, err := strconv.Atoi(ipSlice[1])
		if err != nil {
			slog.Error(" strconv.Atoi(ipSlice[1]) error:", err)
			return
		}
		addrs = append(addrs, address{IFName: key, Ipv4: ipSlice[0], Netmask: netmask})
	}

	//返回 网卡名称 ip 子网掩码
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   addrs,
	})
}
