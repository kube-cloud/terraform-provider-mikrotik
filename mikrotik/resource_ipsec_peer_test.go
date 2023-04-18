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
 * IPSec Peer Resource Create Test
 */
func TestIpSecPeer_Create(t *testing.T) {

	profileName := "ipsec-profile"
	profileDhGroup := "modp2048"
	profileDpdInterval := "2m"
	profileDpdMaxFailure := 5
	profileEncAlgorithms := "aes-256,aes-192,aes-128"
	profileHashAlgorithm := "sha1"
	profileLifetime := "1h30m"
	profileNatTraversal := true
	profileProposalCheck := "obey"

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_peer.testacc"

	// Define IPSec Peer Values
	name := "ipsec-peer"
	address := "192.16.2.14/32"
	exchangeMode := "ike2"
	sendInitialContact := false
	passive := true
	localAddress := "192.16.3.19"
	port := 0
	updatedName := name + "_updated"
	updatedSendInitialContact := true
	updatedPassive := false

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecPeerDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecPeer(
					profileName,
					profileDhGroup,
					profileDpdInterval,
					profileDpdMaxFailure,
					profileEncAlgorithms,
					profileHashAlgorithm,
					profileLifetime,
					profileNatTraversal,
					profileProposalCheck,
					name,
					address,
					exchangeMode,
					sendInitialContact,
					passive,
					localAddress,
					port,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPeerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "exchange_mode", exchangeMode),
					resource.TestCheckResourceAttr(resourceName, "send_initial_contact", strconv.FormatBool(sendInitialContact)),
					resource.TestCheckResourceAttr(resourceName, "passive", strconv.FormatBool(passive)),
					resource.TestCheckResourceAttr(resourceName, "local_address", localAddress),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecPeer(
					profileName,
					profileDhGroup,
					profileDpdInterval,
					profileDpdMaxFailure,
					profileEncAlgorithms,
					profileHashAlgorithm,
					profileLifetime,
					profileNatTraversal,
					profileProposalCheck,
					updatedName,
					address,
					exchangeMode,
					updatedSendInitialContact,
					updatedPassive,
					localAddress,
					port,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPeerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "exchange_mode", exchangeMode),
					resource.TestCheckResourceAttr(resourceName, "send_initial_contact", strconv.FormatBool(updatedSendInitialContact)),
					resource.TestCheckResourceAttr(resourceName, "passive", strconv.FormatBool(updatedPassive)),
					resource.TestCheckResourceAttr(resourceName, "local_address", localAddress),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecPeerExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecPeer(rs.Primary.ID)

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
func testAccCheckIpSecPeerDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not IPSec Peer
		if rs.Type != "mikrotik_ipsec_peer" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindIpSecPeer(rs.Primary.ID)

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
func testAccIpSecPeer(
	profileName string,
	profileDhGroup string,
	profileDpdInterval string,
	profileDpdMaxFailure int,
	profileEncAlgorithms string,
	profileHashAlgorithm string,
	profileLifetime string,
	profileNatTraversal bool,
	profileProposalCheck string,
	name string,
	address string,
	exchangeMode string,
	sendInitialContact bool,
	passive bool,
	localAddress string,
	port int,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_ipsec_profile" "testaccprofile" {
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
		resource "mikrotik_ipsec_peer" "testacc" {
			name = %q
			address = %q
			profile = mikrotik_ipsec_profile.testaccprofile.name
			exchange_mode = %q
			send_initial_contact = %t
			passive = %t
			local_address = %q
			port = %d
		}
	`,
		profileName,
		profileDhGroup,
		profileDpdInterval,
		profileDpdMaxFailure,
		profileEncAlgorithms,
		profileHashAlgorithm,
		profileLifetime,
		profileNatTraversal,
		profileProposalCheck,
		name,
		address,
		exchangeMode,
		sendInitialContact,
		passive,
		localAddress,
		port,
	)
}
