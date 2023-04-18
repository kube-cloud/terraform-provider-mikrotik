package client

import (
	"fmt"
	"log"
)

/**
 * Define IPSec Policy Group Structure
 */
type IpSecPolicyGroup struct {
	Id   string `mikrotik:".id"`
	Name string `mikrotik:"name"`
}

/**
 * Function used to ADD IPSec Policy Group on Mikrotik Router
 */
func (client Mikrotik) AddIpSecPolicyGroup(peer *IpSecPolicyGroup) (*IpSecPolicyGroup, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Policy Group")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/policy/group/add", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy Group ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return peer by ID
	return client.FindIpSecPolicyGroup(id)
}

/**
 * Function used to List IPSec Policy Group from Mikrotik Router
 */
func (client Mikrotik) ListIpSecPolicyGroup() ([]IpSecPolicyGroup, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Policy Groups")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/group/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy Group List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Policy Group Array Pointer
	ipSecPolicyGroups := []IpSecPolicyGroup{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPolicyGroups)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecPolicyGroups, nil
}

/**
 * Function used to FIND IPSec Policy Group by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecPolicyGroup(id string) (*IpSecPolicyGroup, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Policy Group: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/group/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy Group Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Policy Group Pointer
	ipSecPolicyGroup := IpSecPolicyGroup{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPolicyGroup)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if ipSecPolicyGroup.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("IPSec Policy Group `%s` not found", id))
	}

	// Return result
	return &ipSecPolicyGroup, nil
}

/**
 * Function used to UPDATE IPSec Policy Group on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecPolicyGroup(peer *IpSecPolicyGroup) (*IpSecPolicyGroup, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/policy/group/set", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy Group ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return peer by ID
	return client.FindIpSecPolicyGroup(peer.Id)
}

/**
 * Function used to DELETE IPSec Policy Group by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecPolicyGroup(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Policy Group: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/group/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
