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
 * TFTP Resource Create Test
 */
func TestTftp_Create(t *testing.T) {

	// Initialize Resource Name
	resourceName := "mikrotik_tftp.testacc"

	// Define Tftp Expected Values
	ipAddresses := "10.10.10.1"
	requestFileName := "pxelinux.0"
	realFileName := "pxelinux.0"
	allow := true
	readOnly := true
	disabled := false
	comment := "Test TFTP"
	updatedIpAddresses := "10.11.11.1"
	updatedRequestFileName := "updated.pxelinux.0"
	updatedRealFileName := "updated.pxelinux.0"
	updatedDisabled := true

	// Initialize Test
	resource.Test(t, resource.TestCase{

		// Initialize Test Case Precheck Callback
		PreCheck: func() { testAccPreCheck(t) },

		// Initialize Test Case Provider Factory Callback
		ProviderFactories: testAccProviderFactories,

		// Initialize Check destroy Callback
		CheckDestroy: testAccCheckTftpDestroy,

		// Initialize Test Steps
		Steps: []resource.TestStep{
			{
				// Configure Test Resource
				Config: testAccTftp(
					ipAddresses,
					requestFileName,
					realFileName,
					allow,
					readOnly,
					disabled,
					comment,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccTftpExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses", ipAddresses),
					resource.TestCheckResourceAttr(resourceName, "request_file_name", requestFileName),
					resource.TestCheckResourceAttr(resourceName, "real_file_name", realFileName),
					resource.TestCheckResourceAttr(resourceName, "allow", strconv.FormatBool(allow)),
					resource.TestCheckResourceAttr(resourceName, "read_only", strconv.FormatBool(readOnly)),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(disabled)),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			{
				// Configure Test Resource
				Config: testAccTftp(
					updatedIpAddresses,
					updatedRequestFileName,
					updatedRealFileName,
					allow,
					readOnly,
					updatedDisabled,
					comment,
				),

				// Check Test Resource
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccTftpExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses", updatedIpAddresses),
					resource.TestCheckResourceAttr(resourceName, "request_file_name", updatedRequestFileName),
					resource.TestCheckResourceAttr(resourceName, "real_file_name", updatedRealFileName),
					resource.TestCheckResourceAttr(resourceName, "allow", strconv.FormatBool(allow)),
					resource.TestCheckResourceAttr(resourceName, "read_only", strconv.FormatBool(readOnly)),
					resource.TestCheckResourceAttr(resourceName, "disabled", strconv.FormatBool(updatedDisabled)),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
		},
	})
}

/**
 * Function used to Test if Terraform Resource Exists
 */
func testAccTftpExists(resourceName string) resource.TestCheckFunc {

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
		record, err := c.FindTftp(rs.Primary.ID)

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
func testAccCheckTftpDestroy(s *terraform.State) error {

	// Build Client
	c := client.NewClient(client.GetConfigFromEnv())

	// Iterate over Resources
	for _, rs := range s.RootModule().Resources {

		// If Resource is not TFTP
		if rs.Type != "mikrotik_tftp" {

			// Continue Iteration
			continue
		}

		// Find Resource
		remoteRecord, err := c.FindTftp(rs.Primary.ID)

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
func testAccTftp(
	ipAddresses string,
	requestFileName string,
	realFileName string,
	allow bool,
	readOnly bool,
	disabled bool,
	comment string,
) string {

	// Format and Resource Resource String
	return fmt.Sprintf(`
		resource "mikrotik_tftp" "testacc" {
			ip_addresses = %q
			request_file_name = %q
			real_file_name = %q
			allow = %t
			read_only = %t
			disabled = %t
			comment = %q
		}
	`, ipAddresses, requestFileName, realFileName, allow, readOnly, disabled, comment)
}
