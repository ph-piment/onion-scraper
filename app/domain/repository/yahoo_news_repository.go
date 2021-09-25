package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ph-piment/onion-scraper/app/domain/entity"
)

type YahooNewsRepository interface {
	ScrapingListFromWEB(ctx context.Context) ([]*entity.YahooNews, error)
	ImportToDB(ctx context.Context, db *sqlx.DB, list []*entity.YahooNews, now time.Time) error
}
