package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var baseURL string = "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=go+docs"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// 테이블 존재 여부 확인
	tableCount := doc.Find("div.sc_page").Length()
	fmt.Printf("Found %d tables with class LeaderBoardTable\n", tableCount)
	if tableCount == 0 {
		fmt.Println("테이블을 찾지 못했습니다. CSS 선택자를 다시 확인하세요.")
		return 0
	}

	fmt.Println("네이버에서 go docs 검색 결과")
	doc.Find("div.sc_page div.sc_page_inner a").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
