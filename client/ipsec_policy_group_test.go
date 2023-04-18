package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecPolicyGroup ADD and Remove Operations
 */
func TestAddIpSecPolicyGroupAndDeleteIpSecPolicyGroup(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define IpSecPolicyGroup Expected Values
	name := "TestName"

	// New Name
	updatedName := "NewTestName"

	// Expected IPSec Policy Group
	expectedIpSecPolicyGroup := &IpSecPolicyGroup{
		Name: name,
	}

	// Adding IpSecPolicyGroup
	ipSecPolicyGroup, err := c.AddIpSecPolicyGroup(expectedIpSecPolicyGroup)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Policy Group with: %v", err)
	}

	// Check ID
	expectedIpSecPolicyGroup.Id = ipSecPolicyGroup.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicyGroup, expectedIpSecPolicyGroup) {

		// Print Error
		t.Errorf("The IPSec Policy Group does not match what we expected. actual: %v expected: %v", ipSecPolicyGroup, expectedIpSecPolicyGroup)
	}

	// Update Fields
	expectedIpSecPolicyGroup.Name = updatedName

	// Execute Update
	ipSecPolicyGroup, err = c.UpdateIpSecPolicyGroup(expectedIpSecPolicyGroup)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Policy Group with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicyGroup, expectedIpSecPolicyGroup) {

		// Print Error
		t.Errorf("The IPSec Policy Group does not match what we expected. actual: %v expected: %v", ipSecPolicyGroup, expectedIpSecPolicyGroup)
	}

	// Find IPSecProfile
	foundIpSecPolicyGroup, err := c.FindIpSecPolicyGroup(ipSecPolicyGroup.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Policy Group with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicyGroup, foundIpSecPolicyGroup) {

		// Print Error
		t.Errorf("The IPSec Policy Group does not match what we expected. actual: %v expected: %v", ipSecPolicyGroup, foundIpSecPolicyGroup)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecPolicyGroup(ipSecPolicyGroup.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Policy Group with: %v", err)
	}
}
