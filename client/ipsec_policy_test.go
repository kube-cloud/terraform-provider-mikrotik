package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecPolicy ADD and Remove Operations
 */
func TestAddIpSecPolicyAndDeleteIpSecPolicy(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Expected IPSec Proposal
	expectedIpSecProposal := &IpSecProposal{
		Name:           "TestName",
		AuthAlgorithms: "sha512,sha256,sha1",
		EncAlgorithms:  "aes-256-cbc,aes-192-cbc,aes-128-cbc",
		Lifetime:       "30m",
		PfsGroup:       "modp2048",
		Disabled:       false,
	}

	// Adding IPsecProposal
	ipSecProposal, _ := c.AddIpSecProposal(expectedIpSecProposal)

	// Expected IPSec Profile
	expectedIpSecProfile := &IpSecProfile{
		Name:          "TestName",
		DhGroup:       "modp2048",
		DpdInterval:   "2m",
		DpdMaxFailure: 5,
		EncAlgorithms: "aes-256,aes-192,aes-128",
		HashAlgorithm: "sha1",
		Lifetime:      "1h30m",
		NatTraversal:  true,
		ProposalCheck: "obey",
	}

	// Adding IpSecProfile
	ipSecProfile, _ := c.AddIpSecProfile(expectedIpSecProfile)

	// Expected IPSec Peer
	expectedIpSecPeer := &IpSecPeer{
		Name:               "TestName",
		Address:            "192.16.2.14/32",
		Profile:            ipSecProfile.Name,
		ExchangeMode:       "ike2",
		SendInitialContact: false,
		Passive:            true,
		LocalAddress:       "192.16.3.19",
		Port:               0,
	}

	// Adding IpSecPeer
	ipSecPeer, _ := c.AddIpSecPeer(expectedIpSecPeer)

	// Define IpSecPolicy Expected Values
	peer := ipSecPeer.Name
	tunnel := true
	sourceAddress := "172.20.0.0/16"
	sourcePort := 0
	destinationAddress := "10.20.0.0/16"
	destinationPort := 0
	protocol := "all"
	template := false
	action := "encrypt"
	level := "require"
	ipSecProtocol := "esp"
	proposal := ipSecProposal.Name
	disabled := true
	updatedAction := "encrypt"
	updatedLevel := "require"
	updatedIpSecProtocol := "esp"
	updatedSourcePort := 0
	updatedProtocol := "egp"
	updatedDisabled := false

	// Expected IPSec Policy
	expectedIpSecPolicy := &IpSecPolicy{
		Peer:               peer,
		Tunnel:             tunnel,
		SourceAddress:      sourceAddress,
		SourcePort:         sourcePort,
		DestinationAddress: destinationAddress,
		DestinationPort:    destinationPort,
		Protocol:           protocol,
		Template:           template,
		Action:             action,
		Level:              level,
		IpSecProtocol:      ipSecProtocol,
		Proposal:           proposal,
		Disabled:           disabled,
	}

	// Adding IpSecPolicy
	ipSecPolicy, err := c.AddIpSecPolicy(expectedIpSecPolicy)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Policy with: %v", err)
	}

	// Check ID
	expectedIpSecPolicy.Id = ipSecPolicy.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicy, expectedIpSecPolicy) {

		// Print Error
		t.Errorf("The IPSec Policy does not match what we expected. actual: %v expected: %v", ipSecPolicy, expectedIpSecPolicy)
	}

	// Update Fields
	expectedIpSecPolicy.Level = updatedLevel
	expectedIpSecPolicy.IpSecProtocol = updatedIpSecProtocol
	expectedIpSecPolicy.SourcePort = updatedSourcePort
	expectedIpSecPolicy.Action = updatedAction
	expectedIpSecPolicy.Protocol = updatedProtocol
	expectedIpSecPolicy.Disabled = updatedDisabled

	// Execute Update
	ipSecPolicy, err = c.UpdateIpSecPolicy(expectedIpSecPolicy)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Policy with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicy, expectedIpSecPolicy) {

		// Print Error
		t.Errorf("The IPSec Policy does not match what we expected. actual: %v expected: %v", ipSecPolicy, expectedIpSecPolicy)
	}

	// Find IPSecProfile
	foundIpSecPolicy, err := c.FindIpSecPolicy(ipSecPolicy.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Policy with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPolicy, foundIpSecPolicy) {

		// Print Error
		t.Errorf("The IPSec Policy does not match what we expected. actual: %v expected: %v", foundIpSecPolicy, ipSecPolicy)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecPolicy(ipSecPolicy.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Policy with: %v", err)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecPeer(ipSecPeer.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Peer with: %v", err)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecProfile(ipSecProfile.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Profile with: %v", err)
	}

	// Delete IPSecProposal
	err = c.DeleteIpSecProposal(ipSecProposal.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Proposal with: %v", err)
	}
}
