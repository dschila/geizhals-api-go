package models

type Article struct {
	Name      string  `json:"name"`
	PriceFrom float64 `json:"priceFrom"`
	PriceTo   float64 `json:"priceTo"`
	ImageURL  string  `json:"imageUrl"`
}

type Provider struct {
	Name         string  `json:"name"`
	OfferURL     string  `json:"offerUrl"`
	Price        float64 `json:"price"`
	Availability string  `json:"availability"`
	Shipping     string  `json:"shipping"`
}
