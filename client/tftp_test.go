package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for Tftp ADD and Remove Operations
 */
func TestAddTftpAndDeleteTftp(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

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

	// Expected TFTP
	expectedTftp := &Tftp{
		IpAddresses:     ipAddresses,
		RequestFileName: requestFileName,
		RealFileName:    realFileName,
		Allow:           allow,
		ReadOnly:        readOnly,
		Disabled:        disabled,
		Comment:         comment,
	}

	// Adding TFTP
	tftp, err := c.AddTftp(expectedTftp)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an TFTP with: %v", err)
	}

	// Check ID
	expectedTftp.Id = tftp.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(tftp, expectedTftp) {

		// Print Error
		t.Errorf("The TFTP does not match what we expected. actual: %v expected: %v", tftp, expectedTftp)
	}

	// Update Fields
	expectedTftp.IpAddresses = updatedIpAddresses
	expectedTftp.RequestFileName = updatedRequestFileName
	expectedTftp.RealFileName = updatedRealFileName
	expectedTftp.Disabled = updatedDisabled

	// Execute Update
	tftp, err = c.UpdateTftp(expectedTftp)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an TFTP with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(tftp, expectedTftp) {

		// Print Error
		t.Errorf("The TFTP does not match what we expected. actual: %v expected: %v", tftp, expectedTftp)
	}

	// Find TFTP
	foundTftp, err := c.FindTftp(tftp.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an TFTP with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(tftp, foundTftp) {

		// Print Error
		t.Errorf("The TFTP does not match what we expected. actual: %v expected: %v", foundTftp, tftp)
	}

	// Delete TFTP
	err = c.DeleteTftp(tftp.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an TFTP with: %v", err)
	}
}
