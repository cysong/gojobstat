package fetchers

import (
	"bytes"
	"golang.org/x/net/html"
	"log"
)

type SeekFether struct{}

func (s *SeekFether) Fetch(url string) (resultJson string, urls []string, err error) {
	page := GetByUrl(url) // the raw page of the url

	reader := bytes.NewReader(page)
	doc, err := html.Parse(reader)
	if err != nil {
		log.Println("error: ", err.Error())
		return "", nil, err
	}

	// now need process to get the body string and urls i need
	// and get the json string for the result
	str := parseSections(doc)
	log.Println(str)

	return str, nil, nil
}

func parseSections(doc *html.Node) (jobsList string) {
	f = func(n *html.Node) {
		if n != nil && n.Type == html.ElementNode && n.Data == "section" {
			for _, section := range n.Attr {
				if section.Key == "id" && section.Val == "jobsListing" {
					// here, we find a job
					// now start parsing it
					jobsList += parseArticle(n)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

func parseArticle(doc *html.Node) (jsStr string) {
	log.Println("========================Listing all the jobs ==============================")
	f = func(n *html.Node) {
		if n != nil && n.Type == html.DocumentNode && n.Data == "article" {
			jsStr += "find a article, "
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

var f func(*html.Node)
