package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type Image struct {
	PictUrl string
	Title string
}

func main(){
	image:=Image{}
	scraper:=newScraper()
	scraper.OnHTML(`meta[property='og:image']`, func(e *colly.HTMLElement) {
		image.PictUrl=e.Attr("content")
	})
	scraper.OnHTML(`meta[property='og:title']`, func(e *colly.HTMLElement) {
		image.Title=e.Attr("content")
	})
	if os.Args[1]==""{
		log.Fatal("Paste the instagram url")
	}
	err:=scraper.Visit(os.Args[1])
	if err!=nil{
		log.Fatal(err)
	}
	download(scraper,image)
}

func download(scraper *colly.Collector, image Image){
	scraper.OnResponse(func(response *colly.Response) {
		picture, err:=os.Create(image.Title+".jpg")
		if err!=nil{
			log.Fatal(err)
		}
		defer picture.Close()
		n2,err:=picture.Write(response.Body)
		fmt.Printf("Wrote %d bytes\n", n2)
		picture.Sync()
		})
	scraper.Visit(image.PictUrl)
}

func newScraper() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"
	return c
}