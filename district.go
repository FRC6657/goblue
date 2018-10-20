package goblue

import (
	"fmt"
)

type District struct {
	Abbreviation string `json:"abbreviation"`
	DisplayName  string `json:"display_name"`
	Key          string `json:"key"`
	Year         int    `json:"year"`
}

// TeamDistricts : GET /team/{team_key}/districts
func (c *Client) TeamDistricts(teamKey string) (t []District) {
	if err := c.make(fmt.Sprintf("/team/%s/districts", teamKey), &t); err != nil {
		return []District{}
	}
	return
}
