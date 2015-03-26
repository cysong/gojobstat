package model

type Seek struct {
	Title      string
	TotalCount int
	Data       []SeekJob
}

type SeekJob struct {
	Id                int
	ListDate          string
	Title             string
	Teaser            string
	BulletPoints      []string
	Advertiser        Advertiser
	Logo              Logo
	IsPremium         bool
	IsStandOut        bool
	Location          string
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
