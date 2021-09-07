package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"github.com/ph-piment/onion-scraper/models"
	"github.com/xo/dburl"
)

func main() {
	db, err := dburl.Open("postgresql://root:root@localhost:5432/os?sslmode=disable")
	if err != nil { /* ... */
		fmt.Printf("Error: %v", err)
		return
	}

	c := colly.NewCollector()

	now := time.Now()
	selector := ".newsFeed_list > li"
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		title := e.DOM.Find(".newsFeed_item_title").Text()
		sub := e.DOM.Find(".newsFeed_item_sub").Text()
		if len(title) == 0 {
			return
		}

		fmt.Printf("title: %v, sub: %v\n", title, sub)
		news := &models.News{
			Title:       title,
			Description: sub,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt: sql.NullTime{
				Time:  now,
				Valid: false,
			},
		}
		err := news.Insert(context.Background(), db)
		if err != nil {
			panic(err)
		}
	})

	c.Visit("https://news.yahoo.co.jp/topics/domestic")
}
