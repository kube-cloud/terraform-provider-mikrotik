package mikrotik

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * IPSec Profile Resource Create Test
 */
func TestIpSecProfile_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_profile.testacc"

	// Define IPSec Profile Values
	name := "ipsec-profile"
	dhGroup := "modp2048"
	dpdInterval := "2m"
	dpdMaxFailure := 5
	encAlgorithms := "aes-256,aes-192,aes-128"
	hashAlgorithm := "sha1"
	lifetime := "1h30m"
	natTraversal := true
	proposalCheck := "obey"
	updatedName := name + "_updated"
	updatedEncAlgorithms := "aes-256"
	updatedHashAlgorithm := "sha256"
	updatedNatTransversal := false

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecProfileDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecProfile(
					name,
					dhGroup,
					dpdInterval,
					dpdMaxFailure,
					encAlgorithms,
					hashAlgorithm,
					lifetime,
					natTraversal,
					proposalCheck,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecProfile(
					updatedName,
					dhGroup,
					dpdInterval,
					dpdMaxFailure,
					updatedEncAlgorithms,
					updatedHashAlgorithm,
					lifetime,
					updatedNatTransversal,
					proposalCheck,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecProfileExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecProfile(rs.Primary.ID)

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
func testAccCheckIpSecProfileDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not IPSec Profile
		if rs.Type != "mikrotik_ipsec_profile" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindIpSecProfile(rs.Primary.ID)

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
func testAccIpSecProfile(
	name string,
	dhGroup string,
	dpdInterval string,
	dpdMaxFailure int,
	encAlgorithms string,
	hashAlgorithm string,
	lifetime string,
	natTraversal bool,
	proposalCheck string,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_ipsec_profile" "testacc" {
			name = %q
			dh_group = %q
			dpd_interval = %q
			dpd_max_failure = %d
			enc_algorithms = %q
			hash_algorithm = %q
			lifetime = %q
			nat_traversal = %t
			proposal_check = %q
		}
	`, name, dhGroup, dpdInterval, dpdMaxFailure, encAlgorithms, hashAlgorithm, lifetime, natTraversal, proposalCheck)
}
