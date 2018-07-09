package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
	api_key    string
}

type Log struct {
	log_type int
	datetime int
	duration int
}

type Monitor struct {
	id              string
	friendly_name   string
	url             string
	monitor_type    int
	sub_type        string
	keyword_type    string
	keyword_value   string
	http_username   string
	http_password   string
	port            string
	interval        int
	status          int
	create_datetime int
	monitor_group   int
	is_group_main   int
	logs            []Log
}

type Pagination struct {
	offset int
	limit  int
	total  int
}

type MonitorResp struct {
	stat       string
	pagination Pagination
	monitors   []Monitor
}

func (c *Client) getMonitors() ([]Monitor, error) {
	data := url.Values{}
	data.Set("api_key", c.api_key)
	data.Set("format", "json")

	rel := &url.URL{Path: "/getMonitors"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var monitors_resp MonitorResp
	err = json.NewDecoder(resp.Body).Decode(&monitors_resp)
	if err != nil {
		return nil, err
	}
	return monitors_resp.monitors, err
}
