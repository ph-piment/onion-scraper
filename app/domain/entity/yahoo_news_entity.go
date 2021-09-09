package entity

type YahooNews struct {
	id          uint64
	title       string
	description string
}

func NewYahooNews(
	id uint64,
	title string,
	description string,
) *YahooNews {
	return &YahooNews{
		id:          id,
		title:       title,
		description: description,
	}
}

func (y *YahooNews) GetID() uint64 {
	return y.id
}

func (y *YahooNews) GetTitle() string {
	return y.title
}

func (y *YahooNews) GetDescription() string {
	return y.description
}
