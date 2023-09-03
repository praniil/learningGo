package arrays

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	price string
}

func Csv() {
	file, err := os.Create("BookInfo.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Title", "Price"}
	writer.Write(headers)

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
		
	)
	
	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting to: ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println(r.Ctx.Get("url"))
	})

	c.OnHTML(".next > a", func(e *colly.HTMLElement){
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement){
		var book Book
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.price = e.ChildText(".price_color")
		array := []string{book.Title, book.price}
		writer.Write(array)
	})

	startUrl := "https://books.toscrape.com"
	c.Visit(startUrl)
}
