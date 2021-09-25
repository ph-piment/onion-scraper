package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
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
	rows, err := uc.repo.ScrapingListFromWEB(ctx)
	if err != nil {
		return err
	}

	db, err := sqlx.Open("postgres", "user=root dbname=os password=root sslmode=disable")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	err = uc.repo.ImportToDB(ctx, tx, rows, now)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
