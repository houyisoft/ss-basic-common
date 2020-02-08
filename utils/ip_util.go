package utils

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"net"
	"net/http"
	"os"
)

// 得到用户的IP
func GetIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

var ipTransfer *ip2region.Ip2Region

//得到ip地址相关信息, 格式: 国家-省-城市
func GetIpInfo(ipAddress string) string {
	var ipDbPath string
	pwd, pwdErr := os.Getwd()
	if pwdErr == nil {
		ipDbPath = pwd + "/config/ip2region.db"
	}
	if ipDbPath != "" {
		ipTransfer, _ = ip2region.New(ipDbPath)
		defer ipTransfer.Close()
		ip, ipErr := ipTransfer.BinarySearch(ipAddress)
		if ipErr != nil {
			return "解析失败"
		}
		if ip.Province == "0" && ip.Country == "0" && ip.Region == "" {
			return "未知地区"
		}
		ipStr := ""
		if ip.Country != "0" {
			ipStr += ip.Country
		}
		if ip.Province != "0" {
			ipStr += "-" + ip.Province
		}
		if ip.City != "0" {
			ipStr += "-" + ip.City
		}
		if ip.Region != "0" {
			ipStr += "-" + ip.Region
		}
		return ipStr
	}
	return "IP加载错误"
}
