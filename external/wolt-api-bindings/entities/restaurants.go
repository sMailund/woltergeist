package entities

type Restaurants struct {
	Created struct {
		Date int64 `json:"$date"`
	} `json:"created"`
	ExpiresInSeconds int `json:"expires_in_seconds"`
	Filtering        struct {
		Filters []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Values []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"values"`
		} `json:"filters"`
	} `json:"filtering"`
	Name      string `json:"name"`
	PageTitle string `json:"page_title"`
	Sections  []struct {
		Items []struct {
			Filtering struct {
				Filters []struct {
					ID     string   `json:"id"`
					Values []string `json:"values"`
				} `json:"filters"`
			} `json:"filtering"`
			Image struct {
				Blurhash string   `json:"blurhash"`
				URL      string   `json:"url"`
				Variants []string `json:"variants"`
			} `json:"image"`
			Link struct {
				Target      string `json:"target"`
				TargetSort  string `json:"target_sort"`
				TargetTitle string `json:"target_title"`
				Title       string `json:"title"`
				Type        string `json:"type"`
			} `json:"link"`
			Sorting struct {
				Sortables []struct {
					ID    string `json:"id"`
					Value int    `json:"value"`
				} `json:"sortables"`
			} `json:"sorting"`
			Template string `json:"template"`
			Title    string `json:"title"`
			TrackID  string `json:"track_id"`
			Venue    struct {
				Address                string        `json:"address"`
				Badges                 []interface{} `json:"badges"`
				Categories             []interface{} `json:"categories"`
				City                   string        `json:"city"`
				Currency               string        `json:"currency"`
				Delivers               bool          `json:"delivers"`
				DeliveryPrice          string        `json:"delivery_price"`
				DeliveryPriceHighlight bool          `json:"delivery_price_highlight"`
				Estimate               int           `json:"estimate"`
				EstimateRange          string        `json:"estimate_range"`
				Franchise              string        `json:"franchise"`
				ID                     string        `json:"id"`
				Location               []float64     `json:"location"`
				Name                   string        `json:"name"`
				Online                 bool          `json:"online"`
				PriceRange             int           `json:"price_range"`
				ProductLine            string        `json:"product_line"`
				Rating                 struct {
					Rating int     `json:"rating"`
					Score  float64 `json:"score"`
				} `json:"rating"`
				ShortDescription string   `json:"short_description"`
				ShowWoltPlus     bool     `json:"show_wolt_plus"`
				Slug             string   `json:"slug"`
				Tags             []string `json:"tags"`
			} `json:"venue"`
			Overlay string `json:"overlay,omitempty"`
		} `json:"items"`
		Link struct {
			Target      string `json:"target"`
			TargetSort  string `json:"target_sort"`
			TargetTitle string `json:"target_title"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"link"`
		Name     string `json:"name"`
		Template string `json:"template"`
	} `json:"sections"`
	ShowLargeTitle bool `json:"show_large_title"`
	ShowMap        bool `json:"show_map"`
	Sorting        struct {
		Sortables []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"sortables"`
	} `json:"sorting"`
	TrackID string `json:"track_id"`
}
