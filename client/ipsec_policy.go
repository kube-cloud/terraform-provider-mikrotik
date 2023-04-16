package client

import (
	"fmt"
	"log"
)

/**
 * Define IPSec Policy Structure
 */
type IpSecPolicy struct {
	Id                 string `mikrotik:".id"`
	Peer               string `mikrotik:"peer"`
	Tunnel             bool   `mikrotik:"tunnel"`
	SourceAddress      string `mikrotik:"src-address"`
	SourcePort         int    `mikrotik:"src-port"`
	DestinationAddress string `mikrotik:"dst-address"`
	DestinationPort    int    `mikrotik:"dst-port"`
	Protocol           string `mikrotik:"protocol"`
	Template           bool   `mikrotik:"template"`
	Action             string `mikrotik:"action"`
	Level              string `mikrotik:"level"`
	IpSecProtocol      string `mikrotik:"ipsec-protocols"`
	Proposal           string `mikrotik:"proposal"`
}

/**
 * Function used to ADD IPSec Policy on Mikrotik Router
 */
func (client Mikrotik) AddIpSecPolicy(ipsecPolicy *IpSecPolicy) (*IpSecPolicy, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Policy")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Force Template to False
	ipsecPolicy.Template = false

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/policy/add", ipsecPolicy)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return ipsecpolicy by ID
	return client.FindIpSecPolicy(id)
}

/**
 * Function used to List IPSec Policy from Mikrotik Router
 */
func (client Mikrotik) ListIpSecPolicy() ([]IpSecPolicy, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Policys")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Policy Array Pointer
	ipSecPolicies := []IpSecPolicy{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPolicies)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecPolicies, nil
}

/**
 * Function used to FIND IPSec Policy by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecPolicy(id string) (*IpSecPolicy, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Policy: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Policy Pointer
	ipSecPolicy := IpSecPolicy{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPolicy)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if ipSecPolicy.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("IPSec Policy `%s` not found", id))
	}

	// Return result
	return &ipSecPolicy, nil
}

/**
 * Function used to UPDATE IPSec Policy on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecPolicy(ipsecPolicy *IpSecPolicy) (*IpSecPolicy, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/policy/set", ipsecPolicy)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Policy ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return ipsecpolicy by ID
	return client.FindIpSecPolicy(ipsecPolicy.Id)
}

/**
 * Function used to DELETE IPSec Policy by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecPolicy(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Policy: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/policy/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
