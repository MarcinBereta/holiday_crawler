package crawlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func generateItakaString(country string) string {
	currentTime := time.Now()
	maxDate := "31.07.2024"
	userDate := currentTime.Format("02.01.2006")
	var stringBuild strings.Builder = strings.Builder{}
	stringBuild.WriteString("https://www.itaka.pl/all-inclusive/")
	stringBuild.WriteString(country)
	stringBuild.WriteString("/?dateFrom=")
	stringBuild.WriteString(userDate)
	stringBuild.WriteString("&dateTo=")
	stringBuild.WriteString(maxDate)
	stringBuild.WriteString("&durationMax=8&durationMin=6&priceTo=3500&order=reviewsRatingDesc&page=1&participants=0&adults=2&priceFrom=2500&departuresByPlane=KTW%2CKRK")

	return stringBuild.String()
}

func GetItakaOffers(country string) []Offer{
	url := generateItakaString(country)
	c := colly.NewCollector()

	var offers []Offer
	c.OnHTML("div.styles_c__f1i9i", func(e *colly.HTMLElement) {
		location := strings.Builder{}
		e.ForEach("div.styles_destination__tOoSF", func(i int, el *colly.HTMLElement) {
			location.WriteString(el.Text)
			location.WriteString(", ")
		})

		name := e.ChildText("h5.styles_title__kH0gG")
		price := e.ChildText("span.styles_current-price__value__NY1hb")
		departureDate := e.ChildText("div.styles_c__GqLxf")
		departurePlace := e.ChildText("div.styles_label___8Mr4")
		ratingValue := e.ChildText("span.styles_c__rIHSD")
		link := e.ChildAttr("a.styles_c__MESiM", "href")

		offer := Offer{
			Name: name,
			Price: price,
			DepartureDate: departureDate,
			DeparturePlace: departurePlace,
			RatingValue: ratingValue,
			Link: fmt.Sprintf("https://itaka.pl%v", link),
		}

		offers = append(offers, offer)

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.Visit(url)

	return offers
}