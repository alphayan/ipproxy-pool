package ip89

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alphayan/ipproxy-pool/fetcher"
	"github.com/alphayan/ipproxy-pool/models/ipModel"
	"github.com/alphayan/ipproxy-pool/util"
	logger "github.com/sirupsen/logrus"
)

func Ip89() []*ipModel.IP {
	logger.Info("[89ip] fetch start")
	defer func() {
		recover()
		logger.Warnln("[89ip] fetch error")
	}()
	list := make([]*ipModel.IP, 0)

	indexUrl := "https://www.89ip.cn/"
	fetchIndex := fetcher.Fetch(indexUrl)
	pageNum := fetchIndex.Find("#layui-laypage-1 > a:nth-child(7)").Text()
	num, _ := strconv.Atoi(pageNum)
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf("%s/index_%d.html", indexUrl, i)
		fetch := fetcher.Fetch(url)
		fetch.Find("table > tbody").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
				proxyIp := strings.TrimSpace(selection.Find("td:nth-child(1)").Text())
				proxyPort := strings.TrimSpace(selection.Find("td:nth-child(2)").Text())
				proxyLocation := strings.TrimSpace(selection.Find("td:nth-child(3)").Text())

				ip := new(ipModel.IP)
				ip.ProxyHost = proxyIp
				ip.ProxyPort, _ = strconv.Atoi(proxyPort)
				ip.ProxyType = "http"
				ip.ProxyLocation = proxyLocation
				ip.ProxySpeed = 100
				ip.ProxySource = "https://www.89ip.cn"
				ip.CreateTime = util.FormatDateTime()
				ip.UpdateTime = util.FormatDateTime()
				list = append(list, ip)
			})
		})
	}
	logger.Info("[89ip] fetch done")
	return list
}
