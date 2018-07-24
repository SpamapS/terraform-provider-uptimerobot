package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"net/url"
	"spamaps.org/uptimerobot"
)

func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTIMEROBOT_API_KEY", nil),
			},
		},
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
			HttpClient: &http.Client{},
			Api_key:    d.Get("api_key").(string),
		}
		return &c, nil
	}
}
