package models

type Article struct {
	Name      string
	PriceFrom float64
	PriceTo   float64
	ImageURL  string
}

type Provider struct {
	Name         string
	OfferURL     string
	Price        float64
	Availability string
	Shipping     string
}
