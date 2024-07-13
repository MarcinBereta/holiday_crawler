package main

import (
	"holiday_crawler/crawlers"
	"html/template"
	"os"
)
func main() {
	country := os.Args[1]
	var allOffers []crawlers.Offer
	wakacjeOffers := crawlers.GetWakacjeOffers(country)
	itakaOffers := crawlers.GetItakaOffers(country)
	rainbowOffers := crawlers.GetRainbowOffers(country)

	allOffers = append(allOffers, wakacjeOffers...)
	allOffers = append(allOffers, itakaOffers...)
	allOffers = append(allOffers, rainbowOffers...)

	var tmplFile = "offersHTML.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	var f *os.File
	f, err = os.Create("offers.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, allOffers)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
