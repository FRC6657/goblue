package goblue

type Status struct {
	CurrentSeason  int      `json:"current_season"`
	MaxSeason      int      `json:"max_season"`
	IsDatafeedDown bool     `json:"is_datafeed_down"`
	DownEvents     []string `json:"down_events"`
	Ios            struct {
		MinAppVersion    int `json:"min_app_version"`
		LatestAppVersion int `json:"latest_app_version"`
	} `json:"ios"`
	Android struct {
		MinAppVersion    int `json:"min_app_version"`
		LatestAppVersion int `json:"latest_app_version"`
	} `json:"android"`
}

// Status : GET /status
func (c *Client) Status() (s Status) {
	if err := c.make("/status", &s); err != nil {
		return Status{}
	}
	return
}
