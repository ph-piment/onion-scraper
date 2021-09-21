package repository_impl

import (
	"context"
	"testing"
	"time"

	"github.com/ph-piment/onion-scraper/app/domain/entity"
)

func Test_NewYahooNewsRepository(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "NewYahooNewsRepository",
		},
	}
	for _, r := range tests {
		t.Run(r.name, func(t *testing.T) {
			got := NewYahooNewsRepository()
			if got == nil {
				t.Errorf("NewYahooNewsRepository() = nil")
			}
		})
	}
}

func Test_ScrapingListFromWEB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ScrapingListFromWEB",
		},
	}
	for _, r := range tests {
		t.Run(r.name, func(t *testing.T) {
			got := NewYahooNewsRepository()
			if got == nil {
				t.Errorf("NewYahooNewsRepository() = nil")
			}
			rows, err := got.ScrapingListFromWEB(context.Background())
			if err != nil {
				t.Errorf("ScrapingListFromWEB error = %v", err)
			}
			if len(rows) == 0 {
				t.Error("ScrapingListFromWEB count = 0")
			}
		})
	}
}

func Test_ImportToDB(t *testing.T) {
	type fields struct {
		entities []*entity.YahooNews
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "ImportToDB",
			fields: fields{
				entities: []*entity.YahooNews{
					entity.NewYahooNews(1, "aaa", "bbb"),
					entity.NewYahooNews(2, "ccc", "ddd"),
					entity.NewYahooNews(3, "eee", "fff"),
				},
			},
		},
	}
	for _, r := range tests {
		t.Run(r.name, func(t *testing.T) {
			got := NewYahooNewsRepository()
			if got == nil {
				t.Errorf("NewYahooNewsRepository() = nil")
			}
			err := got.ImportToDB(context.Background(), r.fields.entities, time.Now())
			if err != nil {
				t.Errorf("ImportToDB error = %v", err)
			}
		})
	}
}
