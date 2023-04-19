package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for FirewallRule ADD and Remove Operations
 */
func TestAddFirewallRuleAndDeleteFirewallRule(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define FirewallRule Expected Values
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

	// Expected Firewall Rule
	expectedFirewallRule := &FirewallRule{
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

	// Adding FirewallRule
	firewallRule, err := c.AddFirewallRule(expectedFirewallRule)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an Firewall Rule with: %v", err)
	}

	// Check ID
	expectedFirewallRule.Id = firewallRule.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRule, expectedFirewallRule) {

		// Print Error
		t.Errorf("The Firewall Rule does not match what we expected. actual: %v expected: %v", firewallRule, expectedFirewallRule)
	}

	// Update Fields
	expectedFirewallRule.Disabled = updatedDisabled
	expectedFirewallRule.LogPrefix = updatedLogPrefix
	expectedFirewallRule.SourcePort = updatedSourcePort
	expectedFirewallRule.SourceAddress = updatedSourceAddress
	expectedFirewallRule.DestinationPort = updatedDestinationPort

	// Execute Update
	firewallRule, err = c.UpdateFirewallRule(expectedFirewallRule)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an Firewall Rule with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRule, expectedFirewallRule) {

		// Print Error
		t.Errorf("The Firewall Rule does not match what we expected. actual: %v expected: %v", firewallRule, expectedFirewallRule)
	}

	// Find IPSecProfile
	foundFirewallRule, err := c.FindFirewallRule(firewallRule.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an Firewall Rule with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallRule, foundFirewallRule) {

		// Print Error
		t.Errorf("The Firewall Rule does not match what we expected. actual: %v expected: %v", firewallRule, foundFirewallRule)
	}

	// Delete IPSecProfile
	err = c.DeleteFirewallRule(firewallRule.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an Firewall Rule with: %v", err)
	}
}
