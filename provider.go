package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/url"
	"spamaps.org/uptimerobot"
)

func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_monitor": uptimerobotMonitor(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		u, err := url.Parse("https://api.uptimerobot.com/v2")
		if err != nil {
			return nil, err
		}
		c := uptimerobot.Client{
			BaseURL:    u,
			UserAgent:  "terraform-provider-uptimerobot",
			HttpClient: nil,
			Api_key:    d.Get("api_key").(string),
		}
		return c, nil
	}
}
