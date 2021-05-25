package models

type Article struct {
	Name      string
	PriceFrom string
	PriceTo   string
	ImageURL  string
}

type Provider struct {
	Name         string
	OfferURL     string
	Price        string
	Availability string
	Shipping     string
}
