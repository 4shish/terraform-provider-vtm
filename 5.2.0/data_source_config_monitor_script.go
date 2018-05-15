// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	vtm "github.com/pulse-vadc/go-vtm/5.2"
)

func dataSourceMonitorScript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMonitorScriptRead,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Object text
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceMonitorScriptRead(d *schema.ResourceData, tm interface{}) error {
	objectName := d.Get("name").(string)
	object, err := tm.(*vtm.VirtualTrafficManager).GetMonitorScript(objectName)
	if err != nil {
		if err.ErrorId == "resource.not_found" {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Failed to read vtm_monitor_script '%v': %v", objectName, err.ErrorText)
	}
	d.Set("content", object)
	d.SetId(objectName)
	return nil
}