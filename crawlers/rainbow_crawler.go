package crawlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func generateRainbowUrl(country string) string {
	currentTime := time.Now()
	maxDate := "2024.07.31"
	userDate := currentTime.Format("2006.01.02")
	var stringBuild strings.Builder = strings.Builder{}
	stringBuild.WriteString("https://r.pl/")
	stringBuild.WriteString(country)
	stringBuild.WriteString("?wybraneDokad=")
	stringBuild.WriteString(country)
	stringBuild.WriteString("&wybraneSkad=KTW&wybraneSkad=KRK&typTransportu=AIR&data=")
	stringBuild.WriteString(userDate)
	stringBuild.WriteString("&data=")
	stringBuild.WriteString(maxDate)
	stringBuild.WriteString("&dorosli=1994-01-01&dorosli=1994-01-01&dzieci=nie&liczbaPokoi=1&dowolnaLiczbaPokoi=nie&dataWylotu&dlugoscPobytu=*-*&dlugoscPobytu.od=6&dlugoscPobytu.do=8&cena=avg&cena.od=2500&cena.do=3500&ocenaKlientow=*-*&odlegloscLotnisko=*-*&hotelUrl&produktUrl&sortowanie=ocena-desc&strona=1&wyzywienia=all-inclusive")

	return stringBuild.String()
}

func GetRainbowOffers(country string) []Offer {
	url := generateRainbowUrl(country)
	c := colly.NewCollector()
	var offers []Offer
	c.OnHTML("a.n-bloczek", func(e *colly.HTMLElement) {
		name := e.ChildText("span.r-bloczek-tytul")
		price := e.ChildText("span.r-bloczek-cena__aktualna")
		departurePlace := e.ChildText("span.r-bloczek-przystanki__nazwa")
		departureDate := e.ChildText("div.r-bloczek-wlasciwosci__dni")
		link := e.Attr("href")

		offer := Offer{
			Name: name,
			Price: price,
			DepartureDate: departureDate,
			DeparturePlace: departurePlace,
			Link: fmt.Sprintf("https://r.pl%v", link),
		}

		offers = append(offers, offer)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Visit(url)

	return offers
}