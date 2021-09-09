package usecase

import (
	"context"
	"time"

	"github.com/ph-piment/onion-scraper/app/domain/repository"
)

type YahooNews interface {
	Import(ctx context.Context, now time.Time) error
}

type yahooNews struct {
	repo repository.YahooNewsRepository
}

func NewYahooNews(
	repo repository.YahooNewsRepository,
) YahooNews {
	return &yahooNews{
		repo: repo,
	}
}

func (uc *yahooNews) Import(ctx context.Context, now time.Time) error {
	rows, err := uc.repo.ScrapingListFromWEB(ctx, now)
	if err != nil {
		return err
	}
	err = uc.repo.ImportToDB(ctx, rows, now)
	if err != nil {
		return err
	}
	return nil
}
