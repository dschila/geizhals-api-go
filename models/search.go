package models

type SearchArticle struct {
	Name         string
	LowestPrice  float64
	OfferCount   int
	URL          string
	ImageURL     string
	Availability string
}
