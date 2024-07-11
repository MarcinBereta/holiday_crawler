package crawlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/thedevsaddam/gojsonq"
)

func reverse(s string) string { 
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    return string(rns) 
} 

func replaceF2(input string) string {
	return strings.ReplaceAll(input, "F2%", "/")
}

func removeLastNChars(s string, n int) string {
	if len(s) <= n {
		return ""
	}
	return s[:len(s)-n]
}

func generateWakacjeUrl() string {
    currentTime := time.Now()
    maxDate := "2024-07-31"
    userDate := currentTime.Format("2006-01-02")
    var stringBuild strings.Builder = strings.Builder{}
    stringBuild.WriteString("https://www.wakacje.pl/wczasy/grecja/?od-")
    stringBuild.WriteString(userDate)
    stringBuild.WriteString(",do-")
    stringBuild.WriteString(maxDate)
    stringBuild.WriteString(",6-8-dni,od-2500zl,do-2800zl,samolotem,all-inclusive,z-katowic,z-krakowa,ocena-malejaco,za-osobe")
    return stringBuild.String()
}

func GetWakacjeOffers() {
    url := generateWakacjeUrl()
    c := colly.NewCollector()
    c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		
		jq := gojsonq.New().JSONString(e.Text)

		for counter := 0; counter < 10; counter++ {
			path := "props.stores.storeOffers.offers.data.[" + strconv.Itoa(counter) + "]"
			keys := [8]string{"name", "price", "departureDate", "returnDate", "departurePlace", "ratingValue", "link", "duration"}
		
			name := jq.Copy().Find(path + "." + keys[0])
			price := jq.Copy().Find(path + "." + keys[1])
			departureData := jq.Copy().Find(path + "." + keys[2])
			returnDate := jq.Copy().Find(path + "." + keys[3])
			departurePlace := jq.Copy().Find(path + "." + keys[4])
			ratingValue := jq.Copy().Find(path + "." + keys[5])
			link := jq.Copy().Find(path + "." + keys[6])
			duration := jq.Copy().Find(path + "." + keys[7])

		
			if name == nil || price == nil || departureData == nil || returnDate == nil || departurePlace == nil || ratingValue == nil{
				break
			}
		
			fmt.Printf("Name: %v\n", name)
			fmt.Printf("Price: %v PLN\n", price)
			fmt.Printf("Departure data: %v\n", departureData)
			fmt.Printf("Return date: %v\n", returnDate)
			fmt.Printf("Departure place: %v\n", departurePlace)
			fmt.Printf("Rating value: %v\n", ratingValue)
			fmt.Printf("Duration: %v\n", duration)

			reversedLink := reverse(link.(string))
			reversedLink = removeLastNChars(reversedLink, 12)
			reversedLink = replaceF2(reversedLink)
			fmt.Printf("Link: https://wakacje.pl%v?od-%v,%v-dni,all-inclusive,z-%v", reversedLink, departureData, duration, departurePlace)


			fmt.Println()
		}

    })

    c.Visit(url)
}