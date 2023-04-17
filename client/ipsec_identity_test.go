package client

import (
	"reflect"
	"testing"
)

/**
 * Test Method for IpSecIdentity ADD and Remove Operations
 */
func TestAddIpSecIdentityAndDeleteIpSecIdentity(t *testing.T) {

	// Get Client from Environments Configuration
	c := NewClient(GetConfigFromEnv())

	// Expected IPSec Policy Group
	expectedIpSecPolicyGroup := &IpSecPolicyGroup{
		Name: "TestGroup",
	}

	// Adding IpSecPolicyGroup
	ipSecPolicyGroup, _ := c.AddIpSecPolicyGroup(expectedIpSecPolicyGroup)

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

	// Expected IPSec Identity
	expectedIpSecPeer := &IpSecPeer{
		Name:               "TestName",
		Address:            "192.16.2.14/32",
		Profile:            expectedIpSecProfile.Name,
		ExchangeMode:       "ike2",
		SendInitialContact: false,
		Passive:            true,
		LocalAddress:       "192.16.3.19",
		Port:               0,
	}

	// Adding IpSecPeer
	ipSecPeer, _ := c.AddIpSecPeer(expectedIpSecPeer)

	// Define IpSecIdentity Expected Values
	peer := ipSecPeer.Name
	authMethod := "pre-shared-key"
	secret := "iosxwcf13t6èèu"
	noTrackChain := ""
	myId := "auto"
	remoteId := "auto"
	matchBy := ""
	generatePolicy := "no"
	disabled := false
	policyTemplateGroup := expectedIpSecPolicyGroup.Name

	// Updated Generated Pocily (updatedGeneratePolicy)
	updatedGeneratePolicy := "port-strict"

	// Updated MyID
	updatedMyId := "fqdn:MyID"

	// Expected IPSec Identity
	expectedIpSecIdentity := &IpSecIdentity{
		Peer:                peer,
		AuthMethod:          authMethod,
		Secret:              secret,
		NoTrackChain:        noTrackChain,
		MyId:                myId,
		RemoteId:            remoteId,
		MatchBy:             matchBy,
		GeneratePolicy:      generatePolicy,
		PolicyTemplateGroup: policyTemplateGroup,
		Disabled:            disabled,
	}

	// Adding IpSecIdentity
	ipSecIdentity, err := c.AddIpSecIdentity(expectedIpSecIdentity)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Adding an IPSec Identity with: %v", err)
	}

	// Check ID
	expectedIpSecIdentity.Id = ipSecIdentity.Id

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecIdentity, expectedIpSecIdentity) {

		// Print Error
		t.Errorf("The IPSec Identity does not match what we expected. actual: %v expected: %v", ipSecIdentity, expectedIpSecIdentity)
	}

	// Update Fields
	expectedIpSecIdentity.GeneratePolicy = updatedGeneratePolicy
	expectedIpSecIdentity.MyId = updatedMyId

	// Execute Update
	ipSecIdentity, err = c.UpdateIpSecIdentity(expectedIpSecIdentity)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Updating an IPSec Identity with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecIdentity, expectedIpSecIdentity) {

		// Print Error
		t.Errorf("The IPSec Identity does not match what we expected. actual: %v expected: %v", ipSecIdentity, expectedIpSecIdentity)
	}

	// Find IPSecProfile
	foundIpSecIdentity, err := c.FindIpSecIdentity(ipSecIdentity.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Find an IPSec Identity with: %v", err)
	}

	// If Deep Compare Failed
	if !reflect.DeepEqual(ipSecIdentity, foundIpSecIdentity) {

		// Print Error
		t.Errorf("The IPSec Identity does not match what we expected. actual: %v expected: %v", ipSecIdentity, foundIpSecIdentity)
	}

	// Delete IPSecProfile
	err = c.DeleteIpSecIdentity(ipSecIdentity.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Identity with: %v", err)
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

	// Delete IPSecProfile
	err = c.DeleteIpSecPolicyGroup(ipSecPolicyGroup.Id)

	// If There is Error
	if err != nil {

		// Log
		t.Errorf("Error Delete an IPSec Policy Group with: %v", err)
	}
}
