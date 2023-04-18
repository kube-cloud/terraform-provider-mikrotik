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
 * IPSec Identity Resource Create Test
 */
func TestIpSecIdentity_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_ipsec_identity.testacc"

	// Define IPSec Peer Values
	peerName := "ipsec-peer"

	// Define IPSec Policy Group Values
	policyGroupName := "ipsec-policy-group"

	// Define IpSecIdentity Expected Values
	authMethod := "pre-shared-key"
	secret := "iosxwcf13t6èèu"
	username := ""
	password := ""
	eapMethods := ""
	certificate := ""
	remoteCertificate := ""
	key := ""
	remoteKey := ""
	noTrackChain := ""
	myId := "auto"
	remoteId := "auto"
	matchBy := ""
	modeConfig := ""
	generatePolicy := "no"
	comment := ""
	disabled := false

	updatedDisabled := true
	updatedComment := "Updated Comment"

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckIpSecIdentityDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccIpSecIdentity(
					peerName,
					policyGroupName,
					authMethod,
					secret,
					username,
					password,
					eapMethods,
					certificate,
					remoteCertificate,
					key,
					remoteKey,
					noTrackChain,
					myId,
					remoteId,
					matchBy,
					modeConfig,
					generatePolicy,
					comment,
					disabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecIdentityExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "peer", peerName),
					resource.TestCheckResourceAttr(resourceName, "policy_template_group", policyGroupName),
					resource.TestCheckResourceAttr(resourceName, "auth_method", authMethod),
					resource.TestCheckResourceAttr(resourceName, "secret", secret),
					resource.TestCheckResourceAttr(resourceName, "username", username),
					resource.TestCheckResourceAttr(resourceName, "password", password),
					resource.TestCheckResourceAttr(resourceName, "eap_methods", eapMethods),
					resource.TestCheckResourceAttr(resourceName, "certificate", certificate),
					resource.TestCheckResourceAttr(resourceName, "remote_certificate", remoteCertificate),
					resource.TestCheckResourceAttr(resourceName, "key", key),
					resource.TestCheckResourceAttr(resourceName, "remote_key", remoteKey),
					resource.TestCheckResourceAttr(resourceName, "generate_policy", generatePolicy),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			{
				// Configure Test Resource
				Config: testAccIpSecIdentity(
					peerName,
					policyGroupName,
					authMethod,
					secret,
					username,
					password,
					eapMethods,
					certificate,
					remoteCertificate,
					key,
					remoteKey,
					noTrackChain,
					myId,
					remoteId,
					matchBy,
					modeConfig,
					generatePolicy,
					updatedComment,
					updatedDisabled,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpSecIdentityExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "peer", peerName),
					resource.TestCheckResourceAttr(resourceName, "policy_template_group", policyGroupName),
					resource.TestCheckResourceAttr(resourceName, "auth_method", authMethod),
					resource.TestCheckResourceAttr(resourceName, "secret", secret),
					resource.TestCheckResourceAttr(resourceName, "username", username),
					resource.TestCheckResourceAttr(resourceName, "password", password),
					resource.TestCheckResourceAttr(resourceName, "eap_methods", eapMethods),
					resource.TestCheckResourceAttr(resourceName, "certificate", certificate),
					resource.TestCheckResourceAttr(resourceName, "remote_certificate", remoteCertificate),
					resource.TestCheckResourceAttr(resourceName, "key", key),
					resource.TestCheckResourceAttr(resourceName, "remote_key", remoteKey),
					resource.TestCheckResourceAttr(resourceName, "generate_policy", generatePolicy),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(updatedDisabled)),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccIpSecIdentityExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindIpSecIdentity(rs.Primary.ID)

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
func testAccCheckIpSecIdentityDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not IPSec Identity
		if rs.Type != "mikrotik_ipsec_identity" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindIpSecIdentity(rs.Primary.ID)

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
func testAccIpSecIdentity(
	peerName string,
	policyGroupName string,
	authMethod string,
	secret string,
	username string,
	password string,
	eapMethods string,
	certificate string,
	remoteCertificate string,
	key string,
	remoteKey string,
	noTrackChain string,
	myId string,
	remoteId string,
	matchBy string,
	modeConfig string,
	generatePolicy string,
	comment string,
	disabled bool,
) string {

	// Profile Name
	ipsecProfileName := "ipsec-profile-name"

	// Policy Template Group
	ipsecPolicyTemplateGroupName := policyGroupName

	// Peer Name
	ipsecPeerName := peerName

	// Format and Resource Resource String
	return fmt.Sprintf(`
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
		resource "mikrotik_ipsec_identity" "testacc" {
			peer = mikrotik_ipsec_peer.testacc_peer.name
			auth_method = %q
			secret = %q
			username = %q
			password = %q
			eap_methods = %q
			certificate = %q
			remote_certificate = %q
			key = %q
			remote_key = %q
			no_track_chain = %q
			my_id = %q
			remote_id = %q
			match_by = %q
			mode_config = %q
			generate_policy = %q
			policy_template_group = mikrotik_ipsec_policy_group.testacc_policy_group.name
			comment = %q
			disabled = %t
		}
	`,
		ipsecProfileName, "modp2048", "2m", 5, "aes-256,aes-192,aes-128", "sha1", "1h30m", true, "obey",
		ipsecPolicyTemplateGroupName,
		ipsecPeerName, "192.16.2.14/32", "ike2", false, true, "192.16.3.19", 0,

		authMethod, secret, username, password, eapMethods, certificate, remoteCertificate, key,
		remoteKey, noTrackChain, myId, remoteId, matchBy, modeConfig, generatePolicy, comment, disabled)
}
