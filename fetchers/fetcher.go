package fetchers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Fetcher interface {
	// Fetch returns the result in json format of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (resultJson string, urls []string, err error)
}

// this is a method to download the target url
// and return the content as a byte array
func GetByUrl(url string) []byte {
	log.Println("download start download url: ", url)
	response, err := http.Get(url)
	log.Println("download finished url: ", url)

	if err != nil {
		log.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		log.Println("content type:", response.Header.Get("content-type"))
		if err != nil {
			log.Printf("%s", err)
		}
		return contents
	}
	return nil
}
