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

	resourceName := "mikrotik_bridge_interface_port.testacc"
	iface := "ether1"
	bridge := "test-brigde"
	horizon := "none"
	learn := "auto"
	unknown_multicast_flood := true
	unknown_unicast_flood := false
	broadcast_flood := false
	trusted := true
	hardware_offload := false
	auto_isolate := false
	restricted_role := false
	restricted_tcn := false
	bpdu_guard := false
	priority := 90
	path_cost := 20
	internal_path_cost := 30
	edge := "no-discover"
	point_to_point := "yes"
	disabled := false
	comment := "Test Bridge Port Comment"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckBridgeInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgeInterfacePort(iface, bridge, horizon, learn, unknown_multicast_flood,
					unknown_unicast_flood, broadcast_flood, trusted, hardware_offload,
					auto_isolate, restricted_role, restricted_tcn, bpdu_guard, priority,
					path_cost, internal_path_cost, edge, point_to_point, disabled, comment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeInterfaceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "iface", iface),
					resource.TestCheckResourceAttr(resourceName, "bridge", bridge),
					resource.TestCheckResourceAttr(resourceName, "horizon", horizon),
					resource.TestCheckResourceAttr(resourceName, "learn", learn),
					resource.TestCheckResourceAttr(resourceName, "edge", edge),
					resource.TestCheckResourceAttr(resourceName, "point_to_point", point_to_point),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
					resource.TestCheckResourceAttr(resourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(resourceName, "path_cost", strconv.Itoa(path_cost)),
					resource.TestCheckResourceAttr(resourceName, "internal_path_cost", strconv.Itoa(internal_path_cost)),
				),
			},
			{
				Config: testAccBridgeInterfacePort(iface, bridge, horizon, learn, unknown_multicast_flood,
					unknown_unicast_flood, broadcast_flood, trusted, hardware_offload,
					auto_isolate, restricted_role, restricted_tcn, bpdu_guard, priority,
					path_cost+5, internal_path_cost+10, edge, point_to_point, disabled, comment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeInterfaceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "iface", iface),
					resource.TestCheckResourceAttr(resourceName, "bridge", bridge),
					resource.TestCheckResourceAttr(resourceName, "horizon", horizon),
					resource.TestCheckResourceAttr(resourceName, "learn", learn),
					resource.TestCheckResourceAttr(resourceName, "edge", edge),
					resource.TestCheckResourceAttr(resourceName, "point_to_point", point_to_point),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
					resource.TestCheckResourceAttr(resourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(resourceName, "path_cost", strconv.Itoa(path_cost)+5),
					resource.TestCheckResourceAttr(resourceName, "internal_path_cost", strconv.Itoa(internal_path_cost)+10),
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
		return nil
	}
}

func testAccCheckBridgeInterfaceDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_bridge_interface_port" {
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

func testAccBridgeInterfacePort(iface string, bridge string, horizon string, learn string,
	unknown_multicast_flood bool, unknown_unicast_flood bool, broadcast_flood bool, trusted bool,
	hardware_offload bool, auto_isolate bool, restricted_role bool, restricted_tcn bool, bpdu_guard bool,
	priority int, path_cost int, internal_path_cost int, edge string, point_to_point string, disabled bool, comment string) string {
	return fmt.Sprintf(`
		resource "mikrotik_bridge_interface_port" "testacc" {
			interface = %q
			bridge = %q
			horizon = %q
			learn = %q
			unknown_multicast_flood = %t
			unknown_unicast_flood = %t
			broadcast_flood = %t
			trusted = %t
			hardware_offload = %t
			auto_isolate = %t
			restricted_role = %t
			restricted_tcn = %t
			bpdu_guard = %t
			priority = %d
			path_cost = %d
			internal_path_cost = %d
			edge = %q
			point_to_point = %q
			disabled = %t
			comment = %q
		}
	`, iface, bridge, horizon, learn, unknown_multicast_flood,
		unknown_unicast_flood, broadcast_flood, trusted, hardware_offload,
		auto_isolate, restricted_role, restricted_tcn, bpdu_guard, priority,
		path_cost, internal_path_cost, edge, point_to_point, disabled, comment)
}
