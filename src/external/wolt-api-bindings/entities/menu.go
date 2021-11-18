package entities

type Menu struct {
	Categories []struct {
		Description   string      `json:"description"`
		ID            string      `json:"id"`
		Image         interface{} `json:"image"`
		ImageBlurhash interface{} `json:"image_blurhash"`
		ItemCount     int         `json:"item_count"`
		Layout        string      `json:"layout"`
		Name          string      `json:"name"`
	} `json:"categories"`
	ExtraInfoPopup interface{} `json:"extra_info_popup"`
	ID             string      `json:"id"`
	Items          []struct {
		AlcoholPercentage        int           `json:"alcohol_percentage"`
		AllowedDeliveryMethods   []string      `json:"allowed_delivery_methods"`
		Baseprice                int           `json:"baseprice"`
		Category                 string        `json:"category"`
		Checksum                 string        `json:"checksum"`
		Description              string        `json:"description"`
		DisabledInfo             interface{}   `json:"disabled_info"`
		Enabled                  bool          `json:"enabled"`
		HasExtraInfo             bool          `json:"has_extra_info"`
		ID                       string        `json:"id"`
		Image                    string        `json:"image"`
		ImageBlurhash            string        `json:"image_blurhash"`
		IsVenueTip               bool          `json:"is_venue_tip"`
		MaxQuantityPerPurchase   interface{}   `json:"max_quantity_per_purchase"`
		Name                     string        `json:"name"`
		NoContactDeliveryAllowed bool          `json:"no_contact_delivery_allowed"`
		Options                  []interface{} `json:"options"`
		OriginalPrice            interface{}   `json:"original_price"`
		QuantityLeft             interface{}   `json:"quantity_left"`
		QuantityLeftVisible      bool          `json:"quantity_left_visible"`
		QuantityTotal            interface{}   `json:"quantity_total"`
		SellByWeightConfig       interface{}   `json:"sell_by_weight_config"`
		Tags                     []interface{} `json:"tags"`
		Times                    []struct {
			AvailableDaysOfWeek []int `json:"available_days_of_week"`
			AvailableEndDate    int   `json:"available_end_date"`
			AvailableStartDate  int   `json:"available_start_date"`
			VisibleDaysOfWeek   []int `json:"visible_days_of_week"`
			VisibleEndDate      int   `json:"visible_end_date"`
			VisibleStartDate    int   `json:"visible_start_date"`
		} `json:"times"`
		Type          string      `json:"type"`
		UnitInfo      interface{} `json:"unit_info"`
		UnitPrice     interface{} `json:"unit_price"`
		Validity      interface{} `json:"validity"`
		VatPercentage int         `json:"vat_percentage"`
	} `json:"items"`
	Language  string `json:"language"`
	Languages []struct {
		Autotranslated bool   `json:"autotranslated"`
		Language       string `json:"language"`
		Name           string `json:"name"`
		Primary        bool   `json:"primary"`
	} `json:"languages"`
	Name    string `json:"name"`
	Options []struct {
		Defaultvalue string `json:"defaultvalue"`
		HasExtraInfo bool   `json:"has_extra_info"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Values       []struct {
			HideWhenSelected  bool   `json:"hide_when_selected"`
			ID                string `json:"id"`
			MaximumSelections int    `json:"maximum_selections"`
			MinimumSelections int    `json:"minimum_selections"`
			MultichoiceWeight int    `json:"multichoice_weight"`
			Name              string `json:"name"`
			Price             int    `json:"price"`
		} `json:"values"`
	} `json:"options"`
	Recommendations []struct {
		Item    string        `json:"item"`
		Options []interface{} `json:"options"`
	} `json:"recommendations"`
	RecommendedItems []struct {
		AlcoholPercentage        int           `json:"alcohol_percentage"`
		AllowedDeliveryMethods   []string      `json:"allowed_delivery_methods"`
		Baseprice                int           `json:"baseprice"`
		Category                 string        `json:"category"`
		Checksum                 string        `json:"checksum"`
		Description              string        `json:"description"`
		DisabledInfo             interface{}   `json:"disabled_info"`
		Enabled                  bool          `json:"enabled"`
		HasExtraInfo             bool          `json:"has_extra_info"`
		ID                       string        `json:"id"`
		Image                    interface{}   `json:"image"`
		ImageBlurhash            interface{}   `json:"image_blurhash"`
		IsVenueTip               bool          `json:"is_venue_tip"`
		MaxQuantityPerPurchase   interface{}   `json:"max_quantity_per_purchase"`
		Name                     string        `json:"name"`
		NoContactDeliveryAllowed bool          `json:"no_contact_delivery_allowed"`
		Options                  []interface{} `json:"options"`
		OriginalPrice            interface{}   `json:"original_price"`
		QuantityLeft             interface{}   `json:"quantity_left"`
		QuantityLeftVisible      bool          `json:"quantity_left_visible"`
		QuantityTotal            interface{}   `json:"quantity_total"`
		SellByWeightConfig       interface{}   `json:"sell_by_weight_config"`
		Tags                     []interface{} `json:"tags"`
		Times                    []struct {
			AvailableDaysOfWeek []int `json:"available_days_of_week"`
			AvailableEndDate    int   `json:"available_end_date"`
			AvailableStartDate  int   `json:"available_start_date"`
			VisibleDaysOfWeek   []int `json:"visible_days_of_week"`
			VisibleEndDate      int   `json:"visible_end_date"`
			VisibleStartDate    int   `json:"visible_start_date"`
		} `json:"times"`
		Type          string      `json:"type"`
		UnitInfo      interface{} `json:"unit_info"`
		UnitPrice     interface{} `json:"unit_price"`
		Validity      interface{} `json:"validity"`
		VatPercentage int         `json:"vat_percentage"`
	} `json:"recommended_items"`
}
