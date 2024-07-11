package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func generateString() string {
	currentTime := time.Now()
	maxDate := "29.09.2024"
	userDate := currentTime.Format("02.01.2006")
	var stringBuild strings.Builder = strings.Builder{}
	stringBuild.WriteString("https://www.itaka.pl/wyniki-wyszukiwania/wakacje/?dateFrom=")
	stringBuild.WriteString(userDate)
	stringBuild.WriteString("&dateTo=")
	stringBuild.WriteString(maxDate)
	stringBuild.WriteString("&priceFrom=2500&priceTo=3500&order=priceAsc&page=1&destinations=albania%2Cbulgaria%2Cgrecja%2Cegipt%2Cturcja%2Chiszpania&participants%5B0%5D%5Badults%5D=2")
	return stringBuild.String()
}

func main() {
	url := generateString()
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(5),
	)
	c.OnHTML("div.styles_c__f1i9i", func(e *colly.HTMLElement) {
		location := strings.Builder{}
		e.ForEach("div.styles_destination__tOoSF", func(i int, el *colly.HTMLElement) {
			location.WriteString(el.Text)
			location.WriteString(", ")
		})

		title := e.ChildText("h5.styles_title__kH0gG")
		price := e.ChildText("span.styles_current-price__value__NY1hb")
		fmt.Println(location.String(), title, price)

	})

	c.Visit(url)
	c.Wait()
}
