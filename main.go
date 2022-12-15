package main

import (
	"context"
	"github.com/gojobstat/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"

	//"encoding/json"
	"errors"
	"github.com/gojobstat/mongodb"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gojobstat/fetchers"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	baseUrl        = "https://api.seek.com.au/v2/jobs/search?keywords=&hirerId=&hirerGroup=&graduateSearch=false&displaySuburb=&suburb=&location=&nation=3001&area=&isAreaUnspecified=false&worktype=&salaryRange=0-999999&salaryType=annual&sortMode=ListedDate&engineConfig=&usersessionid=owzcdtqlasq144424pmlhfhx&eventCaptureSessionId=e1847392-9b47-4535-b278-21952c751c43&userqueryid=138283483704875751&include=expanded&_=1427325069071"
	dropTable      = false
	dateRange      = 2
	classification = 6281 //Information & Communication Technology
	total          = 0
	added          = 0
	// sentinel value
	loading = errors.New("url load in progress")
	conn    mongodb.Conn
	col     mongo.Collection
	ctx     context.Context
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
	parseArgs()
	conn := &mongodb.Conn{}
	col = conn.GetCol("seek", "jobs")
	if dropTable {
		col.Drop(ctx)
		log.Println("drop collection", conn.GetColFullName(col))
	}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	rows, _ := col.CountDocuments(ctx, bson.D{})
	baseUrl += "&classification=" + strconv.Itoa(classification)
	if rows == 0 {
		baseUrl += "&dateRange=999" //query all records
	} else {
		baseUrl += "&dateRange=" + strconv.Itoa(dateRange) //query last dateRange days
	}
}

func parseArgs() {
	if len(os.Args) > 1 {
		log.Println("args:", os.Args[1:])
		for _, arg := range os.Args[1:] {
			if arg == "dropTable" {
				dropTable = true
				dateRange = 999
			} else if strings.Index(arg, "dateRange") == 0 {
				strs := strings.Split(arg, "=")
				if len(strs) > 1 {
					dateRange, _ = strconv.Atoi(strs[1])
				}
			} else if strings.Index(arg, "class") == 0 {
				strs := strings.Split(arg, "=")
				if len(strs) > 1 {
					classification, _ = strconv.Atoi(strs[1])
				}
			} else {
				log.Println("unrecognized arg:", arg)
			}
		}
	}
}

func main() {
	defer conn.Close()
	sk := &fetchers.SeekFether{}
	Crawl(1, sk)

	//for k, v := range fetched.result {
	//	log.Println(k, ":", v)
	//}
	log.Printf("total records:%d, added:%d", total, added)
}

// the crawl is currentlly a single threaded function
func Crawl(paging int, fetcher fetchers.Fetcher) {
	body, err := fetcher.Fetch(baseUrl + "&page=" + strconv.Itoa(paging))
	if err != nil {
		panic(err)
	}
	total += len(body)
	if len(body) != 0 {
		for _, ele := range body {
			if ele.Id == 0 {
				log.Println("job id is 0")
				continue
			}
			if isJobExists(col, ele.Id) {
				continue
			}
			_, err := col.InsertOne(ctx, ele)
			if err != nil {
				panic(err)
			}
			added++
		}

		Crawl(paging+1, fetcher)
	}

	//log.Println(body)

	//var tmpResult map[string]int
	//err = json.Unmarshal([]byte(body), &tmpResult)
	//if err != nil {
	//	log.Println("parsing json err:", err)
	//	return
	//}
	//if len(tmpResult) != 0 {
	//	// fetched.Lock()
	//	// add to result
	//	for k, v := range tmpResult {
	//		if count, ok := fetched.result[k]; ok {
	//			fetched.result[k] = count + v
	//		} else {
	//			fetched.result[k] = v
	//		}
	//	}
	//	// fetched.Unlock()
	//
	//	Crawl(paging+1, fetcher)
	//}

}

func isJobExists(col mongo.Collection, jobId int) bool {
	count, _ := col.CountDocuments(ctx, bson.M{"id": jobId})
	return count > 0
}

func findByJobId(col mongo.Collection, jobId int) (model.SeekJob, error) {
	result := col.FindOne(ctx, bson.M{"id": jobId}, options.FindOne())
	var job model.SeekJob
	err := result.Err()
	if err != nil {
		return job, err
	}
	err = result.Decode(&job)
	return job, err
}

// this is a util function
func WriteToFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
}
