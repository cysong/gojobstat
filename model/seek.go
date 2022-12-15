package model

type Seek struct {
	Title string    `json:"title"`
	Count int       `json:"totalCount"`
	Data  []SeekJob `json:"data"`
}

type SeekJob struct {
	Id                int
	ListingDate       string
	Title             string
	Teaser            string
	BulletPoints      []string
	Advertiser        Advertiser
	Logo              Logo
	IsPremium         bool
	IsStandOut        bool
	Location          string
	Area              string
	Classification    Classiciation
	SubClassification SubClassiciation
	Salary            string
	Suburb            string
}

type Advertiser struct {
	Id          string
	Description string
}

type Logo struct {
	Id          string
	Description string
}

type Classiciation struct {
	Id          string
	Description string
}

type SubClassiciation struct {
	Id          string
	Description string
}
