package mikrotik

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * IPSec Policy Group Resource Create Test
 */
func TestIpSecPolicyGroup_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_policy_group.testacc"

	// Define IPSec Policy Group Values
	name := "ipsec-pocily-group"
	updatedName := name + "_updated"

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecPolicyGroupDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecPolicyGroup(
					name,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPolicyGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecPolicyGroup(
					updatedName,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPolicyGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecPolicyGroupExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecPolicyGroup(rs.Primary.ID)

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
func testAccCheckIpSecPolicyGroupDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not IPSec Policy Group
		if rs.Type != "mikrotik_ipsec_policy_group" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindIpSecPolicyGroup(rs.Primary.ID)

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
func testAccIpSecPolicyGroup(
	name string,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_ipsec_policy_group" "testacc" {
			name = %q
		}
	`, name)
}
