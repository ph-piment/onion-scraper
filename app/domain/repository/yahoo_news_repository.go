package repository

import (
	"context"
	"time"

	"github.com/ph-piment/onion-scraper/app/domain/entity"
)

type YahooNewsRepository interface {
	ScrapingListFromWEB(ctx context.Context) ([]*entity.YahooNews, error)
	ImportToDB(ctx context.Context, db interface{}, list []*entity.YahooNews, now time.Time) error
}
