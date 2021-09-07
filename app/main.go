package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	selector := ".newsFeed_list > li"
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		title := e.DOM.Find(".newsFeed_item_title").Text()
		sub := e.DOM.Find(".newsFeed_item_sub").Text()
		if len(title) == 0 {
			return
		}
		fmt.Printf("title: %v, sub: %v\n", title, sub)
	})

	c.Visit("https://news.yahoo.co.jp/topics/domestic")
}
