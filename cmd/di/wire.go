//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	repository_impl "github.com/ph-piment/onion-scraper/app/infrastructure/repository_impl"
	usecase "github.com/ph-piment/onion-scraper/app/usecase/yahoo"
)

func InitializeEvent() usecase.YahooNews {
	wire.Build(
		usecase.NewYahooNews,
		repository_impl.NewYahooNewsRepository,
	)
	return usecase.NewYahooNews(nil)
}
