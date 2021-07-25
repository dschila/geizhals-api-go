package models

type SearchArticle struct {
	Name         string       `json:"name"`
	LowestPrice  float64      `json:"lowestPrice"`
	OfferCount   int          `json:"offerCount"`
	URL          string       `json:"url"`
	ImageURL     string       `json:"imageUrl"`
	Availability Availability `json:"availability"`
}

type Availability int

const (
	ANY Availability = iota
	AVAILABLE
	SHORTLY
	UNKOWN
)

var AvailabilityFromString = map[string]Availability{
	"available": AVAILABLE,
	"shortly":   SHORTLY,
	"any":       ANY,
}
