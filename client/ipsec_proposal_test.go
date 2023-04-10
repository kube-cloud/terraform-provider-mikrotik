package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecProposal ADD and Remove Operations
 */
func TestAddIpAddressAndDeleteIpSecProposal(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define IpSecProposal Expected Values
	name := "TestName"
	authAlgorithms := "sha512,sha256,sha1"
	encAlgorithms := "aes-256-cbc,aes-192-cbc,aes-128-cbc"
	lifetime := "30m"
	pfsGroup := "modp2048"
	disabled := false
	updatedEncAlgorithms := "aes-256-cbc"

	// Expected IPSec Proposal
	expectedIpSecProposal := &IpSecProposal{
		Name:           name,
		AuthAlgorithms: authAlgorithms,
		EncAlgorithms:  encAlgorithms,
		Lifetime:       lifetime,
		PfsGroup:       pfsGroup,
		Disabled:       disabled,
	}

	// Adding IPsecProposal
	ipSecProposal, err := c.AddIpSecProposal(expectedIpSecProposal)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Proposal with: %v", err)
	}

	// Check ID
	expectedIpSecProposal.Id = ipSecProposal.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProposal, expectedIpSecProposal) {

		// Print Error
		t.Errorf("The IPSec Proposal does not match what we expected. actual: %v expected: %v", ipSecProposal, expectedIpSecProposal)
	}

	// Update Encryption Algorithms
	expectedIpSecProposal.EncAlgorithms = updatedEncAlgorithms

	// Execute Update
	ipSecProposal, err = c.UpdateIpSecProposal(expectedIpSecProposal)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Proposal with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProposal, expectedIpSecProposal) {

		// Print Error
		t.Errorf("The IPSec Proposal does not match what we expected. actual: %v expected: %v", ipSecProposal, expectedIpSecProposal)
	}

	// Find IPSecProposal
	foundIpSecProposal, err := c.FindIpSecProposal(ipSecProposal.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Proposal with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProposal, foundIpSecProposal) {

		// Print Error
		t.Errorf("The IPSec Proposal does not match what we expected. actual: %v expected: %v", ipSecProposal, foundIpSecProposal)
	}

	// Delete IPSecProposal
	err = c.DeleteIpSecProposal(ipSecProposal.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Proposal with: %v", err)
	}
}
