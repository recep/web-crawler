package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {

	url := flag.String("url", "", "url of website")
	flag.Parse()

	if *url == "" {
		fmt.Println("Please give me a url")
		return
	}

	urls, err := urlParse(*url)
	if err != nil {
		fmt.Println(err)
		return
	}

	reqUrls := strings.Fields(urls)

	fmt.Printf("%-90s %s\n", "Url", "Status Code")
	fmt.Println(strings.Repeat("=", 100))

	for _, uri := range reqUrls {
		resp, err := http.Get(uri)
		if err != nil {
			continue
		}

		code := resp.StatusCode
		c := strconv.Itoa(code)

		fmt.Printf("%-90s", uri)

		switch code {
		case 200, 201, 202, 203, 204, 205, 206, 207, 208, 226:
			color.Blue(c)
		case 300, 301, 302, 303, 304, 305, 306, 307, 308:
			color.Yellow(c)
		case 400, 401, 403, 404, 405, 406, 412, 415:
			color.Red(c)
		case 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511:
			color.Red(c)
		default:
			fmt.Println()
		}
		fmt.Println(strings.Repeat("-", 100))

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
