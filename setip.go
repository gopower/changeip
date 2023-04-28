package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

func setIp(c *gin.Context) {
	var addr address
	if err := c.ShouldBindJSON(&addr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ipStr := fmt.Sprintf("%s/%d", addr.Ipv4, addr.Netmask)

	// 打开yaml文件
	yamlConfig, err := os.ReadFile("interfaces.yaml")
	if err != nil {
		slog.Error("os.ReadFile error:", err)
		return
	}

	// 转换 the YAML string to a struct
	var config NetworkConfig
	err = yaml.Unmarshal([]byte(yamlConfig), &config)
	if err != nil {
		slog.Error("yaml.Unmarshal error:", err)
		return
	}

	config.Network.Ethernets[addr.IFName].Addresses[0] = ipStr

	// Convert the struct back to YAML
	updatedYaml, err := yaml.Marshal(&config)
	if err != nil {
		slog.Error("yaml.Marshal error:", err)
		return
	}

	err = os.WriteFile("interfaces.yaml", []byte(updatedYaml), 0644)
	if err != nil {
		slog.Error("os.WriteFile error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "修改ip中"})

}
