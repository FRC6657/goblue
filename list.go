package goblue

import (
	"fmt"
)

// Teams : GET /teams/{page_num}
func (c *Client) Teams(pageNum int) (t []Team) {
	if err := c.make(fmt.Sprintf("/teams/%d", pageNum), &t); err != nil {
		return []Team{}
	}
	return
}

// TeamsSimple : GET /teams/{page_num}/simple
func (c *Client) TeamsSimple(pageNum int) (t []TeamSimple) {
	if err := c.make(fmt.Sprintf("/teams/%d/simple", pageNum), &t); err != nil {
		return []TeamSimple{}
	}
	return
}

// TeamsKeys : GET /teams/{page_num}/keys
func (c *Client) TeamsKeys(pageNum int) (t []string) {
	if err := c.make(fmt.Sprintf("/teams/%d/keys", pageNum), &t); err != nil {
		return nil
	}
	return
}

// TeamsYear : GET /teams/{year}/{page_num}
func (c *Client) TeamsYear(year, pageNum int) (t []Team) {
	if err := c.make(fmt.Sprintf("/teams/%d/%d", year, pageNum), &t); err != nil {
		return []Team{}
	}
	return
}

// TeamsYearSimple : GET /teams/{year}/{page_num}/simple
func (c *Client) TeamsYearSimple(year, pageNum int) (t []TeamSimple) {
	if err := c.make(fmt.Sprintf("/teams/%d/%d/simple", year, pageNum), &t); err != nil {
		return []TeamSimple{}
	}
	return
}

// TeamsYearKeys : GET /teams/{year}/{page_num}/keys
func (c *Client) TeamsYearKeys(year, pageNum int) (t []string) {
	if err := c.make(fmt.Sprintf("/teams/%d/%d/keys", year, pageNum), &t); err != nil {
		return nil
	}
	return
}

/*
TODO:
GET /team/{team_key}/events/{year}/statuses
*/

// Events : GET /events/{year}
func (c *Client) Events(year int) (t []Event) {
	if err := c.make(fmt.Sprintf("/events/%d", year), &t); err != nil {
		return []Event{}
	}
	return
}

// EventsSimple : GET /events/{year}/simple
func (c *Client) EventsSimple(year int) (t []EventSimple) {
	if err := c.make(fmt.Sprintf("/events/%d/simple", year), &t); err != nil {
		return []EventSimple{}
	}
	return
}
