package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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
	var posts []extractedMonzee
	totalPages := getPages()

	for i := 1; i <= totalPages; i++ {
		extractedPosts := getPage(i)
		posts = append(posts, extractedPosts...)
	}
	writePosts(posts)
}

func writePosts(posts []extractedMonzee) {
	file, err := os.Create("posts.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "SUMMARY", "DATE"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, post := range posts {
		postSlice := []string{"https://monzee.tistory.com" + post.id, post.title, post.summary, post.date}
		pwErr := w.Write(postSlice)
		checkErr(pwErr)
	}
}

func getPage(page int) []extractedMonzee {
	var posts []extractedMonzee
	c := make(chan extractedMonzee)
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
		go extractPost(card, c)
	})

	for i:=0; i<searchCards.Length(); i++{
		post := <-c
		posts = append(posts, post)
	}

	return posts

}

func extractPost(card *goquery.Selection, c chan<- extractedMonzee) {
	id, _ := card.Find(".article-content>a").Attr("href")
	title := card.Find(".title").Text()
	summary := card.Find(".summary").Text()
	date := card.Find(".date").Text()
	c <- extractedMonzee{id: id, title: title, summary: summary, date: date}
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
