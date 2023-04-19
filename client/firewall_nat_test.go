package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for FirewallNat ADD and Remove Operations
 */
func TestAddFirewallNatAndDeleteFirewallNat(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define FirewallNat Expected Values
	chain := "input"
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
	packetMark := "mymark"
	connectionMark := "mymark"
	routingMark := "mymark"
	routingTable := "mytable"
	connectionType := "ftp"
	sourceAddressList := ""
	destinationAddressList := ""
	layer7Protocol := ""
	sourceMacAddress := "00:AB:AC:29:22:31"
	ipSecPolicy := "in,ipsec"
	inBridgePort := ""
	outBridgePort := ""
	inBridgePortList := ""
	outBridgePortList := ""
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

	// Expected Firewall Nat
	expectedFirewallNat := &FirewallNat{
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
		PacketMark:             packetMark,
		ConnectionMark:         connectionMark,
		RoutingMark:            routingMark,
		RoutingTable:           routingTable,
		ConnectionType:         connectionType,
		SourceAddressList:      sourceAddressList,
		DestinationAddressList: destinationAddressList,
		Layer7Protocol:         layer7Protocol,
		SourceMacAddress:       sourceMacAddress,
		IpSecPolicy:            ipSecPolicy,
		InBridgePort:           inBridgePort,
		OutBridgePort:          outBridgePort,
		InBridgePortList:       inBridgePortList,
		OutBridgePortList:      outBridgePortList,
		Action:                 action,
		Log:                    log,
		LogPrefix:              logPrefix,
	}

	// Adding FirewallNat
	firewallNat, err := c.AddFirewallNat(expectedFirewallNat)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an Firewall Nat with: %v", err)
	}

	// Check ID
	expectedFirewallNat.Id = firewallNat.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallNat, expectedFirewallNat) {

		// Print Error
		t.Errorf("The Firewall Nat does not match what we expected. actual: %v expected: %v", firewallNat, expectedFirewallNat)
	}

	// Update Fields
	expectedFirewallNat.Disabled = updatedDisabled
	expectedFirewallNat.LogPrefix = updatedLogPrefix
	expectedFirewallNat.SourcePort = updatedSourcePort
	expectedFirewallNat.SourceAddress = updatedSourceAddress
	expectedFirewallNat.DestinationPort = updatedDestinationPort

	// Execute Update
	firewallNat, err = c.UpdateFirewallNat(expectedFirewallNat)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an Firewall Nat with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallNat, expectedFirewallNat) {

		// Print Error
		t.Errorf("The Firewall Nat does not match what we expected. actual: %v expected: %v", firewallNat, expectedFirewallNat)
	}

	// Find IPSecProfile
	foundFirewallNat, err := c.FindFirewallNat(firewallNat.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an Firewall Nat with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallNat, foundFirewallNat) {

		// Print Error
		t.Errorf("The Firewall Nat does not match what we expected. actual: %v expected: %v", firewallNat, foundFirewallNat)
	}

	// Delete IPSecProfile
	err = c.DeleteFirewallNat(firewallNat.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an Firewall Nat with: %v", err)
	}
}
