package mikrotik

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Firewall Nat Resource Create Test
 */
func TestFirewallNat_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_firewall_nat.testacc"

	// Define Firewall Nat Values
	chain := "srcnat"
	disabled := false
	updatedDisabled := true

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckFirewallNatDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccFirewallNat(
					chain,
					disabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccFirewallNatExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "chain", chain),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
				),
			},
			{
				// Configure Test Resource
				Config: testAccFirewallNat(
					chain,
					updatedDisabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccFirewallNatExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "chain", chain),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(updatedDisabled)),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccFirewallNatExists(resourceName string) resource.TestCheckFunc {

	// Find and return result
	return func(s *terraform.State) error {

		// Find resource
		rs, ok := s.RootModule().Resources[resourceName]

		// If not OK
		if !ok {

			// Return Not Found Error Message
			return fmt.Errorf("Not found: %s", resourceName)
		}

		// If Resource ID is Empty
		if rs.Primary.ID == "" {

			// Return ENot Exists Error Message
			return fmt.Errorf("%s does not exist in the statefile", resourceName)
		}

		// Build Client
		c := client.NewClient(client.GetConfigFromEnv())

		// Find Resource by ID
		record, err := c.FindFirewallNat(rs.Primary.ID)

		// If there are Error
		if err != nil {

			// Return Formatted Error Message
			return fmt.Errorf("Unable to get remote record for %s: %v", resourceName, err)
		}

		// If No resource Found
		if record == nil {

			// Return Formatted Message
			return fmt.Errorf("Unable to get the remote record %s", resourceName)
		}

		// Return Null
		return nil
	}
}

/**
 * Function used to Test if Terraform Resource is Destroyed
 */
func testAccCheckFirewallNatDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not Firewall Nat
		if rs.Type != "mikrotik_firewall_nat" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindFirewallNat(rs.Primary.ID)

		// Process NotFound
		_, ok := err.(*client.NotFound)

		// If Not OK and Error
		if !ok && err != nil {

			// Return Error
			return err
		}

		// If Record Exists
		if remoteRecord != nil {

			// Return Formatted Error
			return fmt.Errorf("remote record (%s) still exists", remoteRecord.Id)
		}
	}

	// Return Nil
	return nil
}

/**
 * Function used to Print Testing Resource
 */
func testAccFirewallNat(
	chain string,
	disabled bool,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_firewall_nat" "testacc" {
			chain = %q
			disabled = %t
		}
	`, chain, disabled)
}
