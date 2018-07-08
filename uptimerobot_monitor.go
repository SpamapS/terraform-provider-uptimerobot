package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func uptimerobotMonitorCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func uptimerobotMonitorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func uptimerobotMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func uptimerobotMonitorDelete(d *schema.ResourceData, m interface{}) error {
	return nil
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
