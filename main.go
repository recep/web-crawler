package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {

	url := flag.String("url", "http", "website's url")
	flag.Parse()

	urls, err := urlParse(*url)
	if err != nil {
		log.Println(err)
	}

	reqUrls := strings.Fields(urls)

	fmt.Printf("%-30s %-10s\n", "Url", "Status Code")
	fmt.Println(strings.Repeat("-", 45))

	for _, uri := range reqUrls {
		resp, err := http.Get(uri)
		if err != nil {
			continue
		}

		code := resp.StatusCode
		c := strconv.Itoa(code)

		fmt.Printf("%-30s", uri)

		switch code {
		case 200, 201, 202, 203, 204, 205, 206, 207, 208, 226:
			color.Blue(c)
		case 300, 301, 302, 303, 304, 305, 306, 307, 308:
			color.Yellow(c)
		case 400, 401, 403, 404, 405, 406, 412, 415:
			color.Red(c)
		case 500, 501:
			color.Red(c)
		}

	}
}

func urlParse(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var titles string

	doc.Find("[href]").Each(func(i int, s *goquery.Selection) {
		Link, _ := s.Attr("href")
		titles += Link + "\n"
	})

	return titles, nil
}
