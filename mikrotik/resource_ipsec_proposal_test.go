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
 * IPSec Proposal Resource Create Test
 */
func TestIpSecProposal_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_proposal.testacc"

	// Define IPSec Proposal Values
	name := "TestName"
	authAlgorithms := "sha512,sha256,sha1"
	encAlgorithms := "aes-256-cbc,aes-192-cbc,aes-128-cbc"
	lifetime := "30m"
	pfsGroup := "modp2048"
	disabled := false
	updatedName := name + "_updated"
	updatedLifetime := "45m"
	updatedEncAlgorithms := "aes-256-cbc"

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecProposalDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecProposal(
					name,
					authAlgorithms,
					encAlgorithms,
					lifetime,
					pfsGroup,
					disabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecProposalExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "auth_algorithms", authAlgorithms),
					resource.TestCheckResourceAttr(resourceName, "enc_algorithms", encAlgorithms),
					resource.TestCheckResourceAttr(resourceName, "lifetime", lifetime),
					resource.TestCheckResourceAttr(resourceName, "pfsGroup", pfsGroup),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecProposal(
					updatedName,
					authAlgorithms,
					updatedEncAlgorithms,
					updatedLifetime,
					pfsGroup,
					disabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecProposalExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "auth_algorithms", authAlgorithms),
					resource.TestCheckResourceAttr(resourceName, "enc_algorithms", updatedEncAlgorithms),
					resource.TestCheckResourceAttr(resourceName, "lifetime", updatedLifetime),
					resource.TestCheckResourceAttr(resourceName, "pfsGroup", pfsGroup),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecProposalExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecProposal(rs.Primary.ID)

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

func testAccCheckIpSecProposalDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_ipsec_proposal" {
			continue
		}

		remoteRecord, err := c.FindIpSecProposal(rs.Primary.ID)

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

/**
 * Function used to Print Testing Resource
 */
func testAccIpSecProposal(
	name string,
	authAlgorithms string,
	encAlgorithms string,
	lifetime string,
	pfsGroup string,
	disabled bool,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_ipsec_proposal" "testacc" {
			name = %q
			auth_algorithms = %q
			enc_algorithms = %q
			lifetime = %q
			pfs_group = %q
			disabled = %t
		}
	`, name, authAlgorithms, encAlgorithms, lifetime, pfsGroup, disabled)
}
