package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for FirewallRaw ADD and Remove Operations
 */
func TestAddFirewallRawAndDeleteFirewallRaw(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define FirewallRaw Expected Values
	chain := "prerouting"
	sourceAddress := "10.0.0.0/8"
	destinationAddress := "10.0.0.0/8"
	sourcePort := 8080
	destinationPort := 9090
	anyPort := 0
	protocol := "tcp"
	inInterface := ""
	outInterface := ""
	inInterfaceList := ""
	outInterfaceList := ""
	sourceAddressList := ""
	destinationAddressList := ""
	sourceMacAddress := "00:AB:AC:29:22:31"
	ipSecPolicy := "in,ipsec"
	action := "accept"
	log := true
	logPrefix := "[TEST-PREFIX] "
	disabled := true

	// New Values
	updatedDisabled := false
	updatedLogPrefix := "[UPDT-TEST-PREFIX] "
	updatedSourcePort := 10
	updatedDestinationPort := 4090
	updatedSourceAddress := "10.20.0.0/16"

	// Expected Firewall Raw
	expectedFirewallRaw := &FirewallRaw{
		Chain:                  chain,
		Disabled:               disabled,
		SourceAddress:          sourceAddress,
		DestinationAddress:     destinationAddress,
		SourcePort:             sourcePort,
		DestinationPort:        destinationPort,
		AnyPort:                anyPort,
		Protocol:               protocol,
		InInterface:            inInterface,
		OutInterface:           outInterface,
		InInterfaceList:        inInterfaceList,
		OutInterfaceList:       outInterfaceList,
		SourceAddressList:      sourceAddressList,
		DestinationAddressList: destinationAddressList,
		SourceMacAddress:       sourceMacAddress,
		IpSecPolicy:            ipSecPolicy,
		Action:                 action,
		Log:                    log,
		LogPrefix:              logPrefix,
	}

	// Adding FirewallRaw
	firewallRaw, err := c.AddFirewallRaw(expectedFirewallRaw)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an Firewall Raw with: %v", err)
	}

	// Check ID
	expectedFirewallRaw.Id = firewallRaw.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRaw, expectedFirewallRaw) {

		// Print Error
		t.Errorf("The Firewall Raw does not match what we expected. actual: %v expected: %v", firewallRaw, expectedFirewallRaw)
	}

	// Update Fields
	expectedFirewallRaw.Disabled = updatedDisabled
	expectedFirewallRaw.LogPrefix = updatedLogPrefix
	expectedFirewallRaw.SourcePort = updatedSourcePort
	expectedFirewallRaw.SourceAddress = updatedSourceAddress
	expectedFirewallRaw.DestinationPort = updatedDestinationPort

	// Execute Update
	firewallRaw, err = c.UpdateFirewallRaw(expectedFirewallRaw)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an Firewall Raw with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRaw, expectedFirewallRaw) {

		// Print Error
		t.Errorf("The Firewall Raw does not match what we expected. actual: %v expected: %v", firewallRaw, expectedFirewallRaw)
	}

	// Find IPSecProfile
	foundFirewallRaw, err := c.FindFirewallRaw(firewallRaw.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an Firewall Raw with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRaw, foundFirewallRaw) {

		// Print Error
		t.Errorf("The Firewall Raw does not match what we expected. actual: %v expected: %v", firewallRaw, foundFirewallRaw)
	}

	// Delete IPSecProfile
	err = c.DeleteFirewallRaw(firewallRaw.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an Firewall Raw with: %v", err)
	}
}
