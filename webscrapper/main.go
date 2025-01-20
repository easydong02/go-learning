package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

var baseURL string = "https://search.naver.com/search.naver?nso=&page=2&query=go+docs&sm=tab_pge&"

type extractedResult struct {
	url string
}

func main() {
	var allUrls []extractedResult
	pages := getPages()
	fmt.Println(pages)

	for i := 0; i < pages/2; i++ {
		urls := getPage(i)
		allUrls = append(allUrls, urls...)
	}

	writeUrls(allUrls)
}

func writeUrls(urls []extractedResult) {
	file, err := os.Create("urls.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"URL"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, url := range urls {
		urlSlice := []string{url.url}
		uwErr := w.Write(urlSlice)
		checkErr(uwErr)
	}
}

func getPage(pageNumber int) []extractedResult {
	var results []extractedResult
	pageURL := baseURL + fmt.Sprintf("start=%d&where=web", 1+(15*pageNumber))
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	finds := doc.Find("a.link_tit")
	finds.Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Attr("href")
		if exists {
			result := extractedResult{url: val}
			results = append(results, result)
		}
	})

	fmt.Println("------------------------------------------")
	return results
}

func getPages() int {
	pages := 0
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
	doc.Find("div.sc_page div.sc_page_inner").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
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
