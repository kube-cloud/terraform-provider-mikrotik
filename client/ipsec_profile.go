package client

import (
	"fmt"
	"log"
)

/**
 * Define IPSec Profile Structure
 */
type IpSecProfile struct {
	Id            string `mikrotik:".id"`
	Name          string `mikrotik:"name"`
	DhGroup       string `mikrotik:"dh-group"`
	DpdInterval   string `mikrotik:"dpd-interval"`
	DpdMaxFailure int    `mikrotik:"dpd-maximum-failures"`
	EncAlgorithms string `mikrotik:"enc-algorithm"`
	HashAlgorithm string `mikrotik:"hash-algorithm"`
	Lifetime      string `mikrotik:"lifetime"`
	NatTraversal  bool   `mikrotik:"nat-traversal"`
	ProposalCheck string `mikrotik:"proposal-check"`
}

/**
 * Function used to ADD IPSec Profile on Mikrotik Router
 */
func (client Mikrotik) AddIpSecProfile(profile *IpSecProfile) (*IpSecProfile, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Profile")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/profile/add", profile)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Profile ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return profile by ID
	return client.FindIpSecProfile(id)
}

/**
 * Function used to List IPSec Profile from Mikrotik Router
 */
func (client Mikrotik) ListIpSecProfile() ([]IpSecProfile, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Profiles")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/profile/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Profile List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Profile Array Pointer
	ipSecProfiles := []IpSecProfile{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecProfiles)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecProfiles, nil
}

/**
 * Function used to FIND IPSec Profile by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecProfile(id string) (*IpSecProfile, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Profile: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/profile/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Profile Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Profile Pointer
	ipSecProfile := IpSecProfile{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecProfile)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if ipSecProfile.Name == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("IPSec Profile `%s` not found", id))
	}

	// Return result
	return &ipSecProfile, nil
}

/**
 * Function used to UPDATE IPSec Profile on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecProfile(profile *IpSecProfile) (*IpSecProfile, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/profile/set", profile)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Profile ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return profile by ID
	return client.FindIpSecProfile(profile.Id)
}

/**
 * Function used to DELETE IPSec Profile by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecProfile(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Profile: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/profile/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
