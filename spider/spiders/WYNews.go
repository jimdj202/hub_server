package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/models"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Spider) GetWYNews() []models.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []models.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "http://news.163.com/special/0001386F/rank_whole.html"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	//request.Header.Add("Referer", `https://www.douban.com/group/explore`)
	//request.Header.Add("Host", `www.douban.com`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find("div.left div.active table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td a").First()
		url, boolUrl := s.Attr("href")
		textUtf8, _ := GbkToUtf8([]byte(s.Text()))

		//url, boolUrl := selection.Find("h3 a").Attr("href")
		//text := selection.Find("h3 a").Text()
		comNum := selection.Find(".cBlue").Text()
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		//imgUrl, _ := selection.Find(".pic-wrap img").Attr("src")
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := models.Item{
				Index: i,
				Title:      string(textUtf8),
				Url:        url,
				ImageUrl:   "",
				TypeDomain: "网易新闻",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: comNum3 ,
				//Date:       time.Time{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
