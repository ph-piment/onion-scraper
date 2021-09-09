package repository

import (
	"context"
	"time"

	"github.com/ph-piment/onion-scraper/app/domain/entity"
)

type YahooNewsRepository interface {
	ScrapingListFromWEB(ctx context.Context, now time.Time) ([]*entity.YahooNews, error)
	ImportToDB(ctx context.Context, list []*entity.YahooNews, now time.Time) error
}
