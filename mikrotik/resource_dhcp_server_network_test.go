package mikrotik

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func TestDhcpServerNetwork_basic(t *testing.T) {

	resourceName := "mikrotik_dhcp_server_network.testacc"

	netmask := "24"
	address := "10.10.10.0/" + netmask
	gateway := "10.10.10.2"
	dnsServer := "10.10.10.3"
	comment := "Terraform managed"
	dnsServerUpdated := "192.168.5.3"
	nextServer := "10.10.10.3"
	ntpServer := "10.10.10.3"
	winsServer := "10.10.10.3"
	bootFileName := "pxelinux.0"
	domain := "kue-cloud.com"
	dhcpOptionSet := ""
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDhcpServerNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpServerNetwork(
					address,
					netmask,
					gateway,
					dnsServer,
					comment,
					nextServer,
					ntpServer,
					winsServer,
					bootFileName,
					domain,
					dhcpOptionSet,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDhcpServerNetworkExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
					resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
					resource.TestCheckResourceAttr(resourceName, "dns_server", dnsServer),
					resource.TestCheckResourceAttr(resourceName, "next_server", nextServer),
					resource.TestCheckResourceAttr(resourceName, "ntp_server", ntpServer),
					resource.TestCheckResourceAttr(resourceName, "wins_server", winsServer),
					resource.TestCheckResourceAttr(resourceName, "boot_file_name", bootFileName),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "dhcp_option_set", dhcpOptionSet),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			{
				Config: testAccDhcpServerNetwork(
					address,
					netmask,
					gateway,
					dnsServerUpdated,
					comment,
					nextServer,
					ntpServer,
					winsServer,
					bootFileName,
					domain,
					dhcpOptionSet,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDhcpServerNetworkExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "netmask", netmask),
					resource.TestCheckResourceAttr(resourceName, "gateway", gateway),
					resource.TestCheckResourceAttr(resourceName, "dns_server", dnsServerUpdated),
					resource.TestCheckResourceAttr(resourceName, "next_server", nextServer),
					resource.TestCheckResourceAttr(resourceName, "ntp_server", ntpServer),
					resource.TestCheckResourceAttr(resourceName, "wins_server", winsServer),
					resource.TestCheckResourceAttr(resourceName, "boot_file_name", bootFileName),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "dhcp_option_set", dhcpOptionSet),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
		},
	})
}

func testAccDhcpServerNetwork(address, netmask, gateway, dns_server, comment, next_server, ntp_server, wins_server, boot_file_name, domain, dhcp_option_set string) string {
	return fmt.Sprintf(`
resource mikrotik_dhcp_server_network "testacc" {
	address    		= %q
	netmask    		= %q
	gateway    		= %q
	dns_server 		= %q
	next_server 	= %q
	ntp_server 		= %q
	wins_server 	= %q
	boot_file_name	= %q
	domain 			= %q
	dhcp_option_set	= %q
	comment    		= %q
}
`, address, netmask, gateway, dns_server, next_server, ntp_server, wins_server, boot_file_name, domain, dhcp_option_set, comment)
}

func testAccCheckDhcpServerNetworkDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_dhcp_server_network" {
			continue
		}

		remoteRecord, err := c.FindDhcpServerNetwork(rs.Primary.ID)

		_, ok := err.(*client.NotFound)
		if !ok && err != nil {
			return err
		}

		if remoteRecord != nil {
			return fmt.Errorf("remote recrod (%s) still exists", remoteRecord.Id)
		}

	}
	return nil
}

func testAccDhcpServerNetworkExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s does not exist in the statefile", resourceName)
		}

		c := client.NewClient(client.GetConfigFromEnv())
		record, err := c.FindDhcpServerNetwork(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Unable to get remote record for %s: %v", resourceName, err)
		}

		if record == nil {
			return fmt.Errorf("Unable to get the remote record %s", resourceName)
		}

		return nil
	}
}
