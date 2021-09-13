package repository_impl

import (
	"context"
	"testing"
	"time"
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
	type fields struct {
		id          uint64
		title       string
		description string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "NewYahooNewsRepository",
			fields: fields{
				id:          1,
				title:       "title1",
				description: "description1",
			},
		},
	}
	for _, r := range tests {
		t.Run(r.name, func(t *testing.T) {
			got := NewYahooNewsRepository()
			if got == nil {
				t.Errorf("NewYahooNewsRepository() = nil")
			}
			rows, err := got.ScrapingListFromWEB(context.Background(), time.Now())
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
}
