package main

import (
	"github.com/gojobstat/fetchers"
)

func main() {
	sk := fetchers.SeekFether{}
	sk.Fetch("https://api.seek.com.au/v2/jobs/search?keywords=&hirerId=&hirerGroup=&page=1&classification=6281&subclassification=6287&graduateSearch=false&displaySuburb=&suburb=&location=&nation=3001&area=&isAreaUnspecified=false&worktype=&salaryRange=0-999999&salaryType=annual&dateRange=999&sortMode=ListedDate&engineConfig=&usersessionid=owzcdtqlasq144424pmlhfhx&eventCaptureSessionId=e1847392-9b47-4535-b278-21952c751c43&userqueryid=138283483704875751&include=expanded&_=1427325069071")

}

// https://api.seek.com.au/v2/jobs/search?keywords=&hirerId=&hirerGroup=&page=1&classification=6281&subclassification=6287&graduateSearch=false&displaySuburb=&suburb=&location=&nation=3001&area=&isAreaUnspecified=false&worktype=&salaryRange=0-999999&salaryType=annual&dateRange=999&sortMode=ListedDate&engineConfig=&usersessionid=owzcdtqlasq144424pmlhfhx&eventCaptureSessionId=e1847392-9b47-4535-b278-21952c751c43&userqueryid=138283483704875751&include=expanded&_=1427325069071
