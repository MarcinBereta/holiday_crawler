package main

import (
	"holiday_crawler/crawlers"
)

func main() {
	crawlers.GetWakacjeOffers()
	crawlers.GetItakaOffers()
	crawlers.GetRainbowOffers()
}
