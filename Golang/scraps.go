package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	
	urls := []string{
		"https://en.wikipedia.org/wiki/Main_Page",
		"https://en.wikipedia.org/wiki/Encyclopedia",
		
	}

	
	c := colly.NewCollector()

	
	file, err := os.Create("scraped_data.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer file.Close()


	writer := csv.NewWriter(file)
	defer writer.Flush()


	writer.Write([]string{"Title", "Description", "Link"})

	
	for _, url := range urls {
		
		c.OnHTML("h1", func(e *colly.HTMLElement) {
		
			title := e.Text

			
			description := "" 
			link := url        

		
			writer.Write([]string{title, description, link})
		})

		
		c.OnError(func(r *colly.Response, err error) {
			log.Printf("Request URL: %s failed with response: %v and error: %v", r.Request.URL, r, err)
		})

	
		err := c.Visit(url)
		if err != nil {
			log.Printf("Error visiting URL %s: %v", url, err)
		}
	}

	fmt.Println("Scraping completed. Data saved to scraped_data.csv")
}
