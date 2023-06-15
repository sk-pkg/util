package util

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

// GetRealIP 返回客户端真实IP
func GetRealIP(c *gin.Context) string {
	var ip string

	ip = c.Request.Header.Get("X-Original-Forwarded-For")
	if ip != "" && ip != "127.0.0.1" {
		ip = strings.Split(ip, ",")[0] // 反向代理时ip地址会自动拼接代理IP，获取真实客户端IP
		return ip
	}

	ip = c.Request.Header.Get("X-Forwarded-For")
	if ip != "" {
		ip = strings.Split(ip, ",")[0] // 反向代理时ip地址会自动拼接代理IP，获取真实客户端IP
		return ip
	}

	ip = c.Request.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
	if err != nil {
		return ""
	}

	return ip
}
