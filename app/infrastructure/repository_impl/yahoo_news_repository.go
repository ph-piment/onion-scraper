package repository_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ph-piment/onion-scraper/app/domain/entity"
	"github.com/ph-piment/onion-scraper/app/domain/repository"
	"github.com/ph-piment/onion-scraper/app/infrastructure/dao"
)

type yahooNewsRepository struct {
	now time.Time
}

func NewYahooNewsRepository() repository.YahooNewsRepository {
	return &yahooNewsRepository{
		now: time.Now(),
	}
}

func (repo *yahooNewsRepository) ScrapingListFromWEB(ctx context.Context) ([]*entity.YahooNews, error) {
	c := colly.NewCollector()

	selector := ".newsFeed_list > li"
	result := make([]*entity.YahooNews, 0, 5)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		title := e.DOM.Find(".newsFeed_item_title").Text()
		sub := e.DOM.Find(".newsFeed_item_sub").Text()
		if len(title) == 0 {
			return
		}

		fmt.Printf("title: %v, sub: %v\n", title, sub)
		news := entity.NewYahooNews(0, title, sub)
		result = append(result, news)
	})

	c.Visit("https://news.yahoo.co.jp/topics/domestic")

	return result, nil
}

func (repo *yahooNewsRepository) ImportToDB(ctx context.Context, db *sqlx.DB, rows []*entity.YahooNews, now time.Time) error {

	bulks := make([]dao.News, len(rows))
	for _, r := range rows {
		bulks = append(
			bulks,
			dao.News{
				Title:       r.GetTitle(),
				Description: r.GetDescription(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		)
	}
	news := dao.News{}
	err := news.BulkInsert(context.Background(), db, bulks, now)
	if err != nil {
		return err
	}

	return nil
}
