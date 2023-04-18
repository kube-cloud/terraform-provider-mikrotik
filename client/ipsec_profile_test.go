package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecProfile ADD and Remove Operations
 */
func TestAddIpSecProfileAndDeleteIpSecProfile(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define IpSecProfile Expected Values
	name := "TestName"
	dhGroup := "modp2048"
	dpdInterval := "2m"
	dpdMaxFailure := 5
	encAlgorithms := "aes-256,aes-192,aes-128"
	hashAlgorithm := "sha1"
	lifetime := "1h30m"
	natTraversal := true
	proposalCheck := "obey"
	updatedEncAlgorithms := "aes-256"
	updatedHashAlgorithm := "sha256"
	updatedNatTransversal := false

	// Expected IPSec Profile
	expectedIpSecProfile := &IpSecProfile{
		Name:          name,
		DhGroup:       dhGroup,
		DpdInterval:   dpdInterval,
		DpdMaxFailure: dpdMaxFailure,
		EncAlgorithms: encAlgorithms,
		HashAlgorithm: hashAlgorithm,
		Lifetime:      lifetime,
		NatTraversal:  natTraversal,
		ProposalCheck: proposalCheck,
	}

	// Adding IpSecProfile
	ipSecProfile, err := c.AddIpSecProfile(expectedIpSecProfile)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Profile with: %v", err)
	}

	// Check ID
	expectedIpSecProfile.Id = ipSecProfile.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProfile, expectedIpSecProfile) {

		// Print Error
		t.Errorf("The IPSec Profile does not match what we expected. actual: %v expected: %v", ipSecProfile, expectedIpSecProfile)
	}

	// Update Fields
	expectedIpSecProfile.EncAlgorithms = updatedEncAlgorithms
	expectedIpSecProfile.HashAlgorithm = updatedHashAlgorithm
	expectedIpSecProfile.NatTraversal = updatedNatTransversal

	// Execute Update
	ipSecProfile, err = c.UpdateIpSecProfile(expectedIpSecProfile)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Profile with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProfile, expectedIpSecProfile) {

		// Print Error
		t.Errorf("The IPSec Profile does not match what we expected. actual: %v expected: %v", ipSecProfile, expectedIpSecProfile)
	}

	// Find IPSecProfile
	foundIpSecProfile, err := c.FindIpSecProfile(ipSecProfile.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Profile with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecProfile, foundIpSecProfile) {

		// Print Error
		t.Errorf("The IPSec Profile does not match what we expected. actual: %v expected: %v", ipSecProfile, foundIpSecProfile)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecProfile(ipSecProfile.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Profile with: %v", err)
	}
}
