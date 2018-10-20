package goblue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	APIBase = "http://thebluealliance.com/api/v3"
)

type Client struct {
	Key string
}

func New(key string) *Client {
	c := new(Client)
	c.Key = key
	return c
}

func (c *Client) makeRequest(endpoint string) (data []byte, err error) {
	req, err := http.NewRequest("GET", APIBase+endpoint, nil)
	if err != nil {
		return
	}
	req.Header.Set("X-TBA-Auth-Key", c.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return body, nil
}

func (c *Client) make(endpoint string, v interface{}) (err error) {
	data, err := c.makeRequest(endpoint)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}
