package qingyunke

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
)

var (
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func getPicture() string {
	prefix:="https://www.fabiaoqing.com/search/bqb/keyword/%E5%88%AB%E5%8F%91%E5%9B%BE%E7%89%87/type/bq/page/"
	url :=fmt.Sprintf("%d.html",rand.Intn(5)+1)
	url = prefix + url
	log.Println("正在"+url+"寻找图片")
	urls := getImgs(url)
	fmt.Println(urls)
	imageUrl := urls[rand.Intn(len(urls))]
	log.Println("取到"+imageUrl)
	return imageUrl
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}
