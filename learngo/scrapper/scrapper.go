package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	title    string
	location string
	work     string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=" + term
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalpages := getPages(baseURL)

	for i := 1; i <= totalpages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalpages; i++ {
		extractedjobs := <-c
		jobs = append(jobs, extractedjobs...)
	}
	writeJobs(jobs)
	fmt.Println("Finished, extracted:", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Title", "location", "Work"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobslice := []string{job.title, job.location, job.work}
		jwErr := w.Write(jobslice)
		checkErr(jwErr)
	}
}

func getPage(i int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&recruitPage=" + strconv.Itoa(i)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".area_job")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	title := card.Find(".job_tit>a").Text()
	location := card.Find(".job_condition>span>a").Text()
	work := card.Find(".job_sector>a").Text()
	c <- extractedJob{title: title, location: location, work: work}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	var pages int
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
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
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
