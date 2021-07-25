package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dschila/geizhals-api-go/helpers"
	"github.com/dschila/geizhals-api-go/models"
	"github.com/gocolly/colly/v2"
)

// Searches for articles based on the query string and the category ID
func Search(query string, category int) []models.SearchArticle {
	articles := []models.SearchArticle{}
	s := colly.NewCollector()
	s.OnHTML(".listview__item > .listview__content", func(h *colly.HTMLElement) {
		isVariant := h.ChildText(".listview__label.listview__label--variant")
		if isVariant == "" {
			article := models.SearchArticle{
				Name:        h.ChildText(".listview__name-wrapper > h3"),
				URL:         h.ChildAttr(".listview__name-wrapper > h3 > a", "href"),
				LowestPrice: helpers.ConvertStringToFloat(h.ChildText(".listview__price-link")),
				ImageURL:    h.ChildAttr(".listview__image-link > img", "src"),
			}

			offerCountArray := regexp.MustCompile("[0-9]+").FindAllString(h.ChildText(".listview__offercount"), -1)
			if len(offerCountArray) > 0 {
				offer, err := strconv.Atoi(offerCountArray[0])
				if err != nil {
					offer = 0
				}
				article.OfferCount = offer
			} else {
				article.OfferCount = 0
			}

			aclasses := h.ChildAttr(".listview__content-right > span", "class")
			aclassarr := strings.Split(aclasses, " ")
			for i := 0; i < len(aclassarr); i++ {
				article.Availability = models.AvailabilityFromString[aclassarr[i]]
			}

			articles = append(articles, article)
		}
	})

	s.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("https://geizhals.de/?fs=%s&hloc=de&in=%d", query, category)
	s.Visit(url)

	return articles
}
