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

func (s *Spider) GetWangYiNews() []models.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []models.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
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

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td a").First()
		url, boolUrl := s.Attr("href")
		text, _ := GbkToUtf8([]byte(s.Text()))

		descText := selection.Find("a.article-list__abstract").Text()
		comNum := selection.Find("span.article-level1").Text()
		//comNum = strings.ReplaceAll(comNum,",","")
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		//imgUrl, _ := selection.Find("a:nth-child(1)").Attr("style")

		//extra := selection.Find("div.f6 span span").Text()
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := models.Item{
				Index: i,
				Title:      string(text),
				Url:        url,
				ImageUrl:   "",
				TypeDomain: "网易新闻",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: comNum3 ,
				Desc: descText,
				//Extra: extra,
				//Date:       time.Time{},

				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
