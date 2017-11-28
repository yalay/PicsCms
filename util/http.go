package util

import (
	"net/http"
	"strings"
)

func headerIp(headerKey string, r *http.Request) string {
	headIps := r.Header.Get(headerKey)
	if headIps == "" {
		return ""
	}

	headIpFields := strings.Split(headIps, ",")
	for _, headIp := range headIpFields {
		if headIp = strings.TrimSpace(headIp); headIp != "" && headIp != "unknown" {
			return headIp
		}
	}
	return ""
}

// 请求头里面记录的代理起始发起端ip
func HeaderForwardedIp(r *http.Request) string {
	return headerIp("X-FORWARDED-FOR", r)
}

// 请求头里面记录的realIp
func HeaderRealIp(r *http.Request) string {
	return headerIp("X-REAL-IP", r)
}

// 获取用户真实ip，优先用forwardedIp（起始发起端ip），realIp是客户端实际ip，RemoteAddr是负载均衡服务器的ip
func RealIp(r *http.Request) string {
	forwardedIp := HeaderForwardedIp(r)
	if forwardedIp != "" {
		return forwardedIp
	}

	realIp := HeaderRealIp(r)
	if realIp != "" {
		return realIp
	}

	remoteAddr := r.RemoteAddr
	if remoteAddr != "" {
		remoteAddr = strings.Split(remoteAddr, ":")[0]
		return remoteAddr
	}
	return "0.0.0.0"
}
