package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/models"
	"io"
	"net/http"
	"strings"
	"time"
)

func (s *Spider) GetSegmentFault() []models.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []models.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://segmentfault.com/hottest"
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
	document.Find(".news-list .news__item-info").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a:nth-child(2)").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("h4").Text()

		descText := s.Find("div.article-excerpt").Text()
		//comNum ,_:= selection.Find("span.icon-message").Attr("data-origincount")
		////comNum = strings.ReplaceAll(comNum,",","")
		//reg, _ := regexp.Compile("\\d+")
		//comNum2 := reg.Find([]byte(comNum))
		//comNum3, _ := strconv.Atoi(string(comNum2))
		imgUrl, _ := selection.Find("a:nth-child(1)").Attr("style")
		if len(imgUrl) > 0{
			imgUrl = strings.Replace(imgUrl,"background-image:url(","",1)
			imgUrl = strings.Replace(imgUrl,")","",1)
		}
		//extra := selection.Find("div.f6 span span").Text()
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := models.Item{
				Index: i,
				Title:      text,
				Url:        "https://segmentfault.com" + url,
				ImageUrl:   imgUrl,
				TypeDomain: "SegmentFault",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: 0 ,
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
