package models

type SearchArticle struct {
	Name         string
	LowestPrice  float64
	OfferCount   int
	URL          string
	ImageURL     string
	Availability Availability
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
