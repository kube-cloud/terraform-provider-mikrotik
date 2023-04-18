package client

import (
	"fmt"
	"log"
)

/**
 * Define IPSec Peer Structure
 */
type IpSecPeer struct {
	Id                 string `mikrotik:".id"`
	Name               string `mikrotik:"name"`
	Address            string `mikrotik:"address"`
	Profile            string `mikrotik:"profile"`
	ExchangeMode       string `mikrotik:"exchange-mode"`
	SendInitialContact bool   `mikrotik:"send-initial-contact"`
	Passive            bool   `mikrotik:"passive"`
	LocalAddress       string `mikrotik:"local-address"`
	Port               int    `mikrotik:"port"`
}

/**
 * Function used to ADD IPSec Peer on Mikrotik Router
 */
func (client Mikrotik) AddIpSecPeer(peer *IpSecPeer) (*IpSecPeer, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Peer")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/peer/add", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Peer ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return peer by ID
	return client.FindIpSecPeer(id)
}

/**
 * Function used to List IPSec Peer from Mikrotik Router
 */
func (client Mikrotik) ListIpSecPeer() ([]IpSecPeer, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Peers")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/peer/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Peer List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Peer Array Pointer
	ipSecPeers := []IpSecPeer{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPeers)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecPeers, nil
}

/**
 * Function used to FIND IPSec Peer by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecPeer(id string) (*IpSecPeer, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Peer: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/peer/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Peer Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Peer Pointer
	ipSecPeer := IpSecPeer{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPeer)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if ipSecPeer.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("IPSec Peer `%s` not found", id))
	}

	// If Port is 0
	if ipSecPeer.Port == 0 {

		// Cancel Port
		ipSecPeer.Port = 0
	}

	// Return result
	return &ipSecPeer, nil
}

/**
 * Function used to UPDATE IPSec Peer on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecPeer(peer *IpSecPeer) (*IpSecPeer, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/peer/set", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Peer ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return peer by ID
	return client.FindIpSecPeer(peer.Id)
}

/**
 * Function used to DELETE IPSec Peer by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecPeer(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Peer: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/peer/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
