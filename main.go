package main

import (
	"log"
)

var (
	baseUrl string
)

func init() {
	baseUrl = "http://www.seek.co.nz/jobs-in-information-communication-technology/developers-programmers/in-new-zealand/#dateRange=999&workType=0&industry=6281&occupation=6287&graduateSearch=false&salaryFrom=0&salaryTo=999999&salaryType=annual&advertiserID=&advertiserGroup=&keywords=&page=1&displaySuburb=&seoSuburb=&isAreaUnspecified=false&location=&area=&nation=3001&sortMode=ListedDate&searchFrom=quick&searchType="
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
