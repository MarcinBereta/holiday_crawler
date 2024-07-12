package crawlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func generateTuiString() string {
	currentTime := time.Now()
	maxDate := "31.07.2024"
	userDate := currentTime.Format("02.01.2006")
	var stringBuild strings.Builder = strings.Builder{}
	stringBuild.WriteString("https://www.tui.pl/wypoczynek/wyniki-wyszukiwania-samolot?q=%3AhotelRating%3AbyPlane%3AT%3Aa%3AKTW%3Aa%3AKRK%3AdF%3A6%3AdT%3A8%3AstartDate%3A")
	stringBuild.WriteString(userDate)
	stringBuild.WriteString("%3AendDate%3A")
	stringBuild.WriteString(maxDate)
	stringBuild.WriteString("%3ActAdult%3A2%3ActChild%3A0%3Ac%3ACFU%3Ac%3AKGS%3Ac%3ACHQ%3Ac%3AMJT%3Ac%3AGPA%3Ac%3ARHO%3Ac%3AZTH%3AminHotelCategory%3AdefaultHotelCategory%3AtripAdvisorRating%3AdefaultTripAdvisorRating%3Abeach_distance%3AdefaultBeachDistance%3AtripType%3AWS&fullPrice=false")

	return stringBuild.String()
}

func GetTuiOffers() {
	url := generateTuiString()
	c := colly.NewCollector()
	c.OnHTML("div.offer-tile", func(e *colly.HTMLElement) {
		name := e.ChildText("span.offer-tile-body__hotel-name")
		price := e.ChildText("span.price-value__amount")
		departureDate := e.ChildText("span.offer-tile-body__info-item-text")
		departurePlace := e.ChildText("span.dropdown-field__value")
		link := e.ChildAttr("a.offer-tile-header", "href")

		fmt.Printf("Name: %v\n", name)
		fmt.Printf("Price: %v PLN\n", price)
		fmt.Printf("Departure place: %v\n", departurePlace)
		fmt.Printf("Departure date: %v\n", departureDate)
		fmt.Printf("Link: https://www.tui.pl%v\n", link)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Visit(url)
}