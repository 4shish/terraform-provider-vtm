// Copyright (C) 2018-2019, Pulse Secure, LLC.
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package main

/*
 * This test covers the following cases:
 *   - Creation and deletion of a vtm_servicediscovery object with minimal configuration
 */

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	vtm "github.com/pulse-vadc/go-vtm/5.2"
)

func TestResourceServicediscovery(t *testing.T) {
	objName := acctest.RandomWithPrefix("TestServicediscovery")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServicediscoveryDestroy,
		Steps: []resource.TestStep{
			{
				Config: getBasicServicediscoveryConfig(objName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicediscoveryExists,
				),
			},
		},
	})
}

func testAccCheckServicediscoveryExists(s *terraform.State) error {
	for _, tfResource := range s.RootModule().Resources {
		if tfResource.Type != "vtm_servicediscovery" {
			continue
		}
		objectName := tfResource.Primary.Attributes["name"]
		tm := testAccProvider.Meta().(*vtm.VirtualTrafficManager)
		if _, err := tm.GetServicediscovery(objectName); err != nil {
			return fmt.Errorf("Servicediscovery %s does not exist: %#v", objectName, err)
		}
	}

	return nil
}

func testAccCheckServicediscoveryDestroy(s *terraform.State) error {
	for _, tfResource := range s.RootModule().Resources {
		if tfResource.Type != "vtm_servicediscovery" {
			continue
		}
		objectName := tfResource.Primary.Attributes["name"]
		tm := testAccProvider.Meta().(*vtm.VirtualTrafficManager)
		if _, err := tm.GetServicediscovery(objectName); err == nil {
			return fmt.Errorf("Servicediscovery %s still exists", objectName)
		}
	}

	return nil
}

func getBasicServicediscoveryConfig(name string) string {
	return fmt.Sprintf(`
        resource "vtm_servicediscovery" "test_vtm_servicediscovery" {
			name = "%s"
			content = "TEST_TEXT"

        }`,
		name,
	)
}
