package main

import (
	"encoding/json"
	"errors"
	"github.com/gojobstat/fetchers"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	baseUrl string

	// sentinel value
	loading = errors.New("url load in progress")

	// fetched tracks URLs that have been (or are being) fetched.
	// The lock must be held while reading from or writing to the map.
	// See http://golang.org/ref/spec#Struct_types section on embedded types.
	fetched = struct {
		m      map[string]error
		result map[string]int
		sync.Mutex
	}{m: make(map[string]error), result: make(map[string]int)}
)

func init() {
	baseUrl = "https://api.seek.com.au/v2/jobs/search?keywords=&hirerId=&hirerGroup=&classification=6281&subclassification=6287&graduateSearch=false&displaySuburb=&suburb=&location=&nation=3001&area=&isAreaUnspecified=false&worktype=&salaryRange=0-999999&salaryType=annual&dateRange=999&sortMode=ListedDate&engineConfig=&usersessionid=owzcdtqlasq144424pmlhfhx&eventCaptureSessionId=e1847392-9b47-4535-b278-21952c751c43&userqueryid=138283483704875751&include=expanded&_=1427325069071"
}

func main() {
	sk := &fetchers.SeekFether{}
	Crawl(1, sk)

	for k, v := range fetched.result {
		log.Println(k, ":", v)
	}
}

// the crawl is currentlly a single threaded function
func Crawl(paging int, fetcher fetchers.Fetcher) {
	body, err := fetcher.Fetch(baseUrl + "&page=" + strconv.Itoa(paging))
	if err != nil {
		log.Println("fetching error:", err)
		return
	}

	log.Println("============result============")
	log.Println(body)
	log.Println("============end result============")

	var tmpResult map[string]int
	err = json.Unmarshal([]byte(body), &tmpResult)
	if err != nil {
		log.Println("parsing json err:", err)
		return
	}
	if len(tmpResult) != 0 {
		// fetched.Lock()
		// add to result
		for k, v := range tmpResult {
			if count, ok := fetched.result[k]; ok {
				fetched.result[k] = count + v
			} else {
				fetched.result[k] = v
			}
		}
		// fetched.Unlock()

		Crawl(paging+1, fetcher)
	}

}

// this is a util function
func WriteToFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		log.Println("...")
		panic(err)
	}
}
