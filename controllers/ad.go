package controllers

import (
	"sync"
	"strings"
)

const (
	kAdViewsNum = 16
)

var adReport *AdReport

type AdReport struct {
	sync.RWMutex
	ipCounts map[string]int
}

func init() {
	adReport = &AdReport{
		ipCounts: make(map[string]int),
	}
}

func NeedAd(ip, ua string) bool {
	ua = strings.ToLower(ua)
	if strings.Contains(ua, "spider") ||
		strings.Contains(ua, "googlebot") {
			return false
	}
	return adReport.needAd(ip)
}

func (a *AdReport) needAd(ip string) bool {
	a.Lock()
	defer a.Unlock()
	curNum, _ := a.ipCounts[ip]
	if curNum >= kAdViewsNum {
		a.ipCounts[ip] = 0
		return true
	}

	a.ipCounts[ip] = curNum + 1
	return false
}

