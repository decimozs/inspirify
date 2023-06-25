package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("\nStatus Code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("\nError: ", err.Error())
	})

	var quote, author string

	// scrape item from the web
	c.OnHTML(".text", func(h *colly.HTMLElement) {
		quote = h.Text
		// Remove quotation marks from the quote
		quote = strings.ReplaceAll(quote, "“", "")
		quote = strings.ReplaceAll(quote, "”", "")
	})

	c.OnHTML(".author", func(h *colly.HTMLElement) {
		author = h.Text

		// Open the file in and overwrite the existing quote and author (clearing existing content)
		file, err := os.OpenFile("quote.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Write the quote and author to the file
		_, err = fmt.Fprintf(file, "Quote: %s\nAuthor: %s\n\n", quote, author)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	})

	c.Visit("http://quotes.toscrape.com/random")
}
