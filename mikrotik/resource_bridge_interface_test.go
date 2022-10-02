package mikrotik

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func TestBridgeInterface_basic(t *testing.T) {

	resourceName := "mikrotik_bridge_interface.testacc"
	mtu := 1500
	name := "test-brigde"
	comment := "test-comment"
	autoMac := false
	adminMac := "74:4D:28:F3:A7:15"
	disabled := false
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckBridgeInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgeInterface(mtu, name, disabled, autoMac, adminMac, comment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeInterfaceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "mtu", strconv.Itoa(mtu)),
				),
			},
			{
				Config: testAccBridgeInterface(mtu, name+"updated", disabled, autoMac, adminMac, comment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeInterfaceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name+"updated"),
					resource.TestCheckResourceAttr(resourceName, "mtu", strconv.Itoa(mtu)),
				),
			},
		},
	})
}

func testAccBridgeInterfaceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s does not exist in the statefile", resourceName)
		}

		c := client.NewClient(client.GetConfigFromEnv())
		record, err := c.FindBridgeInterface(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Unable to get remote record for %s: %v", resourceName, err)
		}

		if record == nil {
			return fmt.Errorf("Unable to get the remote record %s", resourceName)
		}

		fmt.Println("=========================")
		fmt.Println("=========================")
		fmt.Println("=====> LOADED RESOURCES :")
		fmt.Println(record)
		fmt.Println("=========================")
		fmt.Println("=========================")

		return nil
	}
}

func testAccCheckBridgeInterfaceDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_bridge_interface" {
			continue
		}

		remoteRecord, err := c.FindBridgeInterface(rs.Primary.ID)

		_, ok := err.(*client.NotFound)
		if !ok && err != nil {
			return err
		}

		if remoteRecord != nil {
			return fmt.Errorf("remote record (%s) still exists", remoteRecord.Id)
		}

	}
	return nil
}

func testAccBridgeInterface(mtu int, name string, disabled bool, autoMac bool, adminMac string, comment string) string {
	return fmt.Sprintf(`
		resource "mikrotik_bridge_interface" "testacc" {
			mtu = %d
			name = %q
			disabled = %t
			auto_mac = %t
			admin_mac = %q
			comment = %q
		}
	`, mtu, name, disabled, autoMac, adminMac, comment)
}
