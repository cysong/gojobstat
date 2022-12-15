package fetchers

import (
	"encoding/json"
	"github.com/gojobstat/model"
	"github.com/gojobstat/statutil"
	"log"
)

type SeekFether struct{}

func (s *SeekFether) Fetch(url string) (resultJson []model.SeekJob, err error) {
	page := GetByUrl(url) // the raw page of the url
	var result model.Seek
	err = json.Unmarshal(page, &result)
	if err != nil {
		log.Println("error: ", err)
	}

	//return statJobs(result), nil
	return result.Data, err
}

func statJobs(seek model.Seek) (jsStr string) {
	m := make(map[string]int) // programming languages -> int, all pl is in lower cases
	for _, data := range seek.Data {
		r := statutil.Analyse(data.Title)
		log.Println("---------> return map:", len(r))
		if r != nil {
			for k, v := range r {
				if m[k] != -1 {
					m[k] = m[k] + v
				} else {
					m[k] = v
				}
			}
		}
	}

	str, err := json.Marshal(m)
	if err != nil {
		log.Println("error: ", err)
	}
	return string(str)
}
