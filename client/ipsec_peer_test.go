package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecPeer ADD and Remove Operations
 */
func TestAddIpSecPeerAndDeleteIpSecPeer(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

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

	// Define IpSecPeer Expected Values
	name := "TestName"
	address := "192.16.2.14/32"
	profile := expectedIpSecProfile.Name
	exchangeMode := "ike2"
	sendInitialContact := false
	passive := true
	localAddress := "192.16.3.19"
	port := 0
	updatedSendInitialContact := true
	updatedPassive := false

	// Expected IPSec Peer
	expectedIpSecPeer := &IpSecPeer{
		Name:               name,
		Address:            address,
		Profile:            profile,
		ExchangeMode:       exchangeMode,
		SendInitialContact: sendInitialContact,
		Passive:            passive,
		LocalAddress:       localAddress,
		Port:               port,
	}

	// Adding IpSecPeer
	ipSecPeer, err := c.AddIpSecPeer(expectedIpSecPeer)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Peer with: %v", err)
	}

	// Check ID
	expectedIpSecPeer.Id = ipSecPeer.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPeer, expectedIpSecPeer) {

		// Print Error
		t.Errorf("The IPSec Peer does not match what we expected. actual: %v expected: %v", ipSecPeer, expectedIpSecPeer)
	}

	// Update Fields
	expectedIpSecPeer.SendInitialContact = updatedSendInitialContact
	expectedIpSecPeer.Passive = updatedPassive

	// Execute Update
	ipSecPeer, err = c.UpdateIpSecPeer(expectedIpSecPeer)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Peer with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPeer, expectedIpSecPeer) {

		// Print Error
		t.Errorf("The IPSec Peer does not match what we expected. actual: %v expected: %v", ipSecPeer, expectedIpSecPeer)
	}

	// Find IPSecProfile
	foundIpSecPeer, err := c.FindIpSecPeer(ipSecPeer.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Peer with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecPeer, foundIpSecPeer) {

		// Print Error
		t.Errorf("The IPSec Peer does not match what we expected. actual: %v expected: %v", ipSecPeer, foundIpSecPeer)
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
}
