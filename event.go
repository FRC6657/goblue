package goblue

type EventSimple struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	EventCode string `json:"event_code"`
	EventType int    `json:"event_type"`
	District  `json:"district"`
	City      string `json:"city"`
	StateProv string `json:"state_prov"`
	Country   string `json:"country"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Year      int    `json:"year"`
}

type Event struct {
	EventSimple
	ShortName       string  `json:"short_name"`
	EventTypeString string  `json:"event_type_string"`
	Week            int     `json:"week"`
	Address         string  `json:"address"`
	PostalCode      string  `json:"postal_code"`
	GmapsPlaceID    string  `json:"gmaps_place_id"`
	GmapsURL        string  `json:"gmaps_url"`
	Lat             float32 `json:"lat"`
	Lng             float32 `json:"lng"`
	LocationName    string  `json:"location_name"`
	Timezone        string  `json:"timezone"`
	Website         string  `json:"website"`
	FirstEventID    string  `json:"first_event_id"`
	FirstEventCode  string  `json:"first_event_code"`
	Webcasts        []struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		File    string `json:"file"`
	} `json:"webcasts"`
	DivisionKeys      []string `json:"division_keys"`
	ParentEventKey    string   `json:"parent_event_key"`
	PlayoffType       int      `json:"playoff_type"`
	PlayoffTypeString string   `json:"playoff_type_string"`
}
