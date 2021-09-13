package entity

import "testing"

func Test_NewYahooNews(t *testing.T) {
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
			name: "NewYahooNews",
			fields: fields{
				id:          1,
				title:       "title1",
				description: "description1",
			},
		},
	}
	for _, r := range tests {
		t.Run(r.name, func(t *testing.T) {
			got := NewYahooNews(
				r.fields.id,
				r.fields.title,
				r.fields.description,
			)
			if got == nil {
				t.Errorf("NewYahooNews() = nil")
			}
			if g := got.GetID(); g != r.fields.id {
				t.Errorf("GetID(): got = %v, want = %v", g, r.fields.id)
			}
			if g := got.GetTitle(); g != r.fields.title {
				t.Errorf("GetTitle(): got = %v, want = %v", g, r.fields.title)
			}
			if g := got.GetDescription(); g != r.fields.description {
				t.Errorf("GetDescription(): got = %v, want = %v", g, r.fields.description)
			}
		})
	}
}
