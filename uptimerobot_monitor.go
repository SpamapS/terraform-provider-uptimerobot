package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"spamaps.org/uptimerobot"
	"strconv"
)

func uptimerobotMonitorCreate(d *schema.ResourceData, m interface{}) error {
	mon := uptimerobot.Monitor{
		Friendly_name: fmt.Sprintf("%s", d.Get("friendly_name")),
		Url:           fmt.Sprintf("%s", d.Get("url")),
		Monitor_type:  d.Get("type").(int),
	}
	err := m.(*uptimerobot.Client).CreateMonitor(&mon)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%d", mon.Id))
	return nil
}

func uptimerobotMonitorRead(d *schema.ResourceData, m interface{}) error {
	i, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	mons, err := m.(*uptimerobot.Client).GetMonitors([]int{i})
	if err != nil {
		return err
	}
	d.Set("friendly_name", mons[0].Friendly_name)
	d.Set("url", mons[0].Url)
	d.Set("type", fmt.Sprintf("%d", mons[0].Monitor_type))
	return nil
}

func uptimerobotMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	i, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	mon := uptimerobot.Monitor{
		Id:            i,
		Friendly_name: fmt.Sprintf("%s", d.Get("friendly_name")),
		Url:           fmt.Sprintf("%s", d.Get("url")),
		Monitor_type:  d.Get("type").(int),
	}
	return m.(*uptimerobot.Client).EditMonitor(&mon)
}

func uptimerobotMonitorDelete(d *schema.ResourceData, m interface{}) error {
	i, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	return m.(*uptimerobot.Client).DeleteMonitor(i)
}

func uptimerobotMonitor() *schema.Resource {
	return &schema.Resource{
		Create: uptimerobotMonitorCreate,
		Read:   uptimerobotMonitorRead,
		Update: uptimerobotMonitorUpdate,
		Delete: uptimerobotMonitorDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
