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
 * IPSec Policy Resource Create Test
 */
func TestIpSecPolicy_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_policy.testacc"

	// Define IPSec Peer Values
	peerName := "ipsec-peer"

	// Define IPSec Policy Group Values
	policyGroupName := "ipsec-policy-group"

	// Proposal Name
	proposalName := "ipsec-proposal"

	// Define IpSecPolicy Expected Values
	tunnel := true
	sourceAddress := "172.20.0.0/16"
	sourcePort := 0
	destinationAddress := "10.20.0.0/16"
	destinationPort := 0
	protocol := "all"
	template := false
	action := "encrypt"
	level := "require"
	ipSecProtocol := "esp"
	disabled := true

	updatedAction := "encrypt"
	updatedLevel := "require"
	updatedIpSecProtocol := "esp"
	updatedSourcePort := 0
	updatedProtocol := "egp"
	updatedDisabled := false

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecPolicyDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecPolicy(
					proposalName,
					peerName,
					policyGroupName,
					tunnel,
					sourceAddress,
					sourcePort,
					destinationAddress,
					destinationPort,
					protocol,
					template,
					action,
					level,
					ipSecProtocol,
					disabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "peer", peerName),
					resource.TestCheckResourceAttr(resourceName, "tunnel", strconv.FormatBool(tunnel)),
					resource.TestCheckResourceAttr(resourceName, "source_address", sourceAddress),
					resource.TestCheckResourceAttr(resourceName, "source_port", strconv.Itoa(sourcePort)),
					resource.TestCheckResourceAttr(resourceName, "destination_address", destinationAddress),
					resource.TestCheckResourceAttr(resourceName, "destination_port", strconv.Itoa(destinationPort)),
					resource.TestCheckResourceAttr(resourceName, "protocol", protocol),
					resource.TestCheckResourceAttr(resourceName, "template", strconv.FormatBool(template)),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "level", level),
					resource.TestCheckResourceAttr(resourceName, "ipsec_protocol", ipSecProtocol),
					resource.TestCheckResourceAttr(resourceName, "proposal", proposalName),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecPolicy(
					proposalName,
					peerName,
					policyGroupName,
					tunnel,
					sourceAddress,
					updatedSourcePort,
					destinationAddress,
					destinationPort,
					updatedProtocol,
					template,
					updatedAction,
					updatedLevel,
					updatedIpSecProtocol,
					updatedDisabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "peer", peerName),
					resource.TestCheckResourceAttr(resourceName, "tunnel", strconv.FormatBool(tunnel)),
					resource.TestCheckResourceAttr(resourceName, "source_address", sourceAddress),
					resource.TestCheckResourceAttr(resourceName, "source_port", strconv.Itoa(updatedSourcePort)),
					resource.TestCheckResourceAttr(resourceName, "destination_address", destinationAddress),
					resource.TestCheckResourceAttr(resourceName, "destination_port", strconv.Itoa(destinationPort)),
					resource.TestCheckResourceAttr(resourceName, "protocol", updatedProtocol),
					resource.TestCheckResourceAttr(resourceName, "template", strconv.FormatBool(template)),
					resource.TestCheckResourceAttr(resourceName, "action", updatedAction),
					resource.TestCheckResourceAttr(resourceName, "level", updatedLevel),
					resource.TestCheckResourceAttr(resourceName, "ipsec_protocol", updatedIpSecProtocol),
					resource.TestCheckResourceAttr(resourceName, "proposal", proposalName),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(updatedDisabled)),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecPolicyExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecPolicy(rs.Primary.ID)

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
func testAccCheckIpSecPolicyDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not IPSec Policy
		if rs.Type != "mikrotik_ipsec_policy" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindIpSecPolicy(rs.Primary.ID)

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
func testAccIpSecPolicy(
	proposalName string,
	peerName string,
	policyGroupName string,
	tunnel bool,
	sourceAddress string,
	sourcePort int,
	destinationAddress string,
	destinationPort int,
	protocol string,
	template bool,
	action string,
	level string,
	ipsecProtocol string,
	disabled bool,
) string {

	// Profile Name
	ipsecProposalName := proposalName

	// Profile Name
	ipsecProfileName := "ipsec-profile"

	// Policy Template Group
	ipsecPolicyTemplateGroupName := policyGroupName

	// Peer Name
	ipsecPeerName := peerName

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_ipsec_proposal" "testacc_proposal" {
			name = %q
			auth_algorithms = %q
			enc_algorithms = %q
			lifetime = %q
			pfs_group = %q
			disabled = %t
		}
		resource "mikrotik_ipsec_profile" "testacc_profile" {
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
		resource "mikrotik_ipsec_policy_group" "testacc_policy_group" {
			name = %q
		}
		resource "mikrotik_ipsec_peer" "testacc_peer" {
			name = %q
			address = %q
			profile = mikrotik_ipsec_profile.testacc_profile.name
			exchange_mode = %q
			send_initial_contact = %t
			passive = %t
			local_address = %q
			port = %d
		}
		resource "mikrotik_ipsec_policy" "testacc" {
			peer = mikrotik_ipsec_peer.testacc_peer.name
			tunnel = %t
			source_address = %q
			source_port = %d
			destination_address = %q
			destination_port = %d
			protocol = %q
			template = %t
			action = %q
			level = %q
			ipsec_protocol = %q
			proposal = mikrotik_ipsec_proposal.testacc_proposal.name
			disabled = %t
		}
	`,
		ipsecProposalName, "sha512,sha256,sha1", "aes-256-cbc", "30m", "modp2048", false,
		ipsecProfileName, "modp2048", "2m", 5, "aes-256,aes-192,aes-128", "sha1", "1h30m", true, "obey",
		ipsecPolicyTemplateGroupName,
		ipsecPeerName, "192.16.2.14/32", "ike2", false, true, "192.16.3.19", 0,

		tunnel, sourceAddress, sourcePort, destinationAddress, destinationPort,
		protocol, template, action, level, ipsecProtocol, disabled)
}
