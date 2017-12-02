package controllers

import (
	"strings"
	"sync"
)

const (
	kAdViewsBaseNum = 4
	kAdViewsMinNum  = 12
	kAdViewsMaxNum  = 24
)

const (
	kAdCode = `
		<div style="margin: 0 auto; width: 640px; display: block;">
		<iframe src="//ads.exosrv.com/iframe.php?idzone=2469239&size=315x300" width="315" height="300" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>
		<iframe src="//ads.exosrv.com/iframe.php?idzone=2469239&size=315x300" width="315" height="300" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>
		</div>
	`
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

func GetAd() string {
	return kAdCode
}

func (a *AdReport) needAd(ip string) bool {
	a.Lock()
	defer a.Unlock()
	curNum, _ := a.ipCounts[ip]
	if curNum >= kAdViewsMaxNum {
		a.ipCounts[ip] = 0
		return true
	}

	a.ipCounts[ip] = curNum + 1
	if curNum >= kAdViewsMinNum {
		if curNum%kAdViewsBaseNum == 0 {
			return true
		}
	}

	return false
}
