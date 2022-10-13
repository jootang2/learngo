package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedMonzee struct {
	id      string
	title   string
	summary string
	date    string
}

var baseURL string = "https://monzee.tistory.com/"

func main() {
	var jobs []extractedMonzee
	totalPages := getPages()

	for i := 1; i <= totalPages; i++ {
		extractedjobs := getPage(i)
		jobs = append(jobs, extractedjobs...)
	}
	fmt.Println(jobs)


}

func getPage(page int) []extractedMonzee{
	var posts []extractedMonzee
	pageURL := baseURL + "?page=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".article-content")

	searchCards.Each(func(i int, card *goquery.Selection) {
		post := extractPost(card)
		posts = append(posts, post)
	})

	return posts

}

func extractPost(card *goquery.Selection) extractedMonzee{
	id, _ := card.Attr("href")
	title := card.Find(".title").Text()
	summary := card.Find(".summary").Text()
	date := card.Find(".date").Text()
	return extractedMonzee{id: id, title: title, summary: summary, date: date}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".area-paging").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() - 2
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
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
