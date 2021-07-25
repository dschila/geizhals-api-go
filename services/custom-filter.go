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

// Searches for articles based on the query url
func CustomFilter(queryUrl string) []models.SearchArticle {
	articles := []models.SearchArticle{}
	s := colly.NewCollector()
	s.OnHTML(".productlist__product", func(h *colly.HTMLElement) {
		article := models.SearchArticle{
			Name:        h.ChildText(".productlist__link"),
			URL:         h.ChildAttr(".productlist__name > a", "href"),
			LowestPrice: helpers.ConvertStringToFloat(h.ChildText(".productlist__price > span > span")),
			ImageURL:    h.ChildAttr(".catitem__image__container > picture > img", "src"),
		}

		offerCountArray := regexp.MustCompile("[0-9]+").FindAllString(h.ChildText(".productlist__offerscount--standard"), -1)
		if len(offerCountArray) > 0 {
			offer, err := strconv.Atoi(offerCountArray[0])
			if err != nil {
				offer = 0
			}
			article.OfferCount = offer
		} else {
			article.OfferCount = 0
		}

		aclasses := h.ChildAttr(".productlist__deliverytime", "class")
		aclassarr := strings.Split(aclasses, " ")
		for i := 0; i < len(aclassarr); i++ {
			article.Availability = models.AvailabilityFromString[aclassarr[i]]
		}

		articles = append(articles, article)
	})

	s.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("https://geizhals.de/?%s", queryUrl)
	s.Visit(url)

	return articles
}
