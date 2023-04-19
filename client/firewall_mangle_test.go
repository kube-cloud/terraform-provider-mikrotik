package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for FirewallMangle ADD and Remove Operations
 */
func TestAddFirewallMangleAndDeleteFirewallMangle(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Define FirewallMangle Expected Values
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
	packetMark := "mymark"
	connectionMark := "mymark"
	routingMark := "mymark"
	routingTable := ""
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
	connectionState := ""
	connectionNatState := ""
	tcpFlags := ""

	// New Values
	updatedDisabled := false
	updatedLogPrefix := "[UPDT-TEST-PREFIX] "
	updatedSourcePort := 10
	updatedDestinationPort := 4090
	updatedSourceAddress := "10.20.0.0/16"

	// Expected Firewall Mangle
	expectedFirewallMangle := &FirewallMangle{
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
		ConnectionState:        connectionState,
		ConnectionNatState:     connectionNatState,
		TcpFlags:               tcpFlags,
	}

	// Adding FirewallMangle
	firewallMangle, err := c.AddFirewallMangle(expectedFirewallMangle)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an Firewall Mangle with: %v", err)
	}

	// Check ID
	expectedFirewallMangle.Id = firewallMangle.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallMangle, expectedFirewallMangle) {

		// Print Error
		t.Errorf("The Firewall Mangle does not match what we expected. actual: %v expected: %v", firewallMangle, expectedFirewallMangle)
	}

	// Update Fields
	expectedFirewallMangle.Disabled = updatedDisabled
	expectedFirewallMangle.LogPrefix = updatedLogPrefix
	expectedFirewallMangle.SourcePort = updatedSourcePort
	expectedFirewallMangle.SourceAddress = updatedSourceAddress
	expectedFirewallMangle.DestinationPort = updatedDestinationPort

	// Execute Update
	firewallMangle, err = c.UpdateFirewallMangle(expectedFirewallMangle)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an Firewall Mangle with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallMangle, expectedFirewallMangle) {

		// Print Error
		t.Errorf("The Firewall Mangle does not match what we expected. actual: %v expected: %v", firewallMangle, expectedFirewallMangle)
	}

	// Find IPSecProfile
	foundFirewallMangle, err := c.FindFirewallMangle(firewallMangle.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an Firewall Mangle with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(firewallMangle, foundFirewallMangle) {

		// Print Error
		t.Errorf("The Firewall Mangle does not match what we expected. actual: %v expected: %v", firewallMangle, foundFirewallMangle)
	}

	// Delete IPSecProfile
	err = c.DeleteFirewallMangle(firewallMangle.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an Firewall Mangle with: %v", err)
	}
}
