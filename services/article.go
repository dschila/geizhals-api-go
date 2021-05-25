package services

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/proph/geizhals-api-go/models"
)

// Visit the page and extract article informations and all availabel provider
func ArticleDetails(identifier string) (models.Article, []models.Provider) {
	article := models.Article{}
	providers := []models.Provider{}
	s := colly.NewCollector()
	s.OnHTML("html", func(h *colly.HTMLElement) {
		article.Name = h.ChildText(".variant__header__headline")
		prices := h.ChildTexts(".variant__header__pricehistory__pricerange > strong")
		if len(prices) == 2 {
			article.PriceFrom = prices[0]
			article.PriceTo = prices[1]
		}
		h.ForEach("div[id^='offer__']", func(i int, h *colly.HTMLElement) {
			provider := models.Provider{
				Name:         h.ChildText(".merchant__logo-caption"),
				Price:        h.ChildText(".gh_price"),
				OfferURL:     fmt.Sprintf("https://geizhals.de/%s", h.ChildAttr(".offer__clickout > a", "href")),
				Availability: h.ChildText(".offer__delivery-time"),
				Shipping:     h.ChildText(".offer__delivery-payment"),
			}
			providers = append(providers, provider)
		})

		if img, found := h.DOM.Find(".gal_img").Attr("src"); found {
			article.ImageURL = img
		}
	})

	s.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("https://geizhals.de/%s", identifier)
	s.Visit(url)

	return article, providers
}
